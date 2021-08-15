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

func Filter(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	filterType := query.Get("type")

	switch filterType {
	case "issue":
		FilterInIssue(w, r)
		return
	case "suite":
		FilterInSuite(w, r)
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
	r.Handle("/filter", AuthMW(http.HandlerFunc(Filter))).Methods("PUT")

	//User Controllers
	r.HandleFunc("/user", UserRequest).Methods("POST")
	r.Handle("/user", AuthMW(http.HandlerFunc(UserRequest))).Methods("PUT")
	r.Handle("/user", AuthMW(http.HandlerFunc(authenUser))).Methods("GET")
	r.Handle("/user/{Id}", AuthMW(http.HandlerFunc(GetUserByProjectId))).Methods("GET")
	r.Handle("/user/search", AuthMW(http.HandlerFunc(SearchInUser))).Methods("PUT")
	//Project Controllers
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("POST")
	r.Handle("/project", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project/{Id}", AuthMW(http.HandlerFunc(GetProject))).Methods("GET")
	r.Handle("/project", AuthMW(http.HandlerFunc(ProjectRequest))).Methods("PUT")
	r.Handle("/project/search", AuthMW(http.HandlerFunc(SearchInProject))).Methods("PUT")

	//Issue Controllers
	r.HandleFunc("/issue/{projectId}", IssueRequest).Methods("POST")
	r.Handle("/issue/{projectId}", AuthMW(http.HandlerFunc(IssueRequest))).Methods("PUT")
	r.Handle("/issue/{projectId}/{Id}", AuthMW(http.HandlerFunc(GetIssue))).Methods("GET")

	//Automation test suites Controllers
	r.HandleFunc("/suite/{projectId}", SuiteRequest).Methods("POST")
	r.Handle("/suite/{projectId}/{Id}", AuthMW(http.HandlerFunc(GetSuite))).Methods("GET")

	//Trello Integration
	r.HandleFunc("/TestTrello", TestTrello).Methods("POST")
	r.HandleFunc("/GetListBoard", GetListBoard).Methods("GET")

	//Slack Integration
	r.HandleFunc("/TestSlack", TestSlack).Methods("POST")
}
