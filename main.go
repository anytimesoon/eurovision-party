package main

import (
	db "eurovision/db"
	"eurovision/pkg/routes"
	"mime"
)

func init() {
	db.Start()
}

func main() {
	mime.AddExtensionType(".js", "application/javascript")

	routes.Start()

	db.Conn.Close()
}
