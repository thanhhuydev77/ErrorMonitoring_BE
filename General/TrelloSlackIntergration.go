package General

import (
	"github.com/adlio/trello"
	"github.com/slack-go/slack"
	"log"
	"main.go/Models"
)

func TrelloCreateCard(appKey string, token string, boardId string, issue Models.Issue) {
	client := trello.NewClient(appKey, token)
	//get board
	board, err := client.GetBoard(boardId, trello.Defaults())
	log.Print(board.Name)
	//get list
	lists, err := board.GetLists(trello.Defaults())
	log.Print(lists[0].Name)

	//create card
	lists[0].AddCard(&trello.Card{Name: issue.Name, Desc: issue.Description}, trello.Defaults())
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
