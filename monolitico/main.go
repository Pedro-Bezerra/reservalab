package main

import (
	"github.com/Pedro-Bezerra/reservalab-mono/config"
	"github.com/Pedro-Bezerra/reservalab-mono/database"
	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConectarDB()

	database.DB.AutoMigrate(&entity.Solicitacao{}, &entity.Mensagem{}, &entity.Reserva{})

	router := gin.Default()

	router.Run(":8080")
}
