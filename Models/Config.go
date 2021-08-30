package Models

var AppConfig *Config

type Config struct {
	DBConnectionURL  string `json:"DBConnectionURL"`
	AppKey           string `json:"SECRET_KEY"`
	HostMail         string `json:"HostMail"`
	HostMailPassword string `json:"HostMailPassword"`
	UILink           string `json:"UILink"`
	IssuesPath       string `json:"IssuesPath"`
}
