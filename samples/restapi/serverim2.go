package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Project struct {
	ID          string  `json:"id,omitempty"`
	Filename    string  `json:"filename,omitempty"`
	Projectname string  `json:"projectname,omitempty"`
	Location    string  `json:"location,omitempty"`
	Photos      []Photo `json:"photos,omitempty"`
	Videos      []Video `json:"videos,omitempty"`
	Audios      []Audio `json:"audios,omitempty"`
}

type Photo struct {
	Name   string `json:"name,omitempty"`
	Size   string `json:"size,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

type Video struct {
	Name   string `json:"name,omitempty"`
	Size   string `json:"size,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

type Audio struct {
	Name string `json:"name,omitempty"`
	Size string `json:"size,omitempty"`
}

var projects []Project

func GetProjectEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range projects {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Project{})
}

func GetProjectsDataEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(projects)
}

func CreateProjectEndPoint(w http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	// var person Person
	// _ = json.NewDecoder(req.Body).Decode(&person)
	// person.ID = params["id"]
	// persons = append(persons, person)
	// json.NewEncoder(w).Encode(person)
}
func DeleteProjectEndPoint(w http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	// for index, item := range persons {
	// 	if item.ID == params["id"] {
	// 		persons = append(persons[:index], persons[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(persons)
}

func main() {
	fmt.Print("Starting Go Server")
	router := mux.NewRouter()
	projects = append(projects, Project{ID: "1", Filename: "NYWithProximity.im2", Projectname: "Ny With Proximity", Location: "C:\\Infotransit\\Projects\\NY With Proximity"})
	projects = append(projects, Project{ID: "2", Filename: "SanAnt.im2", Projectname: "San Antonio Default", Location: "C:\\Infotransit\\Projects\\San Antonio Presentation"})
	projects = append(projects, Project{ID: "3", Filename: "Default Presentation.im2", Projectname: "Default Presentation", Location: "C:\\Infotransit\\Projects\\Default Presentation"})
	projects = append(projects, Project{ID: "4", Filename: "My Test.im2", Projectname: "My Test For Adhoc", Location: "C:\\Infotransit\\Projects\\NY With Proximity"})
	router.HandleFunc("/projects", GetProjectsDataEndPoint).Methods("GET")
	router.HandleFunc("/project/{id}", GetProjectEndPoint).Methods("GET")
	router.HandleFunc("/project/{id}", CreateProjectEndPoint).Methods("POST")
	router.HandleFunc("/project/{id}", DeleteProjectEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8099", router))
}
