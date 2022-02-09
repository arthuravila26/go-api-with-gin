package main

import (
	"github.com/arthuravila26/go-api-with-gin/models"
	"github.com/arthuravila26/go-api-with-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "4800000000"},
	}
	routes.HandleRequests()
}
