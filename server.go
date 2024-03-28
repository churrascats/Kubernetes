package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/Secret", Secret)
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

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	w.Write([]byte(fmt.Sprintf("User: %s and Password %s", user, password)))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration : %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
