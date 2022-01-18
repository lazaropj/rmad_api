package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	app "github.com/lazaropj/rmad_api/app"
	"github.com/lazaropj/rmad_api/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Port: " + port)

	//r.Use(middleware.ContentTypeMiddleware)
	r.Use(app.JwtAuthentication)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("Post")
	r.HandleFunc("/api/user/login", controllers.Authenticate).Methods("Post")
	r.HandleFunc("/api/travel", controllers.GetTravelsFor).Methods("Get")
	r.HandleFunc("/api/travel/finish/{id}", controllers.FinishTravel).Methods("Put")
	r.HandleFunc("/api/travel", controllers.CreateTravel).Methods("Post")
	r.HandleFunc("/api/travel/vote", controllers.VoteOnTravel).Methods("Post")
	r.HandleFunc("/api/election/average/{travelId}", controllers.GetAverageByTravel).Methods("Get")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
