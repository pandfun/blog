package database

import (
	"blog/config"
	"blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := config.Envs.DBUser + ":" + config.Envs.DBPassword + "@tcp(" + config.Envs.DBAddress + ")/" + config.Envs.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Database connection failed.")
	}

	log.Println("DB: Connected to database!")

	db.AutoMigrate(new(model.Blog))

	DBConn = db
}
