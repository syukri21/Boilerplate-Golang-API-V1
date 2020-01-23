package DB

import (
	"asyifa-backend/src/Config"

	"github.com/jinzhu/gorm"
)

// DatabaseType ...
type DatabaseType struct {
	Dialect  string
	User     string
	Password string
	DbName   string
	db       *gorm.DB
	Error    error
}

// Open ...
func (d *DatabaseType) Open() *DatabaseType {
	switch d.Dialect {
	case "mysql":
		return d.mysql()
	default:
		return d.mysql()
	}
}

// Close ...
func (d *DatabaseType) Close() {
	d.db.Close()
}

// GetDB ...
func (d *DatabaseType) GetDB() *gorm.DB {
	return d.db
}

// Database ...
var Database = DatabaseType{
	Dialect:  Config.DB.Dialect,
	User:     Config.DB.User,
	Password: Config.DB.Password,
	DbName:   Config.DB.DbName,
}
