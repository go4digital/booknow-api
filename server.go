package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go4digital/booknow-api/models"
	leads "github.com/go4digital/booknow-api/repo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from Book Now Api !\n")
	})

	http.HandleFunc("/leads", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			leads, err := leads.GetAllLeads()

			if err != nil {
				log.Fatalf("Unable to get all user. %v", err)
			}

			// send all the users as response
			json.NewEncoder(w).Encode(leads)
		case http.MethodPost:
			var lead models.Lead

			err := json.NewDecoder(r.Body).Decode(&lead)

			if err != nil {
				log.Fatalf("Unable to decode the request body.  %v", err)
			}

			leadId := leads.InsertLead(lead)

			json.NewEncoder(w).Encode(leadId)
		}

	})
	log.Printf("Server runing on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
