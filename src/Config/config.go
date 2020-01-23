package Config

// HostPort host to Listen
const HostPort = 2020

// DB Configuration
var DB = db{
	Dialect:  "mysql",
	User:     "root",
	Password: "",
	DbName:   "ecommerce",
}
