package GeneralFunction

import (
	"crypto/tls"
	"encoding/json"
	gomail "gopkg.in/mail.v2"
	"main.go/Models"
)

//func ConvertUserRequesttoUser(request Models.UserRequest) Models.User {
//	var user Models.User
//	user.Email = request.Email
//	user.PassWord = request.PassWord
//	user.FullName = request.FullName
//	return user
//}

func SendMail(to string, object string, text string) bool {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", Models.AppConfig.HostMail)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", object)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", text)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, Models.AppConfig.HostMail, Models.AppConfig.HostMailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return false
	}

	return true
}

func CreateResponse(Type int, Message string, Data interface{}) string {
	var failRes Models.Respond
	failRes.Status = Type
	failRes.Message = Message
	failRes.Data = Data
	result, _ := json.Marshal(failRes)
	return string(result)
}