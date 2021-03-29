package Controllers

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"io"
	"main.go/Models"
	"net/http"
)

//JWT authorization middleware
func AuthMW(next http.Handler) http.Handler {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		Extractor: jwtmiddleware.FromFirst(jwtmiddleware.FromAuthHeader,
			jwtmiddleware.FromParameter("token")),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(Models.AppConfig.AppKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}

//Validate Token
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	stringresult := `{
		"status": 200,
			"message": "Validate success",
			"data": {
			"status": 1
		}
	}`
	io.WriteString(w, stringresult)
	return
}
