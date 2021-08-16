package Controllers

import (
	"fmt"
	"io"
	"main.go/General"
	"main.go/Models"
	"net/http"
)

func GetListBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()
	appToken := query.Get("appToken")
	userId := query.Get("userId")
	boardId := query.Get("boardId")
	if len(appToken) == 0 || len(userId) == 0 {
		fmt.Println("appToken or userId is empty")
	}

	List := General.GetListInBoard(appToken, userId, boardId)
	if len(List) == 0 {
		result := General.CreateResponse(0, `Search users failed!`, Models.EmptyObject{})
		io.WriteString(w, result)
		return
	}

	result := General.CreateResponse(1, `Search users success!`, List)
	io.WriteString(w, result)
	return
}
