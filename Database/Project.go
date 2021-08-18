package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main.go/CONST"
	"main.go/Models"
)

func CreateProject(project Models.Project) (bool, CONST.ErrorCode) {

	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false, CONST.DATABASE_ERROR
	}

	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)
	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), project)
	if err != nil {
		log.Print(err.Error())
		return false, CONST.UNKNOWN_ERROR
	}
	return true, CONST.NO_ERROR
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
	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

////get a user or all user(id = -1)
func GetProject(Id string) ([]Models.Project, error) {

	//db, err := connectdatabase()
	//// Query all users
	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	list := []Models.Project{}
	var filter bson.D

	if Id == "" {
		filter = bson.D{} //bson.D{{}} specifies 'all documents'
	} else {
		filter = bson.D{primitive.E{Key: "id", Value: Id}}
	}

	client, err := GetMongoClient()
	if err != nil {
		return list, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(CONST.DB).Collection(CONST.Project)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter, options.Find().SetProjection(bson.M{"issues": 0}))
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
	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

func SearchProject(filter string) ([]Models.Project, error) {

	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	list := []Models.Project{}

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
	collection := client.Database(CONST.DB).Collection(CONST.Project)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), query, options.Find().SetProjection(bson.M{"issues": 0}))
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

func UpdateIssueList(project Models.Project) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "id", Value: project.Id}}
	var updater bson.D
	//Define updater for to specifiy change to be updated.
	updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "issues", Value: project.Issues},
	}}}
	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	//Return success without any error.
	return true
}

func GetProjectWithIssue(Id string) ([]Models.Project, error) {

	//db, err := connectdatabase()
	//// Query all users
	if clientInstance == nil {
		log.Print("can not connect to database!")
		return nil, nil
	}
	//defer db.Close()
	list := []Models.Project{}
	var filter bson.D

	if Id == "" {
		filter = bson.D{} //bson.D{{}} specifies 'all documents'
	} else {
		filter = bson.D{primitive.E{Key: "id", Value: Id}}
	}

	client, err := GetMongoClient()
	if err != nil {
		return list, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(CONST.DB).Collection(CONST.Project)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter, options.Find())
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

func UpdateSuiteList(project Models.Project) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "id", Value: project.Id}}
	var updater bson.D
	//Define updater for to specifiy change to be updated.
	updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "suites", Value: project.Suites},
	}}}
	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}

func UpdateIntegration(project Models.Project, Type int) bool {
	if clientInstance == nil {
		Err := "can not connect to database!"
		log.Print(Err)
		return false
	}

	filter := bson.D{primitive.E{Key: "id", Value: project.Id}}
	updater := bson.D{}
	if Type == 1 {
		//Trello
		updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "enableTrello", Value: project.EnableTrello},
			primitive.E{Key: "trelloInfo", Value: project.TrelloInfo},
		}}}
	} else {
		if Type == 2 {
			//Slack
			updater = bson.D{primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "enableSlack", Value: project.EnableSlack},
				primitive.E{Key: "slackInfo", Value: project.SlackInfo},
			}}}
		}
	}

	collection := clientInstance.Database(CONST.DB).Collection(CONST.Project)

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return false
	}
	//Return success without any error.
	return true
}
