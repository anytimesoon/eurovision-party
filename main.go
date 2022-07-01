package main

import (
	db "eurovision/db"
	"eurovision/pkg/routes"
)

func init() {
	db.Start()
}

func main() {
	routes.Start()

	db.Conn.Close()
}
