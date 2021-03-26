package Controllers

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"io"
	"main.go/Database"
	"net/http"
)

//JWT authorization middleware
func AuthMW(next http.Handler) http.Handler {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		Extractor: jwtmiddleware.FromFirst(jwtmiddleware.FromAuthHeader,
			jwtmiddleware.FromParameter("token")),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(Database.APP_KEY), nil
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
