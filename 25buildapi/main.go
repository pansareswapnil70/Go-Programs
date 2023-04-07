package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - files
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helpers
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API for learncodeonline")
	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299,
		Author: &Author{Fullname: "Swapnil Pansare", Website: "learncodeonline.in"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", CoursePrice: 399,
		Author: &Author{Fullname: "Sachin Pansare", Website: "lco.dev"}})
	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	//listen to a port
	log.Fatal(http.ListenAndServe(":5000", r))
}

//controllers - file

// serve home route
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by learncodeonline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("One Course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO data inside")
		return
	}
	//generate unique id,string
	//Append new course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}
