package main

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TODO struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Date  string             `json:"date" bson:"date"`
	Done  bool               `json:"done" bson:"done"`
}

func CreateTodo(c echo.Context) error {
	var todo TODO

	if err := c.Bind(&todo); err != nil {
		return err
	}
	mongoClient := getMongoClient()

	coll := mongoClient.Database("todo").Collection("todos")

	res, err := coll.InsertOne(c.Request().Context(), todo)

	if err != nil {
		return err
	}
	return c.JSON(200, res)
}

func GetTodos(c echo.Context) error {
	mongoClient := getMongoClient()

	coll := mongoClient.Database("todo").Collection("todos")

	var todos []TODO

	cursor, err := coll.Find(c.Request().Context(), bson.D{})

	if err != nil {
		return err
	}

	for cursor.Next(c.Request().Context()) {
		var todo TODO
		err := cursor.Decode(&todo)
		if err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	if err := cursor.Close(c.Request().Context()); err != nil {
		return err
	}

	return c.JSON(200, todos)
}

func MarkTodoDone(c echo.Context) error {
	todoId := c.Param("id")

	// Convert string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}

	mongoClient := getMongoClient()
	coll := mongoClient.Database("todo").Collection("todos")

	// Use objID instead of todoId
	res, err := coll.UpdateOne(c.Request().Context(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"done": true}})
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

func DeleteTodo(c echo.Context) error {
	todoId := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}

	mongoClient := getMongoClient()
	coll := mongoClient.Database("todo").Collection("todos")

	res, err := coll.DeleteOne(c.Request().Context(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
