package migrate

import (
	"github.com/amirhnajafiz/caaas/internal/model"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Migrate handler is used for creating database schemas.
type Handler struct {
	Database *pg.DB
}

func (h Handler) Execute() error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.UserGroup)(nil),
	}

	for _, model := range models {
		err := h.Database.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
