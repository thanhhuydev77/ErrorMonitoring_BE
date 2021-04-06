package Controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io"
	"log"
	"main.go/Business"
	"main.go/GeneralFunction"
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

	IsExsist, passok := Business.Login(p.Email, p.PassWord)

	if !IsExsist {
		result := GeneralFunction.CreateResponse(0, `Can't find user please sign in again!`, `{}`)
		io.WriteString(w, result)
		return
	}

	if !passok {
		//w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Your password is wrong, please type again !"}`)
		return
	}
	// expired after 1000 dates
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": p.Email,
		"exp":  time.Now().Add(time.Hour * time.Duration(1000*24)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(Models.AppConfig.AppKey))
	io.WriteString(w, tokenString)
	return
}

//register a new user
func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user := Models.User{}
	err1 := json.NewDecoder(r.Body).Decode(&user)
	if err1 != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	_, errCode := Business.Register(user)
	if errCode == 0 {
		io.WriteString(w, `{
			 	"status": 200,
				"message":"Register success",
				"data": {
					"status": 1
				}
			}`)
	} else {
		io.WriteString(w, `{"message":"Register fail"}`)
	}
}

func UserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user := Models.UserRequest{}
	err1 := json.NewDecoder(r.Body).Decode(&user)
	if err1 != nil {
		result := GeneralFunction.CreateResponse(0, `Can't login, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if user.Type == "login" {

		IsExist, passOk := Business.Login(user.Email, user.PassWord)

		if (!passOk) || (!IsExist) {
			result := GeneralFunction.CreateResponse(0, `Email or password is incorrect, please try again`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}

		type data struct {
			Email string `json:"email"`
			Token string `json:"token"`
		}

		Data := data{
			Email: user.Email,
			Token: GenerateToken(user.Email),
		}
		result := GeneralFunction.CreateResponse(1, `Login success`, Data)
		io.WriteString(w, result)
		return
	}

	if user.Type == "register" {
		_, errCode := Business.Register(GeneralFunction.ConvertUserRequesttoUser(user))
		switch errCode {
		case 0:
			result := GeneralFunction.CreateResponse(1, `Register success!`, Models.EmptyObject{})
			io.WriteString(w, result)
		case 1:
			result := GeneralFunction.CreateResponse(0, `Email is already exists, please use another email.`, Models.EmptyObject{})
			io.WriteString(w, result)
		default:
			result := GeneralFunction.CreateResponse(0, `Register failed, please try again.`, Models.EmptyObject{})
			io.WriteString(w, result)
		}

		return
	}
}

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
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	Id := vars["Id"]

	List, err := Business.GetUsers(Id)
	if err != nil {
		log.Print(err.Error())
	}
	jsonlist, _ := json.Marshal(List)
	result := List[0]

	stringresult := `{"message": "Get Users success","status": 200,"data":{"user":`
	if len(List) == 1 {
		jsonresult, _ := json.Marshal(result)
		stringresult += string(jsonresult)
	} else {
		stringresult += string(jsonlist)
	}
	stringresult += "}}"
	io.WriteString(w, stringresult)
}
