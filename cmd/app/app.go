package app

import (
	"example/database"
	"example/repo"
	"example/service"

	"gorm.io/gorm"
)

var application Application

func Init() {
	db := database.NewDB()
	re := repo.NewRepo(db)
	se := service.NewService(re)
	application.db = db
	application.service = se
	application.repo = re

}

type Application struct {
	db      *gorm.DB
	service service.Service
	repo    repo.Repository
}
