package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rafael-japyassu/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"message\": \"Hello World\"}")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
