package Database

import (
	"context"
	_ "database/sql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	_ "log"
	"main.go/Models"
	_ "strconv"
)

//login
func Login(Email string, pass string) (bool, bool) {
	exsist := false
	passOK := false
	user := Models.User{}
	// Query all users
	if clientInstance == nil {
		log.Print("can not connect to database!")
		return exsist, false
	}

	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "email", Value: Email}}
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		log.Print(err.Error())
		return false, false
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(User)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print(err.Error())
		return false, false
	}
	exsist = true
	if checkPasswordHash(pass, user.PassWord) {
		passOK = true
	}
	return exsist, passOK
}

//register a new user
func Register(user Models.User) (bool, ErrorCode) {

	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false, DATABASE_ERROR
	}
	if CheckDuplicateEmail(user.Email) {
		return false, DUPLICATE_EMAIL
	}
	user.PassWord, _ = hashPassword(user.PassWord)

	collection := clientInstance.Database(DB).Collection(User)
	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Print(err.Error())
		return false, UNKNOWN_ERROR
	}
	return true, NO_ERROR
}

func Update(user Models.User) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}
	filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
	user.PassWord, _ = hashPassword(user.PassWord)
	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "fullName", Value: user.FullName},
		primitive.E{Key: "passWord", Value: user.PassWord},
		primitive.E{Key: "avatar", Value: user.Avatar},
		primitive.E{Key: "mainPlatform", Value: user.MainPlatform},
		primitive.E{Key: "position", Value: user.Position},
		primitive.E{Key: "organization", Value: user.Organization},
		primitive.E{Key: "projectList", Value: user.ProjectList},
	}}}
	collection := clientInstance.Database(DB).Collection(User)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

func CheckDuplicateEmail(email string) bool {
	listuser, _ := GetUsers("")
	for _, user := range listuser {
		if user.Email == email {
			return true
		}
	}
	return false
}

//get all user name
//func GetAllUserName() []string {
//	var Allusername []string
//	//db, err := connectdatabase()
//	// Query all users
//	if db == nil {
//
//		log.Print("can not connect to database!")
//		return nil
//	}
//	//defer db.Close()
//
//	rows, err := db.Query("select username from USERS")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for rows.Next() {
//		var username string
//		err := rows.Scan(&username)
//		if err != nil {
//			log.Fatal(err)
//		}
//		Allusername = append(Allusername, username)
//	}
//	defer rows.Close()
//	return Allusername
//}
//
////hash password by bycript
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

//
////check password hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//
////get a user or all user(id = -1)
func GetUsers(Id string) ([]Models.User, error) {

	//db, err := connectdatabase()
	//// Query all users
	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	list := []Models.User{}
	var filter bson.D
	if Id == "" {
		filter = bson.D{primitive.E{}} //bson.D{{}} specifies 'all documents'
	} else {
		filter = bson.D{primitive.E{Key: "email", Value: Id}}
	}
	client, err := GetMongoClient()
	if err != nil {
		return list, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(User)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return list, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		var t Models.User
		err := cur.Decode(&t)
		if err != nil {
			return list, err
		}
		list = append(list, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(list) == 0 {
		return list, mongo.ErrNoDocuments
	}
	return list, nil
}
