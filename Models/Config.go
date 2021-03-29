package Models

var AppConfig *Config

type Config struct {
	DBConnectionURL string `json:"DBConnectionURL"`
	AppKey          string `json:"SECRET_KEY"`
}
