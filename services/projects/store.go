package projects

import (
	"go_app/models"
)

type ProjectsStore interface {
	CreateProject(name string, desc string, userId int) (int, error)
	GetProjectsByUser(id int) ([]models.Project, error)
	GetProject(id int) (*models.Project, error)
	DeleteProject(id int) error
}
