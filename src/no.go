package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
)

func main() {
	var reasons []string
	reasonBytes, err := os.ReadFile("reasons.json")
	if err != nil {
		log.Println("Error opening file. ")
	}
	err = json.Unmarshal([]byte(reasonBytes), &reasons)
	fmt.Println("Number of reasons: ", len(reasons))

	// API routes
	http.HandleFunc("/no", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", reasons[rand.IntN(len(reasons))])
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
