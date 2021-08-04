package dao

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go4digital/booknow-api/model"
	"github.com/stretchr/testify/assert"
)

var lead = &model.Lead{
	ID:          1,
	FirstName:   "Test",
	LastName:    "User",
	Email:       "test.user@example.com",
	Phone:       "856974213",
	Description: "I need cleaning service",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}

	return db, mock
}

func TestCreateLead(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	mockleads := NewLeads(db)

	query := "insert into leads"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Description).WillReturnResult(sqlmock.NewResult(1, 1))
	id, err := mockleads.CreateLead(lead)
	if err != nil {
		fmt.Println("this is the error message: ", err)
		t.Errorf("CreateLead() error = %v, wantErr %v", err, 1)
		return
	} else {
		assert.Equal(t, id, 1)
	}
}

func TestGetAllLeads(t *testing.T) {
	db, mock := NewMock()
	mockleads := NewLeads(db)

	want := []*model.Lead{
		{
			ID:          1,
			FirstName:   "lead1",
			LastName:    "test1",
			Email:       "lead1.test1@test.com",
			Phone:       "25625526",
			Description: "Need cleaning service",
		},
		{
			ID:          2,
			FirstName:   "lead2",
			LastName:    "test2",
			Email:       "lead2.test2@test.com",
			Phone:       "1212122",
			Description: "Need cleaning service",
		},
	}
	// query := "SELECT id, firstname, lastname, email, phone, description FROM leads"

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "phone", "description"}).AddRow(1, "lead1", "test1", "lead1.test1@test.com", "25625526", "Need cleaning service").AddRow(2, "lead2", "test2", "lead2.test2@test.com", "1212122", "Need cleaning service")

	mock.ExpectPrepare("SELECT (.+) FROM leads").ExpectQuery().WillReturnRows(rows)

	got, err := mockleads.GetAllLeads()

	if (err != nil) != true {
		t.Errorf("GetAllLeads() error new = %v, wantErr %v", err, true)
		return
	}
	if err == nil && !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllLeads() = %v, want %v", got, want)
	}
}
