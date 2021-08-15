package Models

type SlackInfo struct {
	BotToken string `json:"botToken",bson:"botToken"`
	ChanelId string `json:"chanelId",bson:"chanelId"`
}
