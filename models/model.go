package models

import (
	"books/lib"
	"context"
	"time"
)

type Model struct {
	db lib.Database
}

const defaultTimeout = 60 * time.Second

func New(db lib.Database) Model {
	return Model{db: db}
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
