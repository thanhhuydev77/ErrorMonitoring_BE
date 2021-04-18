package General

import (
	"crypto/tls"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	gomail "gopkg.in/mail.v2"
	"main.go/Models"
	"regexp"
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
func ValidateEmail(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)

}

func CreateUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}
