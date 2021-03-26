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

//func InitAllController(a ApiDB, r *mux.Router, storage *c.Storage) {

//UsersController
//r.HandleFunc("/user/login", a.UserLogin).Methods("POST")
//r.HandleFunc("/user/register", a.UserRegister).Methods("POST")
//r.Handle("/user/get-all-username", c.Cached(storage, "10s", a.GetallUserName)).Methods("GET")
//r.Handle("/user/validate", AuthMW(http.HandlerFunc(ValidateToken))).Methods("POST")
//r.Handle("/user/get-user", AuthMW(http.HandlerFunc(a.GetUser))).Methods("GET")
//r.Handle("/user/get-user/{Id}", AuthMW(http.HandlerFunc(a.GetUser))).Methods("GET")
//}
