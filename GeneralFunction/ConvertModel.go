package GeneralFunction

import (
	"encoding/json"
	"main.go/Models"
)

func ConvertUserRequesttoUser(request Models.UserRequest) Models.User {
	var user Models.User
	user.Email = request.Email
	user.PassWord = request.PassWord
	user.FullName = request.FullName
	return user
}

func CreateResponse(Type int, Message string, Data interface{}) string {
	var failRes Models.Respond
	failRes.Status = Type
	failRes.Message = Message
	failRes.Data = Data
	result, _ := json.Marshal(failRes)
	return string(result)
}
