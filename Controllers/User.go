package Controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/gorilla/mux"
	"io"
	"log"
	"main.go/Business"
	"main.go/Database"
	"main.go/GeneralFunction"
	"main.go/Models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//login with username and pass from body request
func UserLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	param1 := r.URL.Query().Get("token")
	GetEmailFromToken(param1)
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

		IsExist, passOk := Business.Login(user.User.Email, user.User.PassWord)

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
			Email: user.User.Email,
			Token: GenerateToken(user.User.Email),
		}
		result := GeneralFunction.CreateResponse(1, `Login success`, Data)
		io.WriteString(w, result)
		return
	}

	if user.Type == "register" {
		_, errCode := Business.Register(user.User)
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

	if user.Type == "update" {
		isSuccessed := Business.Update(user.User)
		if !isSuccessed {
			result := GeneralFunction.CreateResponse(0, `Update user failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
		}
		result := GeneralFunction.CreateResponse(1, `Update user success!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if user.Type == "forgot-password" {
		min := 100000
		max := 999999
		//code random
		randCode := rand.Intn(max-min) + min

		//token
		token := GenerateToken(user.User.Email)
		//mail
		SentOK := GeneralFunction.SendMail(user.User.Email, Database.EMAILSUBJECT, Database.EMAILTEXT+strconv.Itoa(randCode))

		if SentOK {
			type data struct {
				Code  int    `json:"code"`
				Token string `json:"token"`
				Email string `json:"email"`
			}
			Data := data{
				Code:  randCode,
				Token: token,
				Email: user.User.Email,
			}
			result := GeneralFunction.CreateResponse(1, `Request change password success`, Data)
			io.WriteString(w, result)
			return
		}
		result := GeneralFunction.CreateResponse(0, `Request change password failed`, Models.EmptyObject{})
		io.WriteString(w, result)
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
	//vars := mux.Vars(r)
	//Id := vars["Id"]

	token := r.URL.Query().Get("token")
	email := GetEmailFromToken(token)
	log.Print("email :" + email)
	List, err := Business.GetUsers(email)
	if err != nil {
		result := GeneralFunction.CreateResponse(0, `Unauthentication!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	result := GeneralFunction.CreateResponse(1, `Authentication success!`, List[0])
	io.WriteString(w, result)
	return
}
