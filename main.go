package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in the envirenment")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
