package main

import (
	"fmt"

	"website/routes"
	"website/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Version   string
	Buildtime string
)

// @title           website
// @version         0.0.0
func main() {
	errenv := godotenv.Load(".env")

	if errenv != nil {
		fmt.Println("Failed to read env variable, but will continue")
	}

	app := gin.Default()
	routes.Setup(app)

	listeningPort := utils.GetEnv("PORT", "8081")
	url := fmt.Sprintf("localhost:%s", listeningPort)
	fmt.Printf("application running in %s at %s\n", listeningPort, url)
	err := app.Run(url)
	if err != nil {
		fmt.Println("failed to start server")
	}
	fmt.Println("exit")
}
