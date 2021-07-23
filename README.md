# booknow-api
APIs for booking management system

## Prerequisite

- run `go mod tidy`
- run `go run server.go` this will run server on http://localhost:8080

## Database Setup is necessary to run this API
   We have used Postgres, below are the DB details:-

   > **Database Connection String mentioned in project .env file, If your machine postgres config is different then update in .env file other wise it will throw connection error**

   - DataBase Name:- **BookNow**
   - Table Name:- **Leads**
   - Table script:-
   ```
   CREATE TABLE leads
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    firstname character varying(250) COLLATE pg_catalog."default",
    lastname character varying(250) COLLATE pg_catalog."default",
    email character varying(300) COLLATE pg_catalog."default",
    phone character varying(20) COLLATE pg_catalog."default",
    query text COLLATE pg_catalog."default",
    CONSTRAINT leads_pkey PRIMARY KEY (id)
)
   ```



## Lead API Endpoint details

We have implemented CRUD operation on leads table below are the details:-

- GET:- http://localhost:8080/leads
- POST:- http://localhost:8080/leads
- PUT:- http://localhost:8080/leads
- DELETE:- http://localhost:8080/leads?id=''


**I followed [this link](https://codesource.io/build-a-crud-application-in-golang-with-postgresql/) for Postgres DB Connection**



