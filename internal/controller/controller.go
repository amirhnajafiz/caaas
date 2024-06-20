package controller

import "github.com/go-pg/pg/v10"

type Controller struct {
	database *pg.DB
}
