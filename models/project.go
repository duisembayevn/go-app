package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	UserId int    `json:"userId"`
}
