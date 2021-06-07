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

func SuiteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	suiteRequest := Models.SuiteRequest{}
	err1 := json.NewDecoder(r.Body).Decode(&suiteRequest)

	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if projectId == "" {
		result := General.CreateResponse(0, `Create Suite failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if suiteRequest.Type == "create-suite" {
		CreateOK := Business.CreateSuite(projectId, suiteRequest.Suite)

		if !CreateOK {
			result := General.CreateResponse(0, `Create Suite failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}

		result := General.CreateResponse(1, `Create Suite successfully!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
}
func FilterInSuite(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	suiteFilter := Models.SuiteFilter{}
	err1 := json.NewDecoder(r.Body).Decode(&suiteFilter)

	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	List, err := Business.FilterSuite(suiteFilter)
	if err != nil || len(List) == 0 {
		result := General.CreateResponse(0, `Filter Suite failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	result := General.CreateResponse(1, `Filter Suite success!`, List)
	io.WriteString(w, result)
	return
}
func GetSuite(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	Id := vars["Id"]

	if projectId == "" {
		result := General.CreateResponse(0, `Get Suite failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	suite, GetOK := Business.GetSuite(projectId, Id)

	if !GetOK {
		result := General.CreateResponse(0, `Get Suite failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	result := General.CreateResponse(1, `Get Suite successfully!`, suite)
	io.WriteString(w, result)
	return

}
