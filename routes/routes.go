package routes

import (
	"log"
	"net/http"

	"github.com/rafael-japyassu/go-api-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetPersonalities)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
