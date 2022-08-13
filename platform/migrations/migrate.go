package migrations

import (
	"github.com/grrkaan/go-rest-api/app/models"
	"github.com/grrkaan/go-rest-api/platform/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.Post{})
}
