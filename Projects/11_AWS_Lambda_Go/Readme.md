# AWS Lambda Worker in Go

This project demonstrates how to implement an AWS Lambda function in Go that processes jobs in parallel using worker goroutines, and stores results in DynamoDB. The Lambda function reads jobs from an SQS queue, processes them, and stores the results in a DynamoDB table.

## Architecture Overview

- **SQS**: The Lambda function is triggered by an SQS event, which contains messages representing jobs to be processed.
- **Workers**: The Lambda function launches multiple worker goroutines to process jobs concurrently.
- **DynamoDB**: The results of the jobs are stored in a DynamoDB table.
- **Concurrency Control**: A semaphore is used to limit the number of concurrent workers to avoid exceeding AWS Lambda concurrency limits.

## How It Works

1. **Lambda Handler**:
    - The Lambda function is triggered by an SQS event.
    - The function parses the messages from the SQS event and pushes them into a ```jobs``` channel.
    - Workers fetch jobs from the ```jobs``` channel, process them, and send the results to a ```results``` channel.
    - After all workers finish processing, the results are stored in a DynamoDB table.
  
2. **Worker Pool**:
    - Multiple worker goroutines are launched to handle jobs concurrently.
    - A semaphore is used to limit the number of concurrent workers.

3. **DynamoDB Storage**:
    - The result for each job is stored in a DynamoDB table named ```job-results-outbox-v1```.
  
## Requirements

- Go 1.24 or later
- AWS SDK for Go
- AWS Lambda Go SDK
- AWS account with access to DynamoDB and SQS
- AWS CLI configured with proper permissions

## Setup

### 1. Install Go and Dependencies

Make sure you have Go installed. You can download and install Go from the official [Go website](https://golang.org/).

To install the necessary dependencies, run the following command:

```go get github.com/aws/aws-lambda-go/events```

```go get github.com/aws/aws-lambda-go/lambda```

```go get github.com/aws/aws-sdk-go/aws```

```go get github.com/aws/aws-sdk-go/aws/session```

```go get github.com/aws/aws-sdk-go/service/dynamodb```

### 2. AWS Lambda Configuration

To deploy this Lambda function, you need to configure it in the AWS Lambda Console. Make sure to:

- Set the Lambda function’s runtime to Go 1.x.
- Configure the Lambda to be triggered by an SQS event.
- Set the appropriate IAM roles and permissions to allow Lambda to read from SQS and write to DynamoDB.

### 3. DynamoDB Table

Ensure that a DynamoDB table with the name ```job-results-outbox-v1``` exists in your AWS account. The table should have a partition key named ```JobID``` (string) and a sort key (optional).

### 4. Deploy Lambda Function

To deploy the Lambda function, you can either:

- Package and deploy manually using AWS CLI.
- Use the AWS Lambda Console to upload the ```.zip``` package of your Go code.

### 5. SQS Configuration

Make sure the SQS queue is set up with messages in the following format:

```{
  "job_id": "some-job-id",
  "value": 42
}```

The Lambda function will consume these messages and process the ```value``` field.

## Code Walkthrough

- **Worker function**:
  - The worker function simulates processing by sleeping for a random duration.
  - Each job is processed by multiplying its ```value``` by 3, and the result is stored in DynamoDB.

- **storeInDynamoDB**:
  - This function stores the job result in DynamoDB using the ```PutItem``` API.

- **Lambda Handler**:
  - This is the entry point for the Lambda function.
  - It creates a pool of worker goroutines and processes the jobs from SQS concurrently.
  
## Example Output

When processing a job, the Lambda function logs output similar to:

```Worker 1 processing job : job-id-1```

```Worker 2 processing job : job-id-2```

```Stored result for JobID job-id-1```

```Stored result for JobID job-id-2```

## Error Handling

- If an error occurs while processing a job or storing the result in DynamoDB, it will be logged.
- The Lambda function will retry in case of transient errors, based on the SQS retry configuration.

