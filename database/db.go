package database

import (
	"log"

	"github.com/arthuravila26/go-api-with-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Error ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
