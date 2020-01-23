package Bootstraps

import (
	"asyifa-backend/src/Config"

	"asyifa-backend/src/DB"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type aplication struct {
	Host   int
	Router *mux.Router
	db     *DB.DatabaseType
}

func (v *aplication) listen() *aplication {
	http.ListenAndServe(":"+strconv.Itoa(v.Host), v.Router)
	return v
}

// Run = function for Run aplication
func (v *aplication) Run() {
	v.db.Open()
	log.Fatal(v.handleRoutes().listen())
}

// Migrate ...
func (v *aplication) Migrate() {
	v.db.Open().Migrate().Close()
}

// Aplication = defined aplication
var Aplication = aplication{
	Host:   Config.HostPort,
	Router: mux.NewRouter().StrictSlash(true),
	db:     &DB.Database,
}
