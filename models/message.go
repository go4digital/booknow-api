package models

import "github.com/99designs/gqlgen/graphql"

type Message struct {
	Id              int64            `json:"id"`
	FirstName       string           `json:"firstName"`
	LastName        string           `json:"lastName"`
	Email           string           `json:"email"`
	Phone           string           `json:"phone"`
	Description     string           `json:"description"`
	Address         string           `json:"address"`
	CompanyId       int64            `json:"companyId"`
	Files           []graphql.Upload `json:"files"`
	CompanyFolderId string           `json:"companyFolderId"`
}
