package migrate

import (
	"github.com/amirhnajafiz/caaas/internal/model"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Migrate handler is used for creating database schemas.
type Handler struct {
	database *pg.DB
}

func (h Handler) Execute() error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.UserGroup)(nil),
	}

	for _, model := range models {
		err := h.database.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
