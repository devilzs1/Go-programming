package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)


type Course struct{
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	Price int `json:"price"`
	Owner *Owner `json:"owner"`
}

type Owner struct{
	Name string `json:"name"`
	Portfolio string `json:"portfolio"`
}

// sample DB
var courses []Course

// middleware, helper - file

func (c *Course) isEmpty() bool{
	// return c.CourseName == "" && c.CourseId == "";
	return c.CourseName == ""
}

func main() {
	fmt.Println("Learning API's using GO Programming")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "course-1234", CourseName: "AI-2025", Price: 99, Owner: &Owner{Name : "Adil", Portfolio : "github.com/devilzs1"}})
	courses = append(courses, Course{CourseId: "course-2345", CourseName: "Web-2025", Price: 199, Owner: &Owner{Name : "XYZ", Portfolio : "go.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course/create", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

// serveHome

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h2>Welcome to devilzs1 world. Learning API using Golang made easier</h2>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses in db")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get specific course from db")
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("404 Bad Request! Invalid Request Payload - Expected a proper payload found Empty")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty(){
		json.NewEncoder(w).Encode("404 Bad Request! Invalid Request Payload - No data inside payload JSON")
		return
	}

	for _, existingCourse :=range courses{
		if course.CourseName == existingCourse.CourseName {
			json.NewEncoder(w).Encode("Course with the same name already exists!")
			return
		}
	}

	// generate unique course ID
	rand.Seed(uint64(time.Now().UnixMilli()))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request){
	// It can be done in two ways : 
	// 1. Fetch the course with the specific id and update the row
	// 2. Make a new entry using the id and insert a new row - But for fetching we should have some column of time which we will use to sort Desc

	fmt.Println("Update the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// 1
	for index, course := range courses{
		if course.CourseId == params["id"]{
			var updatedCourse Course
			courses = append(courses[:index],courses[index+1:]... )
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")


	// 2
	// var updatedCourse Course
	// updatedCourse.CourseId = params["id"]
	// courses = append(courses, updatedCourse)
	// json.NewEncoder(w).Encode(courses)

	// return
}


func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// 1
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]... )
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}