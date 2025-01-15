# BookStore Management System

Welcome to the **BookStore Management System** project! This project demonstrates the implementation of a simple bookstore management system using **Go**. It uses the `net/http` package for HTTP server functionality and integrates a MySQL database with the help of **GORM**, an ORM library for Go. It also utilizes the **gorilla/mux** router for routing requests.

## Table of Contents

1. Introduction
2. Getting Started


---

## Introduction

This project is designed to help manage books in a bookstore. The backend is built with Go and connects to a MySQL database to store information about books. This application allows basic CRUD operations (Create, Read, Update, Delete) for managing books in the system.

The system is built using:

- **Go**: The backend programming language
- **MySQL**: The relational database used to store book details
- **GORM**: A Go ORM for interacting with MySQL
- **gorilla/mux**: A HTTP router for routing API requests

---

## Getting Started

### Prerequisites

Make sure you have **Go** and **MySQL** installed on your machine.

- Download Go: https://golang.org/dl/
- Install MySQL: https://dev.mysql.com/downloads/

### Clone the Repository

You can clone the repository to your local machine and navigate to the project directory.

### Setup

1. Initialize a Go module in the project directory.
2. Install the required dependencies:
   - `github.com/jinzhu/gorm`
   - `github.com/jinzhu/gorm/dialects/mysql`
   - `github.com/gorilla/mux`

3. Configure your MySQL database with the following:
   - Create a MySQL database and set up the necessary credentials (username, password, and database name).
   - Update the connection string in the Go code to match your MySQL database configuration.

4. Run the application.

5. The application should now be running. Visit the appropriate URL in your browser (e.g., `http://localhost:4000`).


