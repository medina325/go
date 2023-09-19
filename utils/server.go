package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type course struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c course) getFullInfo() string {
	return fmt.Sprintf("Name: %s, Desc: %s, Price: %d", c.Name, c.Description, c.Price)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := course{
		Name:        "Golang",
		Description: "Oin",
		Price:       100,
	}
	println(course.Name)
	course.getFullInfo()

	json.NewEncoder(w).Encode(course)
}

func hello(w http.ResponseWriter, r *http.Request) {
	println(r)
	w.Write([]byte("Hello, World!!!"))
}

func Server() {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)

	println("Running server in localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
