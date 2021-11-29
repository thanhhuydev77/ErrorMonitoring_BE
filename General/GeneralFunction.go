package General

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/nu7hatch/gouuid"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"log"
	"main.go/CONST"
	"main.go/Models"
	"regexp"
	"time"
)

func SendMail(to string, object string, text string, Name string, Template string) bool {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", Models.AppConfig.HostMail)
	// Set E-Mail receivers
	m.SetHeader("To", to)

	type info struct {
		Name string
		Text string
	}
	t := template.New(Template)
	Info := info{Name: Name, Text: text}
	var err error
	t, err = t.ParseFiles("Template/" + Template)
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, Info); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	// Set E-Mail subject
	m.SetHeader("Subject", object)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", result)

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
	return emailRegex.MatchString(e) && CheckMailExistence(e)
}

func CreateUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}
func GetEmailFromToken(tokenString string) string {

	type MyCustomClaims struct {
		User string `json:"user"`
		jwt.StandardClaims
	}
	var email string
	// sample token is expired.  override time so it parses as valid
	at(time.Unix(0, 0), func() {
		token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
		claims, _ := token.Claims.(*MyCustomClaims)
		email = claims.User

	})
	return email
}
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
func CheckMailExistence(mail string) bool {
	err := checkmail.ValidateHostAndUser(CONST.MAILSMTP, "", mail)
	if _, ok := err.(checkmail.SmtpError); ok && err != nil {
		return false
	}
	return true
}

//func UploadPicture(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	file, handler, err := r.FormFile("file")
//	if err != nil {
//		//fmt.Print(err)
//		io.WriteString(w, `{"message":"Can’t upload avatar"}`)
//		return
//	}
//	defer file.Close()
//
//	f, _ := os.OpenFile("../public/images/avatars/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//
//	defer f.Close()
//	io.Copy(f, file)
//	io.WriteString(w, `{ "status": 200,
//    "message": "Upload avatar success",
//    "data": {
//        "fieldname": "file",
//        "originalname": "`+handler.Filename+`",
//        "destination": "public",
//		 "mimetype": "`+handler.Header.Get("Content-Type")+`",
//        "filename": "`+handler.Filename+`",
//        "path": "public\\images\\avatars\\`+handler.Filename+`",
//        "size": `+strconv.Itoa(int(handler.Size))+`
//    }
//}`)
//}
