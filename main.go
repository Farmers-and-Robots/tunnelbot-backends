package main

import (
	"fmt"
	"github.com/Farmers-and-Robots/tunnelbot-backends/models"
	"github.com/Farmers-and-Robots/tunnelbot-backends/router"
)

func main() {
	fmt.Printf("%v", models.Address{
		City:   "Dubuque",
		State:  "IA",
		Street: "401 S. Main St",
		Zip:    "52001",
	})

	router.NewRouter()
}
