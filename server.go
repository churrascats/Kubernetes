package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ConfigMap", ConfigMap)
	http.HandleFunc("/Hello", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	w.Write([]byte(fmt.Sprintf("Hello, I'm %s and I'm %s years old.", name, age)))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading config file %s", err)
	}

	w.Write([]byte(fmt.Sprintf("Hello, my family is: %s", string(data))))
}
