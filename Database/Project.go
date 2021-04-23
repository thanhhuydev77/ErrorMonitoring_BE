package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main.go/General"
	"main.go/Models"
)

func CreateProject(project Models.Project) (bool, General.ErrorCode) {

	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false, General.DATABASE_ERROR
	}
	project.Id = General.CreateUUID()
	collection := clientInstance.Database(General.DB).Collection(General.Project)
	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), project)
	if err != nil {
		log.Print(err.Error())
		return false, General.UNKNOWN_ERROR
	}
	return true, General.NO_ERROR
}

func ChangeStatusProject(project Models.Project) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "id", Value: project.Id}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "active", Value: project.Active},
	}}}
	collection := clientInstance.Database(General.DB).Collection(General.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

////get a user or all user(id = -1)
func GetProjects(email string, Id string) ([]Models.Project, error) {

	//db, err := connectdatabase()
	//// Query all users
	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	list := []Models.Project{}
	var filter bson.D
	if len(email) > 0 {
		if Id == "" {
			filter = bson.D{primitive.E{Key: "createuser", Value: email}} //bson.D{{}} specifies 'all documents'
		} else {
			filter = bson.D{primitive.E{Key: "createuser", Value: email}, primitive.E{Key: "id", Value: Id}}
		}
	} else {
		if Id == "" {
			filter = bson.D{} //bson.D{{}} specifies 'all documents'
		} else {
			filter = bson.D{primitive.E{Key: "id", Value: Id}}
		}
	}

	client, err := GetMongoClient()
	if err != nil {
		return list, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(General.DB).Collection(General.Project)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return list, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		var t Models.Project
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

func UpdateUserList(project Models.Project) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "id", Value: project.Id}}
	var updater bson.D
	//Define updater for to specifiy change to be updated.
	updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "userlist", Value: project.UserList},
	}}}
	collection := clientInstance.Database(General.DB).Collection(General.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}
