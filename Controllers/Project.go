package Controllers

import (
	"encoding/json"
	"fmt"
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
	case "add-member":
		AddMemberOk := false
		if len(projectRequest.Project.Id) > 0 {
			AddMemberOk = Business.AddMember("", projectRequest.Project)
		}
		if !AddMemberOk {
			result := General.CreateResponse(0, `Add member failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return

		}
		result := General.CreateResponse(1, `Add member success!`, Models.EmptyObject{})
		io.WriteString(w, result)

		return
	case "modify-member":
		ModifyMemberOk := false
		if len(projectRequest.Project.Id) > 0 {
			ModifyMemberOk = Business.ModifyMember("", projectRequest.Project)
		}
		if !ModifyMemberOk {
			result := General.CreateResponse(0, `Mofidy member failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return

		}
		result := General.CreateResponse(1, `Mofidy member success!`, Models.EmptyObject{})
		io.WriteString(w, result)

		return
	case "remove-member":
		RemoveMemberOk := false
		if len(projectRequest.Project.Id) > 0 {
			RemoveMemberOk = Business.RemoveMember("", projectRequest.Project)
		}
		if !RemoveMemberOk {
			result := General.CreateResponse(0, `Remove member failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return

		}
		result := General.CreateResponse(1, `Remove member success!`, Models.EmptyObject{})
		io.WriteString(w, result)

		return
	case "update-Trello":
		{
			ChangeTrelloInfoOK := false
			if len(projectRequest.Project.Id) > 0 {
				ChangeTrelloInfoOK = Business.UpdateTrelloInfo(projectRequest.Project)
			}
			if !ChangeTrelloInfoOK {
				result := General.CreateResponse(0, `Change Trello Info failed!`, Models.EmptyObject{})
				io.WriteString(w, result)
				return
			}
			result := General.CreateResponse(1, `Change Trello Info success!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
	case "update-Slack":
		{
			ChangeTrelloInfoOK := false
			if len(projectRequest.Project.Id) > 0 {
				ChangeTrelloInfoOK = Business.UpdateSlackInfo(projectRequest.Project)
			}
			if !ChangeTrelloInfoOK {
				result := General.CreateResponse(0, `Change Slack Info failed!`, Models.EmptyObject{})
				io.WriteString(w, result)
				return
			}
			result := General.CreateResponse(1, `Change Slack Info success!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
	case "update-Auto-Suggest":
		{
			ChangeTrelloInfoOK := false
			if len(projectRequest.Project.Id) > 0 {
				ChangeTrelloInfoOK = Business.UpdateAutoSuggest(projectRequest.Project)
			}
			if !ChangeTrelloInfoOK {
				result := General.CreateResponse(0, `Change Auto Suggest failed!`, Models.EmptyObject{})
				io.WriteString(w, result)
				return
			}
			result := General.CreateResponse(1, `Change Auto Suggest success!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
	case "update-Mail-Notification":
		{
			ChangeTrelloInfoOK := false
			if len(projectRequest.Project.Id) > 0 {
				ChangeTrelloInfoOK = Business.UpdateAutoSentMail(projectRequest.Project)
			}
			if !ChangeTrelloInfoOK {
				result := General.CreateResponse(0, `Change Mail Notification failed!`, Models.EmptyObject{})
				io.WriteString(w, result)
				return
			}
			result := General.CreateResponse(1, `Change Mail Notification success!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
	}

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	Id := vars["Id"]

	token := r.URL.Query().Get("token")
	email := General.GetEmailFromToken(token)
	log.Print("email :" + email)
	List, err := Business.GetProjects(email, Id)
	if err != nil {
		result := General.CreateResponse(0, `Get project(s) failed!`, Models.EmptyObject{})
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

func SearchInProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	filter := query.Get("filter")
	if len(filter) == 0 {
		fmt.Println("filters not present")
	}

	List, err := Business.SearchProject(filter)
	if err != nil || len(List) == 0 {
		result := General.CreateResponse(0, `Search projects failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	result := General.CreateResponse(1, `Search projects success!`, List)
	io.WriteString(w, result)
	return
}
