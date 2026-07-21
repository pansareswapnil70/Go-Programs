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

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"courseprice"`
	Author      *Author
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func main() {
	courses = append(courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 499, Author: &Author{FullName: "Swapnil", Website: "http://lco.dev"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "PHP", CoursePrice: 399, Author: &Author{FullName: "Shripad", Website: "http://w3schools.dev"}})

	fmt.Println("Welcome to APIs")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Response from API is Hi</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Fetching all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetch One Course")
	params := mux.Vars(r)
	for _, course := range courses {
		if params["id"] == course.CourseId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.Write([]byte("No course found"))
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
