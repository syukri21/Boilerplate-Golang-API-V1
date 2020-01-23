package Bootstraps

import (
	"asyifa-backend/src/Controllers"
	"asyifa-backend/src/Utils"
)

func (v *aplication) handleRoutes() *aplication {

	// Default
	d := Controllers.Default{
		DB:            v.db.GetDB(),
		ResponseJSON:  Utils.ResponseJSON,
		ResponseError: Utils.ResponseError,
	}

	// Users
	users := Controllers.Users{Default: d}
	v.Router.HandleFunc("/v1/api/users", users.GetAll)

	return v
}
