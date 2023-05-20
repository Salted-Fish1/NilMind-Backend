package files

import (
	"context"
	"fmt"
	"golesson/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func Post(c *gin.Context) {
	userId := c.Request.FormValue("userId")
	fmt.Println(userId)
	file, header, err := c.Request.FormFile("file")
	fmt.Println(file)
	if err != nil {
		fmt.Println(file, header, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "File Resolve Fail",
		})
		return
	}
	defer file.Close()
	collection := model.DB.Database("NilMind-backend").Collection("workspaces")
	id, _ := primitive.ObjectIDFromHex(userId)
	document := bson.D{
		{"user_id", id},
		{"name", header.Filename},
		{"created_at", time.Now()},
		{"last_update_time", time.Now()},
	}
	primitive.NewObjectID().Hex()
	result, _ := collection.InsertOne(context.Background(), document)
	fileId := result.InsertedID.(primitive.ObjectID).Hex()
	SaveFile(header, "files/"+fileId+".zip")

	c.JSON(http.StatusOK, gin.H{
		"data": "File Upload Success",
	})
}

func DELETE(c *gin.Context) {
	collection := model.DB.Database("NilMind-backend").Collection("workspaces")
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.D{{"_id", id}}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete the document",
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The document does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Document Delete Success",
	})
}

func FindAll(c *gin.Context) {
	collection := model.DB.Database("NilMind-backend").Collection("workspaces")
	filter := bson.D{}
	result, _ := collection.Find(context.Background(), filter)

	// var doc []model.File
	var doc interface{}
	err := result.Decode(&doc)
	fmt.Println(doc)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The document does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "",
	})
}
