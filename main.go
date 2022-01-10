package main

import (
	"fmt"

	"github.com/lazaropj/rmad_api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go...")
	routes.HandleResquest()
}
