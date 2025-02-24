package database

import (
	"context"
	"log"
	"time"

	"github.com/devilzs1/graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://localhost:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Ping to check if connection is successful
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	log.Println("✅ Connected to MongoDB")

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database("jobs-dashboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}

	return &jobListing

}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollection := db.client.Database("jobs-dashboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}
	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("jobs-dashboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobListing model.JobListing
	inserted, err := jobCollection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"company":     jobInfo.Company,
		"url":         jobInfo.URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedId := inserted.InsertedID.(primitive.ObjectID).Hex()
	log.Fatalf("Inserted job id : %s", insertedId)
	jobListing = model.JobListing{
		ID:      insertedId,
		Title:   jobInfo.Title,
		Company: jobInfo.Company,
		URL:     jobInfo.URL,
	}
	return &jobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo *model.UpdateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("jobs-dashboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	updateJobInfo := bson.M{}
	if jobInfo.Title != "" {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != "" {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != "" {
		updateJobInfo["url"] = jobInfo.URL
	}
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}
	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))
	var jobListing model.JobListing
	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollection := db.client.Database("jobs-dashboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
