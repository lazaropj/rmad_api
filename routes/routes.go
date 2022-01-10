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

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	//r.Use(middleware.ContentTypeMiddleware)
	r.Use(app.JwtAuthentication)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	r.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("Post")
	r.HandleFunc("/api/user/login", controllers.Authenticate).Methods("Post")
	r.HandleFunc("/api/travel", controllers.GetTravelsFor).Methods("Get")
	r.HandleFunc("/api/travel", controllers.CreateTravel).Methods("Post")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
