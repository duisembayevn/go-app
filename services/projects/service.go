package projects

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go_app/models"
	"log"
	"net/http"
	"strconv"
)

// Response models

type CreateProjectResponse struct {
	Id int `json:"id"`
}

type GetProjectByIdResponse struct {
	Project *models.Project `json:"project"`
}

type GetProjectsByUserIdResponse struct {
	Projects []models.Project `json:"projects"`
}

type ProjectsService struct {
	db *sql.DB
}

func NewProjectsService(db *sql.DB) *ProjectsService {
	return &ProjectsService{
		db: db,
	}
}

// Endpoint handlers

func (s ProjectsService) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := CreateProject(s.db, project.Name, project.Desc, project.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response = CreateProjectResponse{
		Id: id,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s ProjectsService) GetProjectByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	project, _ := GetProjectById(s.db, id)

	var response = GetProjectByIdResponse{
		project,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s ProjectsService) GetProjectsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	projects, err := GetProjectsByUserId(s.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var response = GetProjectsByUserIdResponse{
		projects,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s ProjectsService) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	project, err := GetProjectById(s.db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(project)
}

// DB Helpers

func CreateProject(db *sql.DB, name string, desc string, userId int) (int, error) {
	var query = "Insert into `projects` (`name`, `desc`, `userId`) values (?, ?, ?)"
	result, err := db.Exec(query, name, desc, userId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetProjectsByUserId(db *sql.DB, userId int) ([]models.Project, error) {
	var projects []models.Project
	var query = "select * from projects where `userId` = ?"
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.Id, &project.Name, &project.Desc, &project.UserId)
		if err != nil {
			log.Fatal(err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func GetProjectById(db *sql.DB, id int) (*models.Project, error) {
	project := models.Project{}
	var query = "select * from projects where `id` = ?"
	row := db.QueryRow(query, id)
	err := row.Scan(&project.Id, &project.Name, &project.Desc, &project.UserId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
