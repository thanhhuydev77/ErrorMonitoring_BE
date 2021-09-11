package Controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"main.go/Business"
	"main.go/General"
	"main.go/Models"
	"net/http"
)

func IssueRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	issueRequest := Models.IssueRequest{}
	err1 := json.NewDecoder(r.Body).Decode(&issueRequest)

	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if projectId == "" {
		result := General.CreateResponse(0, `Create Issue failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	switch issueRequest.Type {
	case "create-issue":
		CreateOK, ErrMsg := Business.CreateIssue(projectId, issueRequest.Issue)

		if !CreateOK {
			result := General.CreateResponse(0, `Create Issue failed!`, ErrMsg)
			io.WriteString(w, result)
			return
		}
		//mail to assignee,manager,admin

		result := General.CreateResponse(1, `Create Issue successfully!`, ErrMsg)
		io.WriteString(w, result)
		break
	case "update-issue":
		UpdateOK := Business.UpdateIssue(projectId, issueRequest.Issue)

		if !UpdateOK {
			result := General.CreateResponse(0, `Update Issue failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}

		result := General.CreateResponse(1, `Update Issue successfully!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
		break
	case "update-reviewer":
		UpdateOK := Business.UpdateIssueReviewer(projectId, issueRequest.Issue)

		if !UpdateOK {
			result := General.CreateResponse(0, `Update Issue Reviewer failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}

		result := General.CreateResponse(1, `Update Issue Reviewer successfully!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
		break
	}

}

func FilterInIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	issueFilter := Models.IssueFilter{}
	err1 := json.NewDecoder(r.Body).Decode(&issueFilter)

	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	List, err := Business.FilterIssue(issueFilter)
	if err != nil || len(List) == 0 {
		result := General.CreateResponse(0, `Filter Issue failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	List = Models.SortIssueByCreateTime(List)
	result := General.CreateResponse(1, `Filter Issue success!`, List)
	io.WriteString(w, result)
	return
}

func GetIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	Id := vars["Id"]

	if projectId == "" {
		result := General.CreateResponse(0, `Get Issue failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	issue, GetOK := Business.GetIssue(projectId, Id)

	if !GetOK {
		result := General.CreateResponse(0, `Get Issue failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	result := General.CreateResponse(1, `Get Issue successfully!`, issue)
	io.WriteString(w, result)
	return

}
