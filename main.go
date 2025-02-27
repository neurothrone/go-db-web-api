package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"neurothrone/go-db-web-api/data"
	"neurothrone/go-db-web-api/endpoints"
	"neurothrone/go-db-web-api/settings"
)

var config settings.Config

func main() {
	settings.ReadConfig(&config)

	fmt.Println("Config:")
	fmt.Println(config.Database.File)

	data.Init(
		config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port,
	)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")

	router.GET("/", endpoints.Index)
	router.POST("/api/product", endpoints.AddProduct)
	router.GET("/api/product", endpoints.GetAllProducts)
	router.GET("/api/product/:id", endpoints.GetProductById)
	router.PUT("/api/product/:id", endpoints.UpdateProduct)
	router.DELETE("/api/product/:id", endpoints.DeleteProduct)

	// Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//err := router.Run(":8080")
	err := router.Run()
	if err != nil {
		return
	}
}
