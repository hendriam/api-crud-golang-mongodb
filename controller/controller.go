package controller

import (
	"books/lib"
	"books/models"
)

type App struct {
	logging lib.Logging
	model   models.Model
}

func New(db lib.Database) App {
	return App{
		logging: lib.LoadLogging(),
		model:   models.New(db),
	}
}
