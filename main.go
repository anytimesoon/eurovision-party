package main

import (
	db "eurovision/db"
	"eurovision/pkg/routes"
	"log"
)

func init() {
	db.Start()
	log.Println("Database initialization complete")
}

func main() {
	log.Println("Starting server")
	routes.Start()

	db.Conn.Close()
}
