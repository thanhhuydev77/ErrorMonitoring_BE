package Controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"main.go/Business"
	"main.go/General"
	"main.go/Models"
	"net/http"
)

func ProjectRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	projectRequest := Models.ProjectRequest{}
	err1 := json.NewDecoder(r.Body).Decode(&projectRequest)
	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	switch projectRequest.Type {
	case "create-project":
		RegisterOK := false
		if len(projectRequest.Project.Name) > 0 {
			RegisterOK, _ = Business.CreateProject(projectRequest.Project)

		}
		if !RegisterOK {
			result := General.CreateResponse(0, `Create project failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}

		result := General.CreateResponse(1, `Create project successfully!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	case "set-status":
		ChangeStatusOK := false
		if len(projectRequest.Project.Id) > 0 {
			ChangeStatusOK = Business.ChangeStatusProject(projectRequest.Project)
		}
		if !ChangeStatusOK {
			result := General.CreateResponse(0, `Change status failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
		result := General.CreateResponse(1, `Change status success!`, Models.EmptyObject{})
		io.WriteString(w, result)

		return
	}

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	Id := vars["Id"]

	token := r.URL.Query().Get("token")
	email := GetEmailFromToken(token)
	log.Print("email :" + email)
	List, err := Business.GetProjects(email, Id)
	if err != nil {

		result := General.CreateResponse(0, `Get project failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if len(Id) > 0 {
		result := General.CreateResponse(1, `Get project success!`, List[0])
		io.WriteString(w, result)
		return
	}
	result := General.CreateResponse(1, `Get list project success!`, List)
	io.WriteString(w, result)
	return
}
