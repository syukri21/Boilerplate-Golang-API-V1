package Controllers

import (
	"asyifa-backend/src/Models"
	"net/http"
)

// Users ...
type Users struct {
	Default
}

// GetAll ...
func (u *Users) GetAll(w http.ResponseWriter, r *http.Request) {
	users := []Models.User{}
	u.DB.Find(&users, 1)
	u.ResponseJSON(w, http.StatusOK, users)
}
