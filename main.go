package main

import (
	"fmt"

	"github.com/lazaropj/rmad_api/database"
	"github.com/lazaropj/rmad_api/models"
	"github.com/lazaropj/rmad_api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Mendes Leitão", Historia: "Primeiro subprefeito de São José dos Pinhais"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go...")
	routes.HandleResquest()
}
