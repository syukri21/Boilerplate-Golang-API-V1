package DB

import (
	"asyifa-backend/src/Models"
)

// Migrate ...
func (d *DatabaseType) Migrate() *DatabaseType {
	// add Models to migrate
	d.db.AutoMigrate(
		Models.Customer{},
		Models.User{},
	)

	return d
}
