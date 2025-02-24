# GraphQL-MongoDB-Go

## Overview  
This project is a GraphQL API built using ```gqlgen``` and MongoDB with Go. It provides a structured way to interact with a MongoDB database using GraphQL queries and mutations.  

## Prerequisites  
Ensure you have the following installed before proceeding:    
- [gqlgen](https://github.com/99designs/gqlgen)  

## Running the Application  
To start the GraphQL server, run:  
```go run server.go```  
The server will start at ```http://localhost:8080```.  

## Generating GraphQL Code  
If you update the GraphQL schema, regenerate the code using:  
```go run github.com/99designs/gqlgen generate```  

## GraphQL Queries and Mutations  
You can test the API by navigating to ```http://localhost:8080``` and using the GraphQL Playground.  

Example Query:  
```query { jobListings { id title description } }```  

Example Mutation:  
```mutation { updateJobListing(jobId: "123", jobInfo: { title: "New Title" }) { id title } }```  

## Troubleshooting  
### ```gqlgen generate``` Issues  
If you encounter errors while generating GraphQL code, try:  
```rm -rf graph/generated```  
```rm -rf graph/prelude.resolvers.go```  
```go run github.com/99designs/gqlgen generate```  

### MongoDB Connection Issues  
Ensure MongoDB is running and the connection string is correctly set.  
