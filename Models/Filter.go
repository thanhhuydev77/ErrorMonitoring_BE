package Models

type Filter struct {
	Type   int         `json:"type"`
	Filter interface{} `json:"filter"`
}
