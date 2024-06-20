package controller

import "github.com/go-pg/pg/v10"

// Controller is responsible for making database connections.
type Controller struct {
	database *pg.DB
}

func NewController(db *pg.DB) *Controller {
	return &Controller{
		database: db,
	}
}
