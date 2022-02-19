package main

import (
	"github.com/arthuravila26/go-api-with-gin/database"
	"github.com/arthuravila26/go-api-with-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
