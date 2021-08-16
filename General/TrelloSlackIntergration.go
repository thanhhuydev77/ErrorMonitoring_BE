package General

import (
	"github.com/adlio/trello"
	"github.com/slack-go/slack"
	"log"
	"main.go/Models"
)

func TrelloCreateCard(appKey string, token string, boardId string, listId string, issue Models.Issue) {
	client := trello.NewClient(appKey, token)
	//get board
	//board, err := client.GetBoard(boardId, trello.Defaults())
	//get list
	list, err := client.GetList(listId)
	if list == nil {
		log.Print(err.Error())
		return
	}

	//create card
	list.AddCard(&trello.Card{Name: issue.Name, Desc: "## MESSAGE\n" + issue.Description + "\n## EXCEPTION\n" + issue.Frame + "\n## ADDITIONAL DATA\n" + issue.Detail}, trello.Defaults())
	if err != nil {
		log.Println(err.Error())
	}
}
func SlackCreateNortification(botToken string, ChannelId string, issue Models.Issue) {
	api := slack.New(botToken)
	attachment := slack.Attachment{
		Pretext: issue.Description,
		Text:    "automate!",
	}
	_, _, err := api.PostMessage(
		ChannelId,
		slack.MsgOptionText(issue.Detail, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Print("%s\n", err)
	}
}
func GetListInBoard(appToken string, userId string, boardId string) []*trello.List {
	client := trello.NewClient(appToken, userId)
	//get board
	board, _ := client.GetBoard(boardId, trello.Defaults())
	//get list
	lists, _ := board.GetLists(trello.Defaults())
	return lists
}
