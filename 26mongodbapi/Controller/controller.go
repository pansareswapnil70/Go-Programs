package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/pansareswapnil70/mongodbapi/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://learncodeonline:swap5174@cluster0.tsi7rip.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colname = "watchlist"

// Mongo DB Collection
var collection *mongo.Collection

// Connect with mongodb
func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Mongo DB connected successfully")
	collection = client.Database(dbName).Collection(colname)

	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated 1 movie in DB", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	Deletedcount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted 1 movie in DB", Deletedcount)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}}
	DeleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies. Deleted Count is: ", DeleteResult.DeletedCount)
	return DeleteResult.DeletedCount
}

//get all movies from collection

func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer curr.Close(context.Background())
	return movies
}

// Actual Controller file
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
