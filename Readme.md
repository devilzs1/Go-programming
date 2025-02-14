# Go Web Development Projects 

Welcome to my Go programming repository! This repository contains all the web development-related projects I have built using Go. It's designed for anyone looking to explore Go's application in web development, from beginners to intermediate-level Go enthusiasts. 

## Contents 

- **Projects Directory**: This folder contains all the individual projects I have built using Go. Each project demonstrates different web development concepts, frameworks, and patterns in Go. 
- **Web Development in Go**: This repository focuses specifically on web-based applications developed using Go, excluding basic Go language features or libraries. 

--- 

## Table of Contents 

1. Introduction 
2. Getting Started 
3. Project Structure 
4. Projects 
5. Technologies Used 
6. Contributing 
7. License 

--- 

## Introduction 

This repository serves as a learning hub for me to document and share my Go web development journey. Each project inside the **Projects** directory demonstrates specific aspects of web development, such as REST APIs, web servers, routing, authentication, and more, all implemented using Go. 

As I continue to grow my skills, I will keep adding new projects, experiments, and solutions to this repo. 

--- 

## Getting Started 

To get started with any of the projects in this repository, follow these steps: 

### Prerequisites 

Before you can run any project, make sure you have Go installed on your system. You can download it from the official Go website: 

- Go Programming Language Official Site: https://golang.org/dl/ 

### Clone the Repository 

You can clone this repository to your local machine using: 

``` bash
git clone https://github.com/devilzs1/Go-programming.git 
cd Go-programming 
```

### Run a Project 

1. Navigate to the specific project directory inside the `Projects/` folder. 
2. Run the Go application by using: 

``` bash
go run main.go 
```

3. Visit the appropriate URL in your browser (usually `http://localhost:4000` or as specified in each project's README). 

--- 

## Project Structure 

The directory structure of the repository looks like this: 
``` bash
Go-programming/ 
├── README.md 
├── projects/ 
│   ├── project-1/ 
│   │   ├── main.go 
│   │   ├── README.md 
│   │   ├── /handlers 
│   │   ├── /models 
│   │   └── /utils 
│   ├── project-2/ 
│   │   ├── main.go 
│   │   └── /routes 
│   └── project-3/ 
│       ├── main.go 
│       └── /controllers 
└── LICENSE 

```

- **`projects/`**: Contains the subdirectories for each individual project. 
- **`main.go`**: The entry point for each project. 
- **`README.md`** (per project): Contains specific details for each individual project. 
- **`/handlers`, /models, /routes, etc.**: These subdirectories contain project-specific files (like routing, business logic, models, and helpers). 

--- 

## Projects 

Here are some of the projects included in this repository: 

### 1. **Project-1: Simple Go Web Server** 
- A basic HTTP web server using the `net/http` package. 
- Implements routing and serves static files. 
- Demonstrates simple request handling and responses in Go. 

### 2. **Project-2: BookStore Management** 
- A basic HTTP web server using the `net/http` package. 
- Implements routing and serves static files. 
- Demonstrates using MySQL and Go : "github.com/jinzhu/gorm", "github.com/jinzhu/gorm/dialects/mysql"

## Project-3: Stocks Management System
A web application built using Go and PostgreSQL.
- Uses PostgreSQL as the database for managing stock information.
- Demonstrates CRUD operations for stocks, including adding, updating, and deleting stock entries.

## Project-4: Event-Driven Architecture with Kafka
A real-time event-driven system built using Go and Apache Kafka.
- Implements event-driven communication with Apache Kafka as the messaging system.
- Built to handle high-throughput, real-time messaging between services.
- Demonstrates the producer-consumer model using Go and Kafka.

## Project-5: Messaging System with RabbitMQ
A messaging system built with Go and RabbitMQ.
- Implements message queues for asynchronous processing using RabbitMQ.
- Built to handle reliable message delivery and processing.
- Demonstrates how to send and receive messages with RabbitMQ using Go.

---

## 6. **Rate-Limiter-Go**: Rate Limiting in Go
A rate-limiting application built using Go to control the number of requests within a specified time window.

- Demonstrates rate-limiting techniques in Go.
- Useful for building APIs that need to enforce limits on user requests.
- Implements various rate-limiting algorithms (e.g., token bucket, leaky bucket).

---

## 7. **GraphQL-MongoDB-Go**: GraphQL with MongoDB and Go
A Go-based API implementing GraphQL and MongoDB integration.

- Demonstrates building a GraphQL API using Go.
- Uses MongoDB for data storage and querying.
- Provides an example of building a flexible, efficient API with GraphQL.

---

## 8. **gRPC-Go**: gRPC Service with Go
A gRPC-based service implemented using Go.

- Demonstrates the use of gRPC to create efficient and scalable APIs.
- Includes gRPC server and client examples in Go.
- Designed for high-performance, low-latency communication between services.

---

## 9. **Docker_Go**: Dockerized Go Application
A Go application containerized with Docker for deployment.

- Demonstrates how to set up a Go application in a Docker container.
- Includes a Dockerfile and instructions for building and running the app in Docker.
- Focuses on deployment and containerization best practices.


Feel free to check the individual README files for more detailed information about each project. 

--- 

## Technologies Used 

This repository utilizes the following technologies and tools: 

- **Go**: The core programming language. 
- **gorilla/mux**: A powerful HTTP router for Go. 
- **fiber**: A advanced web framework similar to Express for Go. 
- **MongoDB**: NoSQL database used for some of the projects.  
- **MySQL**: MySQL database used for some of the projects.  

--- 

## Contributing 

If you'd like to contribute to this repository or help me improve the projects, feel free to: 

1. Fork the repository. 
2. Create a new branch (`git checkout -b feature-name`). 
3. Make your changes. 
4. Commit your changes (`git commit -am 'Add new feature'`). 
5. Push to the branch (`git push origin feature-name`). 
6. Create a new Pull Request. 

Feel free to open an issue if you have any questions or feedback! 

--- 

