package main

import (
	"fmt"

	"github.com/ozonva/ova-rule-api/configs"
)

func main() {
	fmt.Println("Awesome service ova-rule-api")

	configs.Load()

	fmt.Printf("Server config: %+v\n", *configs.ServerConfig)
	fmt.Printf("Database config: %+v\n", *configs.DatabaseConfig)
}
