package database

import (
	"log"
	"github.com/Pedro-Bezerra/mensagem/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConectarDB() {
	dsn := config.GetDBConnectionString()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados: %v", err)
	}

	DB = database
	log.Println("Banco de dados conectado com sucesso!")
}
