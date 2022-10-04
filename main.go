package main

import (
	"azureDataLineage/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", controllers.Login)
	r.POST("/graph", controllers.Graph)
	r.Run(":8088")

}
