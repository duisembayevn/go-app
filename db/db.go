package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"go_app/config"
	"log"
)

func ConnectDB(config config.Config) *sql.DB {
	db := OpenConnection(CreateMySQLConfig(config))
	CheckConnection(db)
	ApplyMigrations(db)

	return db
}

func CreateMySQLConfig(config config.Config) mysql.Config {
	return mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Addr:                 config.DBAddress,
		DBName:               config.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}

func OpenConnection(config mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CheckConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")
}

func ApplyMigrations(db *sql.DB) {
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "/Users/duisembayev/Desktop/go_app/cmd/migrate/migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations successfully applied!")
}
