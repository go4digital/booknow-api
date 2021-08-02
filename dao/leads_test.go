package leads

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var l = &Lead{
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

func TestInsertLead(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	query := `insert into "leads"("firstname", "lastname", "email", "phone", "query")
	values($1, $2, $3, $4, $5) RETURNING id`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.FirstName, l.LastName, l.Email, l.Phone, l.Description).WillReturnResult(sqlmock.NewResult(0, 1))

	id, _ := l.InsertLead(db)
	assert.Equal(t, id, 1)
}

func TestGetLead(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "SELECT * FROM leads"

	rows := mock.NewRows([]string{"id", "firstname", "lastname", "email", "phone", "query"}).
		AddRow(l.ID, l.FirstName, l.LastName, l.Email, l.Phone, l.Description)

	mock.ExpectQuery(query).WillReturnRows(rows)

	leads, err := GetAllLeads(db)
	assert.NotNil(t, leads)
	assert.NoError(t, err)
}
