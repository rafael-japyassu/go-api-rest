package main

import (
	"fmt"

	"github.com/rafael-japyassu/go-api-rest/database"
	"github.com/rafael-japyassu/go-api-rest/models"
	"github.com/rafael-japyassu/go-api-rest/routes"
)

func main() {
	models.Personalities = append(models.Personalities, models.Personality{Id: 1, Name: "Rafael", History: "History 1..."})
	models.Personalities = append(models.Personalities, models.Personality{Id: 2, Name: "Rafael", History: "History 2..."})
	models.Personalities = append(models.Personalities, models.Personality{Id: 3, Name: "Rafael", History: "History 3..."})

	database.ConnectionDb()

	fmt.Println("Api Rest Go")

	routes.HandleRequest()
}
