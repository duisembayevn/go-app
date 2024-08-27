package db

import (
	"database/sql"
	"go_app/models"
	"log"
)

type Store struct {
	database *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s Store) GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	query := "SELECT * FROM users WHERE `username`= ?"
	row := s.database.QueryRow(query, username)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s Store) CreateUser(username string, firstName string, lastName string, password string) error {
	query := "INSERT INTO `users` (`firstName`, `lastName`, `username`, `password`) VALUES (?, ?, ?, ?)"
	_, err := s.database.Exec(query, firstName, lastName, username, password)
	if err != nil {
		return err
	}
	return nil
}

// Projects Store

func (s Store) CreateProject(name string, desc string, userId int) (int, error) {
	var query = "Insert into `projects` (`name`, `desc`, `userId`) values (?, ?, ?)"
	result, err := s.database.Exec(query, name, desc, userId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s Store) GetProjectsByUser(id int) ([]models.Project, error) {
	var projects []models.Project
	var query = "select * from projects where `userId` = ?"
	rows, err := s.database.Query(query, id)
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

func (s Store) GetProject(id int) (*models.Project, error) {
	project := models.Project{}
	var query = "select * from projects where `id` = ?"
	row := s.database.QueryRow(query, id)
	err := row.Scan(&project.Id, &project.Name, &project.Desc, &project.UserId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s Store) DeleteProject(id int) error {
	var query = "delete from projects where `id` = ?"
	_, err := s.database.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
