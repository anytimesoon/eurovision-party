package main

import (
	db "eurovision/db"
	"eurovision/pkg/routes"
	"mime"
)

func init() {
	db.Connect()
}

func main() {
	mime.AddExtensionType(".js", "application/javascript")

	routes.Start()

	db.Conn.Close()
}
