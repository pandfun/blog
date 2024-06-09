package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pandfun/blog/cmd/api"
	"github.com/pandfun/blog/config"
	"github.com/pandfun/blog/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("Error opening database %v", err)
	}

	initStorage(db)
	server := api.NewAPIServer(config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	log.Println("DB: Connected Successfully!")
}
