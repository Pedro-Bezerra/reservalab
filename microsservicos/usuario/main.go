package main

import (
	"github.com/Pedro-Bezerra/usuario/config"
	"github.com/Pedro-Bezerra/usuario/database"
	"github.com/Pedro-Bezerra/usuario/entity"
	"github.com/Pedro-Bezerra/usuario/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()

	database.ConectarDB()

	database.DB.AutoMigrate(&entity.Usuario{})
}

func main() { 
	router := gin.Default()

	routes.RotasMonitor(router)

	router.Run(":8082")
}
