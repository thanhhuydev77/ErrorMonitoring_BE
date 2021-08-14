package General

import (
	"github.com/adlio/trello"
	"log"
)

func TrelloCreateCard() {
	client := trello.NewClient("c7a6eed6dd294af2536520a37ab4c35e", "187c39388e4e2a52e427ce4b32378a42e548dfd1994f205764e1841a5e7f3010")
	//get board
	board, err := client.GetBoard("XMiar2ZJ", trello.Defaults())
	log.Print(board.Name)
	//get list
	lists, err := board.GetLists(trello.Defaults())
	log.Print(lists[0].Name)

	//create card
	lists[0].AddCard(&trello.Card{Name: "test create card from go", Desc: "auto create"}, trello.Defaults())
	if err != nil {
		log.Println(err.Error())
	}
}
