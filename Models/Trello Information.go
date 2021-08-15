package Models

type TrelloInfo struct {
	AppToken string `json:"appToken",bson:"appToken"`
	UserID   string `json:"userId"bson:"userId"`
	BoardID  string `json:"boardId",bson:"boardId"`
}
