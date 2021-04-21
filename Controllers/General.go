package Controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Server CSS, JS & Images Statically.
	router.
		PathPrefix("/public/").
		Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))
	return router
}

func InitAllController(r *mux.Router) {
	//Users Controller
	r.HandleFunc("/user", UserRequest).Methods("POST")
	r.Handle("/user", AuthMW(http.HandlerFunc(UserRequest))).Methods("PUT")
	r.Handle("/user", AuthMW(http.HandlerFunc(GetUser))).Methods("GET")

	//Projects Controller
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("POST")
	r.Handle("/project", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project/{Id}", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("PUT")

}
