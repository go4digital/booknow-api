package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/services"
)

const (
	ID = "id"
)

func GetAllLeads(request *http.Request, response http.ResponseWriter) {
	leads, err := services.LeadsService.GetAllLeads()
	log.Println(leads)
	if err != nil {
		msg := "No Leads found"
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(msg)
	} else {
		json.NewEncoder(response).Encode(leads)
	}
}

func GetLead(request *http.Request, response http.ResponseWriter) {
	leadId := request.URL.Query().Get(ID)
	if leadId != "" {
		leadId, err := strconv.ParseInt(leadId, 10, 64)

		if err != nil {
			msg := fmt.Sprintf("Invalid lead ID ! %v", leadId)
			log.Println(msg)
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(msg)
		} else {
			lead, err := services.LeadsService.GetLead(leadId)
			if err != nil {
				msg := fmt.Sprintf("No Data found for Id: %v", leadId)
				log.Println(msg)
				response.WriteHeader(http.StatusNotFound)
				json.NewEncoder(response).Encode(msg)
			} else {
				json.NewEncoder(response).Encode(lead)
			}
		}
	}
}

func CreateLead(request *http.Request, response http.ResponseWriter) {
	var lead dao.Lead

	err := json.NewDecoder(request.Body).Decode(&lead)

	if err != nil {
		msg := "Bad request: Invalid request body."
		log.Println(msg)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(msg)
	} else {
		leadId, err := services.LeadsService.CreateLead(&lead)
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
}
func UpdateLead(request *http.Request, response http.ResponseWriter) {
	var lead dao.Lead

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
			rowsAffected, err := services.LeadsService.UpdateLead(&lead)
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
}
func DeleteLead(request *http.Request, response http.ResponseWriter) {
	query := request.URL.Query()

	id := query.Get(ID)

	leadId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Printf("Invalid lead ID ! %v", leadId)
		json.NewEncoder(response).Encode(fmt.Sprintf("Invalid lead ID ! %v", leadId))
		return
	} else {
		rowsAffected, err := services.LeadsService.DeleteLead(leadId)
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
