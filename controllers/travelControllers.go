package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lazaropj/rmad_api/models"
	u "github.com/lazaropj/rmad_api/utils"
)

var CreateTravel = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("accountId").(uint) //Grab the id of the user that send the request
	travel := &models.Travel{}

	err := json.NewDecoder(r.Body).Decode(travel)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body: "+err.Error()))
		return
	}

	travel.UserId = user
	resp := travel.Create()
	u.Respond(w, resp)
}

var GetTravelsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("accountId").(uint)
	data := models.GetTravels(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
