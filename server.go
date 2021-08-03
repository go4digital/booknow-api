package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"database/sql"

	leads "github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/utils"
)

const (
    ID = "id"
)

func getLeads(request *http.Request, response http.ResponseWriter, db *sql.DB) {
    leadId := request.URL.Query().Get(ID)
    if leadId != "" {
        leadId, err := strconv.ParseInt(leadId, 10, 64)

        if err != nil {
            msg := fmt.Sprintf("Invalid lead ID ! %v", leadId)
            log.Println(msg)
            response.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(response).Encode(msg)
        } else {
            lead, err := leads.GetLead(db, leadId)
            if err != nil {
                msg := fmt.Sprintf("No Data found for Id: %v", leadId)
                log.Println(msg)
                response.WriteHeader(http.StatusNotFound)
                json.NewEncoder(response).Encode(msg)
            } else {
                json.NewEncoder(response).Encode(lead)
            }
        }

    } else {
        leads, err := leads.GetAllLeads(db)
        log.Println(err)
        json.NewEncoder(response).Encode(leads)
    }
}

func main() {
	utils.InitLogger()
	port := utils.Getenv("APPLICATION_PORT")

	var db = database.Connect()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, "Hello from Book Now Api !\n")
	})

	http.HandleFunc("/leads", func(response http.ResponseWriter, request *http.Request) {

		switch request.Method {
		case http.MethodGet:
			getLeads(request, response, db);
		case http.MethodPost:
			var lead leads.Lead

			err := json.NewDecoder(request.Body).Decode(&lead)

			if err != nil {
				msg := "Bad request: Invalid request body."
				log.Println(msg)
				response.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(response).Encode(msg)
			} else {
				leadId, err := lead.InsertLead(db)
				if err != nil {
					msg := fmt.Sprintf("Unable to create lead. %v", err)
					log.Println(msg)
					response.WriteHeader(http.StatusExpectationFailed)
					json.NewEncoder(response).Encode(msg)
				} else {
					msg := fmt.Sprintf("Lead created Id: %v", leadId)
					json.NewEncoder(response).Encode(msg)
				}
			}

		case http.MethodPut:
			var lead leads.Lead

			err := json.NewDecoder(request.Body).Decode(&lead)

			if err != nil {
				msg := "Bad request: Invalid request body."
				log.Println(msg)
				response.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(response).Encode(msg)
			} else {
				if lead.ID == 0 {
					msg := fmt.Sprintf("Invalid lead ID ! %v", lead.ID)
					log.Println(msg)
					response.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(response).Encode(msg)
				} else {
					rowsAffected, err := lead.UpdateLead(db)
					msg := ""
					if err != nil {
						msg = fmt.Sprintf("Unable to update lead. %v", err)
						log.Println(msg)
						response.WriteHeader(http.StatusExpectationFailed)
						json.NewEncoder(response).Encode(msg)
					} else {
						msg = fmt.Sprintf("%v Lead Updated", rowsAffected)
						json.NewEncoder(response).Encode(msg)
					}
				}
			}

		case http.MethodDelete:
			query := request.URL.Query()

			id := query.Get(ID)

			leadId, err := strconv.ParseInt(id, 10, 64)

			if err != nil {
				log.Printf("Invalid lead ID ! %v", leadId)
				json.NewEncoder(response).Encode(fmt.Sprintf("Invalid lead ID ! %v", leadId))
				return
			} else {
				rowsAffected, err := leads.DeleteLead(db, leadId)
				if err != nil {
					msg := fmt.Sprintf("Unable to delete lead. %v", err)
					log.Println(msg)
					response.WriteHeader(http.StatusExpectationFailed)
					json.NewEncoder(response).Encode(msg)
				} else {
					msg := ""
					if rowsAffected == 0 {
						msg = fmt.Sprintf("Lead not found Id: %v", leadId)
					} else {
						msg = fmt.Sprintf("Lead deleted Id: %v", leadId)
					}
					json.NewEncoder(response).Encode(msg)
				}
			}
		}

	})
	log.Printf("Server running on localhost:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	defer db.Close()

}
