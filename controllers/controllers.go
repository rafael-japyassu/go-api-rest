package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafael-japyassu/go-api-rest/database"
	"github.com/rafael-japyassu/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"message\": \"Hello World\"}")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality

	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	err := json.NewDecoder(r.Body).Decode(&personality)

	if err != nil {
		log.Fatal("Body inválido")
	}

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.Delete(&personality, id)

	w.WriteHeader(204)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)

	w.WriteHeader(204)
}
