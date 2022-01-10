package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lazaropj/rmad_api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	models.GetDB().Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	models.GetDB().First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	models.GetDB().Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	models.GetDB().Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	models.GetDB().Model(&personalidade).Where("id = ?", id).Updates(personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
