package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lazaropj/rmad_api/models"
	u "github.com/lazaropj/rmad_api/utils"
)

type ElectionPost struct {
	Note int64
	Code string
}

func VoteOnTravel(w http.ResponseWriter, r *http.Request) {
	var electionPost ElectionPost
	err := json.NewDecoder(r.Body).Decode(&electionPost)

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body: "+err.Error()))
		return
	}

	travel := models.GetTravel("code", electionPost.Code)

	if travel == nil {
		u.Respond(w, u.Message(false, "Travel not found by code: "+electionPost.Code))
		return
	}

	account := models.GetUser(r.Context().Value("accountId").(uint))

	models.VoteOnTravel(*travel, *account, electionPost.Note)

	resp := u.Message(true, "success")

	u.Respond(w, resp)

}

func GetAverageByTravel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	travelId := vars["travelId"]

	if travelId == "" {
		u.Respond(w, u.Message(false, "TravelId is required"))
		return
	}

	average := models.GetAverageByTravel(travelId)

	if average == 0 {
		u.Respond(w, u.Message(false, "Average is zero for travelId: "+travelId))
		return
	}

	resp := u.Message(true, "success")
	resp["average"] = average

	u.Respond(w, resp)
}
