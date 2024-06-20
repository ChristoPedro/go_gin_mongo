package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	collection *mongo.Collection
}

func main() {

	mongo_string := os.Getenv("MONGO_STRING")
	db := "sample_mflix"
	collection := "movies"

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongo_string))
	if err != nil {
		panic(err)
	}

	coll := client.Database(db).Collection(collection)
	server := &Server{collection: coll}

	router := gin.Default()
	router.GET("/movies/:movie", server.getMovies)
	router.GET("/movies", server.getAllMovies)
	router.GET("/movie/:id", server.getMovieByID)

	router.Run("localhost:8080")
}

func (s Server) getMovies(c *gin.Context) {

	title := c.Param("movie")

	namePattern := fmt.Sprintf(".*%s.*", regexp.QuoteMeta(title))
	regex := bson.M{"$regex": namePattern, "$options": "i"}
	filter := bson.M{"title": regex}

	var result []bson.M
	cursor, err := s.collection.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding movies from database"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}

func (s Server) getAllMovies(c *gin.Context) {

	var result []bson.M
	cursor, err := s.collection.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title")
		return
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding movies from database"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}

func (s Server) getMovieByID(c *gin.Context) {

	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var result bson.M
	err = s.collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title")
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}
