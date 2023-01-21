package controllers

import (
	"context"
	"net/http"
	"real-time-chat/configuration"
	"real-time-chat/models"
	"real-time-chat/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configuration.GetCollection(configuration.DB, "users")

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "Bad request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		user.Print()

		newUser := models.User{
			Id:   primitive.NewObjectID(),
			Name: user.Name,
		}

		res, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "Error when insert", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, response.UserResponse{Status: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"data": res}})

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		user := c.Query("user")

		err := userCollection.FindOne(ctx, bson.D{{Key: "name", Value: user}})
		if err != nil {
			if err.Err() == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, response.UserResponse{Status: http.StatusNotFound, Message: "User not found", Data: map[string]interface{}{"data": err.Err()}})
				return
			}
			var user models.User
			err.Decode(&user)
			user.Print()
			c.JSON(http.StatusFound, response.UserResponse{Status: http.StatusFound, Message: "Success", Data: map[string]interface{}{"data": user}})
		}

	}
}
