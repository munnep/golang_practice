package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file

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

// fake db
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	r := mux.NewRouter()
	// we only allow the GET method on the default page
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	
	// give courses some data
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 199,Author: &Author{Fullname: "Patrick", Website: "http://p4e.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "GoLang", CoursePrice: 199,Author: &Author{Fullname: "Lotte", Website: "http://p4e.com"}})
	
	fmt.Println("The server is running on port http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers  - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to course API website</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from request (what use is asking for)
	params := mux.Vars(r)

	// loop through course, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	// if nothing is found then give this.
	// try to give back the param id
	json.NewEncoder(w).Encode("No course found: " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// if the body given is empty. no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}

	// if data is empty == {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate a unique id as string
	// append to the courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// return not actually needed but there for clarity
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// if the body given is empty. no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}

	// grap the id from the request
	params := mux.Vars(r)

	// loop id remove add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response where id is not found

	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	// if the body given is empty. no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}

	// grap the id from the request
	params := mux.Vars(r)

	// loop id remove add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted: " + params["id"])
			break
		}
	}
	// TODO: send a response where id is not found
	return
}
