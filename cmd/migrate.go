package cmd

import (
	"github.com/amirhnajafiz/authX/internal/storage"
)

// Migrate command.
type Migrate struct{}

// main function of Migrate command.
func (m Migrate) main() {
	// open db
	_, err := storage.NewConnection(storage.Config{})
	if err != nil {
		panic(err)
	}
}
