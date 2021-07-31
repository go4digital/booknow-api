# booknow-api
APIs for booking management system

## Prerequisite
- run git clone https://github.com/go4digital/booknow-api.git
- run `go mod tidy`
- run `go run server.go` this will run server on http://localhost:8080

<br />
<br />

## Database Setup is necessary to run this API

   ### We have used Postgres, follow below steps to setup

   1. Install Postgres first on your machine follow- https://www.postgresql.org/download/

   2. Postgres Install PGAdmin development platform, for Database, Table and User setup.

      - **Note down the master password which you enter while installation, that password is require to login in pgAdmin tool.**

   3. PgAdmin has one default user name as - "postgres".

   4. Create new user by following this link:- https://chartio.com/learn/postgresql/create-a-user-with-pgadmin/

   5. Create Database follow this link:- https://www.postgresqltutorial.com/postgresql-create-database/

   6. Cretae table follow this link:- https://www.guru99.com/create-drop-table-postgresql.html


<br />
<br />

   > **Database Connection String mentioned in project .env file, please details if you have defferent config**

   > Update Database name
   > Update Port (default postgres port is 5432)
   > Update Username and password

   - DataBase Name:- **BookNow**
   - Table Name:- **Leads**
   - Table script:-
   ```
   CREATE TABLE leads
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    firstname character varying(25) COLLATE pg_catalog."default",
    lastname character varying(25) COLLATE pg_catalog."default",
    email character varying(50) COLLATE pg_catalog."default",
    phone character varying(15) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    CONSTRAINT leads_pkey PRIMARY KEY (id)
)
   ```



## Lead API Endpoint details

We have implemented CRUD operation on leads table below are the details:-

> Replace port number if you are running your Api on defferent port

- GET:- http://localhost:8080/leads
- GET:- http://localhost:8080/leads?id=''
- POST:- http://localhost:8080/leads
- PUT:- http://localhost:8080/leads
- DELETE:- http://localhost:8080/leads?id=''


**I followed [this link](https://codesource.io/build-a-crud-application-in-golang-with-postgresql/) for Postgres DB Connection**



