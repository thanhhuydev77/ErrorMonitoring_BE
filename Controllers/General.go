package Controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main.go/General"
	"main.go/Models"
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

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	filter := query.Get("filter")
	if len(filter) == 0 {
		fmt.Println("filters not present")
	}
	filterType := query.Get("type")

	switch filterType {
	case "project":
		SearchInProject(w, r)
		return
	case "user":
		SearchInUser(w, r)
		return
	default:
		result := General.CreateResponse(0, `invalid type!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
}

func InitAllController(r *mux.Router) {
	//General Controller
	r.Handle("/search", AuthMW(http.HandlerFunc(Search))).Methods("PUT")

	//Users Controller
	r.HandleFunc("/user", UserRequest).Methods("POST")
	r.Handle("/user", AuthMW(http.HandlerFunc(UserRequest))).Methods("PUT")
	r.Handle("/user", AuthMW(http.HandlerFunc(authenUser))).Methods("GET")
	r.Handle("/user/{Id}", AuthMW(http.HandlerFunc(GetUserByProjectId))).Methods("GET")
	r.Handle("/user/search", AuthMW(http.HandlerFunc(SearchInUser))).Methods("PUT")
	//Projects Controller
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("POST")
	r.Handle("/project", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project/{Id}", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("PUT")
	r.Handle("/project/search", AuthMW(http.HandlerFunc(SearchInProject))).Methods("PUT")

}
