package main

import (
	"github.com/Pedro-Bezerra/mensagem/config"
	"github.com/Pedro-Bezerra/mensagem/database"
	"github.com/Pedro-Bezerra/mensagem/entity"
	"github.com/Pedro-Bezerra/mensagem/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()

	database.ConectarDB()

	database.DB.AutoMigrate(&entity.Mensagem{})
}

func main() { 
	router := gin.Default()

	routes.RotasMensagem(router)

	router.Run(":8081")
}
