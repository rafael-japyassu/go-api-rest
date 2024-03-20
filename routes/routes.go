package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rafael-japyassu/go-api-rest/controllers"
	"github.com/rafael-japyassu/go-api-rest/middlewares"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middlewares.SetContentTypeMiddleware)

	const route = "/api/personalities/{id}"

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc(route, controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc(route, controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc(route, controllers.EditPersonality).Methods("PUT")

	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)

	log.Fatal(http.ListenAndServe(":8080", cors))
}
