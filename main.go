package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"main.go/Controllers"
	"main.go/Database"
	"main.go/Models"
	"net/http"
	"os"
)

func main() {
	ReadConfigfile()
	r := Controllers.NewRouter()
	Database.GetMongoClient()
	Controllers.InitAllController(r)
	//allow all method CORS
	handler := cors.AllowAll().Handler(r)
	port := GetPort()
	fmt.Print("Server is running at port" + port + "...")
	http.ListenAndServeTLS(port, "server.crt", "server.key", handler)
}
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8001"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
func ReadConfigfile() {
	jsonFile, err := os.Open("Config/AppConfig.text")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened successfully")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal((byteValue), &Models.AppConfig)
	log.Print("Read Config successfully!")
}
