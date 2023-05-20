package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id               primitive.ObjectID `bson:"_id"`
	User_id          primitive.ObjectID `bson:"user_id"`
	Name             string             `bson:"name"`
	Created_at       time.Time          `bson:"created_at"`
	Last_update_time time.Time          `bson:"last_update_time"`
	Path             string             `bson:"path"`
}

func (file *File) Query(userId string) (files []File, err error) {
	// func (file *File) Query(userId string) (files []interface{}, err error) {
	fmt.Println("File Query")
	fmt.Println(userId)
	collection := DB.Database("NilMind-backend").Collection("workspaces")
	nid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return files, err
	}
	// fmt.Println(nid)
	filter := bson.M{"user_id": nid}

	cur, err := collection.Find(context.Background(), filter)
	cur.All(context.Background(), &files)
	fmt.Println("File Query")
	fmt.Println(files[0].Id)
	if err != nil {
		return files, err
	}
	for i := range files {
		files[i].Path = "http://localhost:8080/files/" + files[i].Id.Hex() + ".zip"
		// fmt.Println(files[i].Path)
	}
	// fmt.Println(files)
	return files, nil
}
