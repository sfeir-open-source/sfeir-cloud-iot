package main

import (
	"github.com/sfeir-cloud-iot/bicycle-api/router"

	_ "github.com/joho/godotenv/autoload"
)

// @Title bicycle API
// @Version 0.1.0
// @BasePath /v1
func main() {
	router.CreateRouter()
}
