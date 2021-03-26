package main

import (
	"fmt"
	"github.com/rs/cors"
	"main.go/Controllers"
	"main.go/Database"
	"net/http"
)

func main() {
	r := Controllers.NewRouter()
	//init Redis server
	Database.GetDbInstance()
	fmt.Print("Server is running at port 8001...")
	//InitAllController(*app, r, Redis)
	//allow all method CORS
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8001", handler)

}
