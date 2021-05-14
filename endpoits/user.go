package endpoits

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"perf_test/structs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "empty response",
		})
		return
	}
	var createUserRequest structs.CreateUserRequest
	if err := json.Unmarshal(jsonData, &createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong bady",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnection := os.Getenv("MONGO")
	if mongoConnection == "" {
		mongoConnection = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnection))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	user_collection := client.Database("testing").Collection("user")
	res, err := user_collection.InsertOne(ctx, createUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail to creating user",
		})
		return
	}
	id := res.InsertedID
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnection := os.Getenv("MONGO")
	if mongoConnection == "" {
		mongoConnection = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnection))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var getUserResponse structs.GetUserResponse

	userCollection := client.Database("testing").Collection("user")
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong id format",
		})
		return
	}
	filter := bson.M{"_id": filterId}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = userCollection.FindOne(ctx, filter).Decode(&getUserResponse)

	if err == mongo.ErrNoDocuments {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, getUserResponse)
}
