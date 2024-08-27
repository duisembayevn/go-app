package dto

import (
	"go_app/models"
)

type ProjectRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	UserId int    `json:"userId"`
}

type ProjectResponse struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	UserId int     `json:"userId"`
	Tasks  *[]Task `json:"tasks"`
}

type GetProjectsByUserResponse struct {
	Projects []models.Project `json:"projects"`
}

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
