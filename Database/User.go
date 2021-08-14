package Database

import (
	"bytes"
	"context"
	_ "database/sql"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	_ "log"
	"main.go/CONST"
	"main.go/Models"
	_ "strconv"
	"time"
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
	collection := client.Database(CONST.DB).Collection(CONST.User)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print(err.Error())
		return false, false
	}
	exsist = true
	if checkPasswordHash(pass, user.Password) {
		passOK = true
	}
	return exsist, passOK
}

//register a new user
func Register(user Models.User) (bool, CONST.ErrorCode) {

	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false, CONST.DATABASE_ERROR
	}
	if CheckDuplicateEmail(user.Email) {
		return false, CONST.DUPLICATE_EMAIL
	}
	user.Password, _ = hashPassword(user.Password)

	collection := clientInstance.Database(CONST.DB).Collection(CONST.User)
	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Print(err.Error())
		return false, CONST.UNKNOWN_ERROR
	}
	return true, CONST.NO_ERROR
}

func Update(user Models.User) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}
	filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
	var updater bson.D
	if len(user.Password) > 0 {
		user.Password, _ = hashPassword(user.Password)
		//Define updater for to specifiy change to be updated.
		updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "fullname", Value: user.FullName},
			primitive.E{Key: "password", Value: user.Password},
			primitive.E{Key: "avatar", Value: user.Avatar},
			primitive.E{Key: "mainplatform", Value: user.MainPlatform},
			primitive.E{Key: "position", Value: user.Position},
			primitive.E{Key: "organization", Value: user.Organization},
			primitive.E{Key: "projectlist", Value: user.ProjectList},
		}}}
	} else {
		updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "fullname", Value: user.FullName},
			primitive.E{Key: "avatar", Value: user.Avatar},
			primitive.E{Key: "mainplatform", Value: user.MainPlatform},
			primitive.E{Key: "position", Value: user.Position},
			primitive.E{Key: "organization", Value: user.Organization},
			primitive.E{Key: "projectlist", Value: user.ProjectList},
		}}}
	}

	collection := clientInstance.Database(CONST.DB).Collection(CONST.User)

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
	collection := client.Database(CONST.DB).Collection(CONST.User)
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

//update Project List
func UpdateProjectList(user Models.User) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
	var updater bson.D
	//Define updater for to specifiy change to be updated.
	updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "projectlist", Value: user.ProjectList},
	}}}
	collection := clientInstance.Database(CONST.DB).Collection(CONST.User)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

func SearchUser(filter string) ([]Models.User, error) {

	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	var list []Models.User

	query := bson.M{
		"$text": bson.M{
			"$search": filter,
		},
	}

	client, err := GetMongoClient()
	if err != nil {
		return list, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(CONST.DB).Collection(CONST.User)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), query, options.Find())
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

func UploadAvatar(file string, filename string) bool {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print(err)
		return false
	}
	conn, _ := GetMongoClient()
	bucket, err := gridfs.NewBucket(
		conn.Database(CONST.DB),
	)
	if err != nil {
		log.Print(err)
		return false
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Print(err)
		return false
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
	return true
}

func DownloadAvatar(fileName string) string {
	conn, _ := GetMongoClient()

	// For CRUD operations, here is an example
	db := conn.Database(CONST.DB)
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Print(err)
		return ""
	}

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		log.Print(err)
		return ""
	}
	fmt.Printf("File size to download: %v\n", dStream)
	ioutil.WriteFile(fileName+".JPG", buf.Bytes(), 0600)
	a := base64.StdEncoding.EncodeToString(buf.Bytes())
	return a
}
