package main

import (
	"github.com/Pedro-Bezerra/reserva/config"
	"github.com/Pedro-Bezerra/reserva/database"
	"github.com/Pedro-Bezerra/reserva/entity"
	"github.com/Pedro-Bezerra/reserva/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()

	database.ConectarDB()

	database.DB.AutoMigrate(&entity.Solicitacao{}, &entity.Reserva{})
}

func main() { 
	router := gin.Default()

	routes.RotasSolicitacao(router)

	router.Run(":8080")
}
