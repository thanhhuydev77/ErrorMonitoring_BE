package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io"
	"log"
	"main.go/Business"
	"main.go/CONST"
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
		log.Print(err1.Error())
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
		errCode := CONST.UNKNOWN_ERROR
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
		EmailExisted := Business.CheckUserExsist(user.User.Email)
		if EmailExisted == nil {
			result := General.CreateResponse(0, `Email unregistered!`, Models.EmptyObject{})
			io.WriteString(w, result)
			return
		}
		user.User = *EmailExisted
		min := 100000
		max := 999999
		//code random
		randCode := rand.Intn(max-min) + min
		//token
		token := GenerateToken(user.User.Email)
		//mail
		SentOK := General.SendMail(user.User.Email, CONST.EMAILSUBJECT, strconv.Itoa(randCode), user.User.FullName)

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

	return
}

func GetUserByProjectId(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	ProjectId := vars["Id"]
	ListUser, Err := Business.GetUsersByProjectId(ProjectId)
	if Err != nil || len(ListUser) == 0 {
		result := General.CreateResponse(0, `Get users by Project Id failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}
	result := General.CreateResponse(1, `Get users by Project Id success!`, ListUser)
	io.WriteString(w, result)
	return

}

func authenUser(w http.ResponseWriter, r *http.Request) {
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

func SearchInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	filter := query.Get("filter")
	if len(filter) == 0 {
		fmt.Println("filters not present")
	}

	List, err := Business.SearchUser(filter)
	if err != nil || len(List) == 0 {
		result := General.CreateResponse(0, `Search users failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	result := General.CreateResponse(1, `Search users success!`, List)
	io.WriteString(w, result)
	return
}

func TestTrello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	General.TrelloCreateCard()

	io.WriteString(w, "ok")
	return
}
