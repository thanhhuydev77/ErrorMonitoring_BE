package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"main.go/Controllers"
	"main.go/Database"
	"main.go/Models"
	"net/http"
)

func main() {
	ReadConfigfile()
	r := Controllers.NewRouter()
	Database.GetMongoClient()
	Controllers.InitAllController(r)
	//allow all method CORS
	handler := cors.AllowAll().Handler(r)
	fmt.Print("Server is running at port 8001...")
	http.ListenAndServeTLS(":8001", "server.crt", "server.key", handler)
}

//func ReadConfigfile() {
//	jsonFile, err := os.Open("Config/AppConfig.text")
//	// if we os.Open returns an error then handle it
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Opened successfully")
//	// defer the closing of our jsonFile so that we can parse it later on
//	defer jsonFile.Close()
//
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//
//	json.Unmarshal((byteValue), &Models.AppConfig)
//	log.Print("Read Config successfully!")
//}
func ReadConfigfile() {
	//jsonFile, err := os.Open("Config/AppConfig.text")
	// if we os.Open returns an error then handle it
	a := new(Models.Config)
	a.HostMailPassword = "Thanhhuyd71t9"
	a.HostMail = "errormonitoringvn@gmail.com"
	a.AppKey = "thisissecreckeyyesitisreallyofcourcetrustmeitiskeyofthisapphahaha"
	a.DBConnectionURL = "mongodb+srv://hathanhhuy:Thanhhuyd71t9@mycluster.5dvo9.mongodb.net/test?authSource=admin&replicaSet=atlas-f8f9l2-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"
	//json.Unmarshal((byteValue), &Models.AppConfig)
	Models.AppConfig = a
	log.Print("Read Config successfully!")
}
