package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Lead struct {
	leads     struct{}  `pg:"leads"`
	ID        int       `json:"id" pg:"id,pk"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Query     string    `json:"query"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create Lead Table
func CreateLeadTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&Lead{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating lead table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Lead table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitializeDB(db *pg.DB) {
	dbConnect = db
	CreateLeadTable(dbConnect)
}

func GetAllLeads(context *gin.Context) {
	var leads []Lead
	err := dbConnect.Model(&leads).Select()

	if err != nil {
		log.Printf("Error while getting all leads, Reason: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Leads",
		"data":    leads,
	})
	return
}

func CreateLead(context *gin.Context) {
	var lead Lead
	lead.CreatedAt = time.Now()
	context.BindJSON(&lead)

	_, insertError := dbConnect.Model(&lead).Insert()
	if insertError != nil {
		log.Printf("Error while inserting new lead into db, Reason: %v\n", insertError)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Lead created Successfully",
	})
	return
}

func GetLead(context *gin.Context) {
	leadId, err := strconv.Atoi(context.Param("leadId"))

	lead := &Lead{ID: leadId}
	err = dbConnect.Model(lead).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting lead Reason: %v\n", err)
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Lead not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Lead",
		"data":    lead,
	})
	return
}

func UpdateLead(context *gin.Context) {
	leadId, err := strconv.Atoi(context.Param("leadId"))
	lead := &Lead{
		ID: leadId,
	}
	lead.UpdatedAt = time.Now()
	context.BindJSON(&lead)

	_, err = dbConnect.Model(lead).WherePK().Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Lead Upated Successfully",
	})
	return
}

func DeleteLead(context *gin.Context) {
	leadId, err := strconv.Atoi(context.Param("leadId"))
	lead := &Lead{ID: leadId}

	_, err = dbConnect.Model(lead).WherePK().Delete()
	if err != nil {
		log.Printf("Error while deleting a lead, Reason: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Lead deleted successfully",
	})
	return
}
