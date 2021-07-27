package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	leads "github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/postgres"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var db = postgres.Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from Book Now Api !\n")
	})

	http.HandleFunc("/leads", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			leads, err := leads.GetAllLeads(db)

			if err != nil {
				log.Fatalf("Unable to get all user. %v", err)
			}

			// send all the users as response
			json.NewEncoder(w).Encode(leads)
		case http.MethodPost:
			var lead leads.Lead

			err := json.NewDecoder(r.Body).Decode(&lead)

			if err != nil {
				log.Fatalf("Unable to decode the request body.  %v", err)
			}

			leadId := lead.InsertLead(db)

			json.NewEncoder(w).Encode(leadId)

		case http.MethodPut:
			var lead leads.Lead

			err := json.NewDecoder(r.Body).Decode(&lead)

			if err != nil {
				log.Fatalf("Unable to decode the request body.  %v", err)
			}

			if lead.ID == 0 {
				log.Fatalf("Invalid lead ID ! %v", lead.ID)
			}

			rowsAffected := lead.UpdateLead(db)

			json.NewEncoder(w).Encode(rowsAffected)
		case http.MethodDelete:
			query := r.URL.Query()

			id := query.Get("id")

			leadId, err := strconv.ParseInt(id, 10, 64)

			if err != nil {
				log.Printf("Invalid lead ID ! %v", leadId)
				json.NewEncoder(w).Encode(fmt.Sprintf("Invalid lead ID ! %v", leadId))
				return
			}
			lead := leads.Lead{ID: leadId}
			rowsAffected := lead.DeleteLead(db)

			json.NewEncoder(w).Encode(rowsAffected)
		}

	})
	log.Printf("Server runing on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	defer db.Close()

}
