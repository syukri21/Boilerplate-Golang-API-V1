package DB

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// import mysql
	_ "github.com/go-sql-driver/mysql"
)

func (d *DatabaseType) mysql() *DatabaseType {
	d.db, d.Error = gorm.Open(d.Dialect, d.User+":"+d.Password+"@/"+d.DbName+"?charset=utf8&parseTime=True&loc=Local")
	if d.Error != nil {
		fmt.Println("error")
	}
	return d
}
