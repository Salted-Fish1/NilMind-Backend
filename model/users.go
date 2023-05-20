package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OldUser struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}

func (user *User) Query(id string, username string) (users []User, err error) {
	fmt.Println("User Query")
	fmt.Println(id)

	collection := DB.Database("NilMind-backend").Collection("users")
	filter := bson.M{"username": username}
	cur, err := collection.Find(context.Background(), filter)

	for cur.Next(context.Background()) {
		var doc bson.M
		err := cur.Decode(&doc)
		if err != nil {
			return nil, err
		}
		fmt.Println(doc)
		// users = append(users, User{
		// 	Id:       doc["_id"],
		// 	Username: doc["username"].(string),
		// 	Password: doc["password"].(string),
		// 	Email:    doc["email"].(string),
		// })
	}

	cur.All(context.Background(), &users)

	// for i := range users {
	// 	users[i].Id = users[i].Id.Hex()
	// }
	// fmt.Println(users)

	if err != nil {
		return users, err
	}
	return
}
