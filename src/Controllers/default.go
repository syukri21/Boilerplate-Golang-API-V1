package Controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// Default ...
type Default struct {
	DB            *gorm.DB
	ResponseJSON  func(w http.ResponseWriter, status int, payload interface{})
	ResponseError func(w http.ResponseWriter, code int, message string)
}
