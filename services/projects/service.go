package projects

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_app/dto"
	"go_app/services/users"
	"net/http"
	"strconv"
)

type ProjectsService struct {
	store ProjectsStore
}

func NewProjectsService(store ProjectsStore) *ProjectsService {
	return &ProjectsService{
		store: store,
	}
}

func (s ProjectsService) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body dto.ProjectRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		var response = dto.ErrorResponse{
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := s.store.CreateProject(body.Name, body.Desc, body.UserId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		var response = dto.ErrorResponse{
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	var response = dto.ProjectResponse{
		id,
		body.Name,
		body.Desc,
		body.UserId,
		nil,
	}

	json.NewEncoder(w).Encode(response)
}

func (s ProjectsService) GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	project, _ := s.store.GetProject(id)

	var response = dto.ProjectResponse{
		project.Id,
		project.Name,
		project.Desc,
		project.UserId,
		nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s ProjectsService) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	project, err := s.store.GetProject(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if project == nil {
		w.Write([]byte("Project not found"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.store.DeleteProject(project.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Project deleted"))
}

func (s ProjectsService) GetProjectsByUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var token = r.Header.Get("Authorization")
	claims, err := users.ValidJWT(token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		var response = dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	projects, err := s.store.GetProjectsByUser(claims.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var response = dto.GetProjectsByUserResponse{
		projects,
	}

	json.NewEncoder(w).Encode(response)
}
