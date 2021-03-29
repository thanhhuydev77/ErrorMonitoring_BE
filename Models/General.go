package Models

type RespondOk struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type RespondFail struct {
	Message string `json:"message"`
}
