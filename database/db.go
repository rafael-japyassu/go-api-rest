package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDb() {

	urlConnection := "host=localhost user=docker password=docker dbname=db_go port=5435 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(urlConnection), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conector com o banco de dados")
	}
}
