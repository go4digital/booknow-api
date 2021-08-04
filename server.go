package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go4digital/booknow-api/controllers"
	"github.com/go4digital/booknow-api/global"
)

func init() {
    global.InitLogger()
}

func main() {
    port := global.Getenv("APPLICATION_PORT")
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, "Hello from Book Now Api !\n")
	})

	http.HandleFunc("/leads", func(response http.ResponseWriter, request *http.Request) {

		switch request.Method {
            case http.MethodGet:
                leadId := request.URL.Query().Get(global.ID)
                if leadId != "" {
                    controllers.GetLead(request, response)
                } else {
                    controllers.GetAllLeads(request, response)
                }
            case http.MethodPost:
                controllers.CreateLead(request, response)

            case http.MethodPut:
                controllers.UpdateLead(request, response)

            case http.MethodDelete:
                controllers.DeleteLead(request, response)
		}
	})
	log.Printf("Server running on localhost:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
