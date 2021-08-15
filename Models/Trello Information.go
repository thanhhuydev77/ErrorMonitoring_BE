package Models

type TrelloInfo struct {
	AppToken string `json:"appToken",bson:"appToken"`
	UserID   string `json:"userId"bson:"userId"`
	BoardID  string `json:"boardId",bson:"boardId"`
	ListID   string `json:"listId",bson:"listId"`
}
type TrelloBoardInfo struct {
	BoardName string `json:"boardName",bson:"boardName"`
	BoardID   string `json:"boardId",bson:"boardId"`
}
