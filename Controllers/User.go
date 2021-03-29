package Controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/gorilla/mux"
	"io"
	"main.go/Business"
	"main.go/Models"
	"net/http"
	"time"
)

//login with username and pass from body request
func UserLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	p := Models.User{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	IsExsist, passok, user := Business.Login(nil, p.UserName, p.PassWord)

	if !IsExsist {
		var failRes Models.RespondFail
		failRes.Message = `Can't find user please sign in again!`
		result, _ := json.Marshal(failRes)
		//w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, string(result))
		return
	}

	if !passok {
		//w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Your password is wrong, please type again !"}`)
		return
	}
	// expired after 1000 dates
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.UserName,
		"exp":  time.Now().Add(time.Hour * time.Duration(1000*24)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(Models.AppConfig.AppKey))
	//tokenString += ""
	//w.Write([]byte(`{"hello": "world"}`))
	//stringresult := `{"OK" }`
	io.WriteString(w, tokenString)
	return
}

//register a new user
//func (a *ApiDB) UserRegister(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	p := MODELS.RequestRegister{}
//	err1 := json.NewDecoder(r.Body).Decode(&p)
//	if err1 != nil {
//		io.WriteString(w, `{"message": "wrong format!"}`)
//		return
//	}
//
//	_, err := BUSINESS.Register(a.Db, p)
//	if err == nil {
//		io.WriteString(w, `{
//			 	"status": 200,
//				"message":"Register success",
//				"data": {
//					"status": 1
//				}
//			}`)
//	} else {
//		io.WriteString(w, `{"message":"Register fail"}`)
//	}
//}

//get all user names
//func (a *ApiDB) GetallUserName(w http.ResponseWriter, r *http.Request) {
//	// Query()["key"] will return an array of items,
//	// we only want the single item.
//
//	allusername := BUSINESS.GetAllUserName(a.Db)
//	w.Header().Add("Content-Type", "application/json")
//	if allusername == nil {
//		//w.WriteHeader(http.StatusBadRequest)
//		io.WriteString(w, `{"message":"get all username unsuccess"}`)
//		return
//	}
//	type result struct {
//		Message string   `json:"message"`
//		Data    []string `json:"data"`
//	}
//	Result, _ := json.Marshal(result{Message: "get all username success", Data: allusername})
//	//w.WriteHeader(200)
//	io.WriteString(w, string(Result))
//}
//
////get a user with id or all user(id =-1)
//func (a *ApiDB) GetUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	vars := mux.Vars(r)
//	Id, err := strconv.Atoi(vars["Id"])
//	//have not id --> get all
//	if err != nil {
//		Id = -1
//	}
//	List := BUSINESS.GetUsers(a.Db, Id)
//	jsonlist, _ := json.Marshal(List)
//	result := List[0]
//
//	stringresult := `{"message": "Get Users success","status": 200,"data":{"user":`
//	if len(List) == 1 {
//		jsonresult, _ := json.Marshal(result)
//		stringresult += string(jsonresult)
//	} else {
//		stringresult += string(jsonlist)
//	}
//	stringresult += "}}"
//	io.WriteString(w, stringresult)
//}
