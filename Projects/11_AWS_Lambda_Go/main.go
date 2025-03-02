package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Job struct {
	JobID string `json:"job_id"`
	Value int    `json:"value"`
}

type JobResult struct {
	JobID  string `json:"job_id"`
	Result int    `json:"result"`
}

// DynamoDB Client
var dbClient *dynamodb.DynamoDB
var tableName = "job-results-outbox-v1"
var region = "ap-south-1"

// Initialize AWS Services
func init() {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	dbClient = dynamodb.New(sess)
}

// Worker function
func worker(id int, jobs <-chan Job, results chan<- JobResult, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	semaphore <- struct{}{}

	for job := range jobs {
		fmt.Printf("Worker %d processing job : %s\n", id, job.JobID)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		results <- JobResult{JobID: job.JobID, Result: job.Value * 3}
	}

	<-semaphore
}

func storeInDynamoDB(result JobResult) error {
	_, err := dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"JobID":  {S: aws.String(result.JobID)},
			"Result": {N: aws.String(fmt.Sprintf("%d", result.Result))},
		},
	})
	return err
}

// AWS Lambda handler: Fetch Jobs from SQS, Process and Store in DynamoDB
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	numWorkers := 3
	semaphore := make(chan struct{}, numWorkers)

	jobs := make(chan Job, len(sqsEvent.Records))
	results := make(chan JobResult, len(sqsEvent.Records))
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg, semaphore)
	}

	// Parse SQS Messages as Jobs
	for _, record := range sqsEvent.Records {
		var job Job
		if err := json.Unmarshal([]byte(record.Body), &job); err != nil {
			log.Fatalf("Error parsing job : %v", err)
			continue
		}
		jobs <- job
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if err := storeInDynamoDB(result); err != nil {
			log.Fatalf("Error storing result in DynamoDB: %v", err)
		} else {
			log.Printf("Stored result for JobID %s\n", result.JobID)
		}
	}
	return nil
}

// // Fan-In function (merges multiple worker outputs)
// func fanIn(channels ...<-chan string) <-chan string {
// 	out := make(chan string)
// 	var wg sync.WaitGroup

// 	for _, ch := range channels {
// 		wg.Add(1)
// 		go func(c <-chan string) {
// 			defer wg.Done()
// 			for msg := range c {
// 				out <- msg
// 			}
// 		}(ch)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

// func handler(ctx context.Context) ([]string, error) {
// 	numWorkers := 3
// 	numJobs := 10
// 	semaphore := make(chan struct{}, numWorkers)

// 	jobs := make(chan Job, numJobs)
// 	results := make(chan string, numJobs)
// 	var wg sync.WaitGroup

// 	// Start worker pool (Fan-Out)
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg, semaphore)
// 	}

// 	// Send jobs
// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- Job{ID: j, Value: j * 10}
// 	}
// 	close(jobs)

// 	// Wait for workers to finish and close results channel
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Collect results using Fan-In
// 	resultChannel := fanIn(results)
// 	var finalResults []string
// 	for res := range resultChannel {
// 		finalResults = append(finalResults, res)
// 	}

// 	return finalResults, nil
// }

func main() {
	fmt.Println("Implementing AWS Lambda Functions using Go Programming")

	lambda.Start(handler)
}
