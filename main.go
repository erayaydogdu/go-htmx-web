package main

import (
	"fmt"
	"go-htmx-web/model"
	"go-htmx-web/routes"
)

func main() {

	fmt.Printf("Starting server...\n")
    model.Setup()
	fmt.Printf("Registering routes...\n")
    routes.SetupAndRun()
	fmt.Printf("Server started!\n")
}
