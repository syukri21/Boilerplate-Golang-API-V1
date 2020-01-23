package main

import (
	Bootstrap "asyifa-backend/src/Bootstraps"
	"asyifa-backend/src/Config"
	"flag"
	"fmt"
)

func main() {
	useMigrate := flag.Bool("migrate", false, "Migrating...")
	flag.Parse()

	if *useMigrate {
		fmt.Println("You are migrating...")
		Bootstrap.Aplication.Migrate()
		fmt.Println("Migrating Done...")

	} else {
		fmt.Println("Listen on ", Config.HostPort)
		Bootstrap.Aplication.Run()
	}

}
