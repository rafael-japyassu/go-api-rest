package main

import (
	"fmt"

	"github.com/rafael-japyassu/go-api-rest/models"
	"github.com/rafael-japyassu/go-api-rest/routes"
)

func main() {
	models.Personalities = append(models.Personalities, models.Personality{Name: "Rafael", History: "History 1..."})
	models.Personalities = append(models.Personalities, models.Personality{Name: "Rafael", History: "History 2..."})
	models.Personalities = append(models.Personalities, models.Personality{Name: "Rafael", History: "History 3..."})

	fmt.Println("Api Rest Go")

	routes.HandleRequest()
}
