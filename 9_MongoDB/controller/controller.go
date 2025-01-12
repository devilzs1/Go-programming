package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/devilzs1/mongodb-go/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


const connString = "mongodb://localhost:27017"
const dbName = "Netflix"
const colName = "Watchlist"


var collection *mongo.Collection

func init(){
	clientOption := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully.")
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance ready")
}


func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie inserted into db : ", inserted.InsertedID)
}

func insertManyMovie(movie []model.Netflix){
	inserted, err := collection.InsertMany(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie inserted into db : ", inserted.InsertedIDs)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Update movie detail : ", result.UpsertedID)
	fmt.Println("Modified count : ", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie detail : ", result.DeletedCount)
}

func deleteManyMovie(movieIds []string){
	var objectIds []primitive.ObjectID
	for _, id := range movieIds {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatalf("Invalid ObjectID format: %v", err)
		}
		objectIds = append(objectIds, objectId)
	}
	filter := bson.M{
		"_id": bson.M{"$in": objectIds},
	}

	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie detail : ", result.DeletedCount)
}

func deleteAllMovie() int64{
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all movie : ", deleted.DeletedCount)
	return deleted.DeletedCount
}


func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()){
		var movie bson.M 
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}


func GetAllMoviesInDB(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkOneMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}