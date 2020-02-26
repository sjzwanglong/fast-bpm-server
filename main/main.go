package main

import (
	"fast-bpm/db"
	"fast-bpm/server"
)

func main() {
	db.InitFactory("mysql")
	server.InitServer()
	server.StartServer()
}
