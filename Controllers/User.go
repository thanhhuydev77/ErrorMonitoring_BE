package Controllers

import (
	"encoding/json"
	_ "github.com/gorilla/mux"
	"io"
	"log"
	"main.go/Business"
	"main.go/General"
	"main.go/Models"
	"math/rand"
	"net/http"
	"strconv"
)

func UserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user := Models.UserRequest{}
	err1 := json.NewDecoder(r.Body).Decode(&user)
	if err1 != nil {
		result := General.CreateResponse(0, `wrong format, please try again!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if user.Type == "login" {

		IsExist, passOk := Business.Login(user.User.Email, user.User.Password)

		if (!passOk) || (!IsExist) {
			result := General.CreateResponse(0, `Email or password is incorrect, please try again`, Models.EmptyObject{})
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
		result := General.CreateResponse(1, `Login success`, Data)
		io.WriteString(w, result)
		return
	}
	if user.Type == "register" {
		errCode := General.UNKNOWN_ERROR
		if General.ValidateEmail(user.User.Email) {
			_, errCode = Business.Register(user.User)
		}
		switch errCode {
		case 0:
			result := General.CreateResponse(1, `Register success!`, Models.EmptyObject{})
			io.WriteString(w, result)
		case 1:
			result := General.CreateResponse(0, `Email is already exists, please use another email.`, Models.EmptyObject{})
			io.WriteString(w, result)
		default:
			result := General.CreateResponse(0, `Register failed, please try again.`, Models.EmptyObject{})
			io.WriteString(w, result)
		}

		return
	}
	if user.Type == "update" {
		token := r.URL.Query().Get("token")
		user.User.Email = General.GetEmailFromToken(token)
		if len(user.User.Password) > 50 {
			user.User.Password = ""
		}
		isSuccessed := Business.Update(user.User)
		if !isSuccessed {
			result := General.CreateResponse(0, `Update user failed!`, Models.EmptyObject{})
			io.WriteString(w, result)
		}
		result := General.CreateResponse(1, `Update user success!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	if user.Type == "forgot-password" {
		EmailExsisted := Business.CheckUserExsist(user.User.Email)
		if !EmailExsisted {
			result := General.CreateResponse(0, `Email unregistered!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
		min := 100000
		max := 999999
		//code random
		randCode := rand.Intn(max-min) + min
		//token
		token := GenerateToken(user.User.Email)
		//mail
		SentOK := General.SendMail(user.User.Email, General.EMAILSUBJECT, General.EMAILTEXT+strconv.Itoa(randCode))

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
			result := General.CreateResponse(1, `Request change password success`, Data)
			io.WriteString(w, result)
			return
		}
		result := General.CreateResponse(0, `Request change password failed`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//Id := vars["Id"]

	token := r.URL.Query().Get("token")
	email := General.GetEmailFromToken(token)
	log.Print("email :" + email)
	List, err := Business.GetUsers(email)
	if err != nil {
		result := General.CreateResponse(0, `Unauthentication!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	result := General.CreateResponse(1, `Authentication success!`, List[0])
	io.WriteString(w, result)
	return
}
