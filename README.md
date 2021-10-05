# Introduction

Booking management system using GraphQL API

# Tools & Technologies

1. Go
2. Postgres
3. GraphQL
4. Bun ORM
5. Logrus

# Setup

1. **Install [Go 1.13 or greater](https://dl.google.com/go/go1.13.darwin-amd64.pkg)**.
   - The recommended approach is to use the installer to get started.
2. **Install Postgres**
   - https://www.postgresql.org/download/
3. **Ensure that you have a `BookNow`** database created
4. ```diff
   - Update `CONNECTION_STRING` inside .env file, keys mentioned below:
    APPLICATION_PORT=8080
    CONNECTION_STRING=postgres://user:pwd@localhost:5432/db?sslmode=disable
    GOOGLE_DRIVE_FOLDER_ID=
    GOOGLE_CAPTCHA_VERIFICATION_URL=https://www.google.com/recaptcha/api/siteverify
    GOOGLE_CAPTCHA_SECRET_KEY=

5. **Start the magic by running the following command**
   - `go run main.go`
6. Any update in .graphql file will require to regenerate gqlgen code, run below command
   - `go get github.com/99designs/gqlgen/cmd@v0.14.0`
   - `go run github.com/99designs/gqlgen generate`

# Sample Queries

```graphql
query {
  {
    messages {
      firstName
      lastName
      email
      phone
      description
      address
    }
  }
}
```

```graphql
mutation {
  saveMessage(
    input: {
      firstName: "Test"
      lastName: "User"
      email: "test.user@example.com"
      phone: "9874561236"
      description: "need cleaning service"
      address: "05 DL chetra gali"
      companyId: 1
      files: "list of files"
    }
  ) {
    id
    firstName
    lastName
    email
    phone
    description
    address
  }

  saveEnquiry(
    input: {
      firstName: "Test"
      lastName: "User"
      email: "test.user@example.com"
      phone: "9874561236"
      description: "need cleaning service"
      address: "05 DL chetra gali"
      companyId: 1
    }
  ) {
    id
    firstName
    lastName
    email
    phone
    description
    address
  }
}
```
# Data creation flow

1. Create company table and insert company info fetched from src/data/contact.json (when user fist time clicks on Contacts page in the tenant's website), company id will be fetched in subsequent calls
2. Create Person & Contact tables and insert records for the tenant after creating the tenant/company
3. Create mapping for CompanyPerson, PersonContact, CompanyContact and insert records
4. Create message from Join/Contact pages
5. On Google login button, capture subcontractor's first name, last name & email address and create Person to above tenant
6. Upload documents and save to tenant's Google drive folder when submitted on Join page

# Google Drive setup for File upload

1. Create GCP account using your gmail account, follow below link

- https://console.cloud.google.com/

2. After that login into gcp console.
3. Create a new project.
4. Select the newly created project.
5. After creating the project, to use the google drive api, we need to enable it for our project, So select the project and click on Enable APIS and SERVICES link and search and select the google drive api and click on enable.
6. We have used google Service Account to authenticate google drive api, To Know more about service account follow this link.

- https://developers.google.com/identity/protocols/oauth2/service-account

6. Follow this link to create service account.
7. Download the service account json file and put the content into `google_service_account_key.json` file in project root.
8. After that login into your google drive and create a new folder where you want to upload files.
9. Right click on the new folder and share it with service account email id, email can we found inside downloaded json.
10. Double click on newly created folder and copy the folder id from the url in the address bar.
11. Add the GOOGLE_DRIVE_FOLDER_ID in .env file
12. After following all the steps, file upload should work fine.
13. All the code related to google drive file upload is in services/fileUpload.go file.

**Follow Example Link to learn how to use gqlgen**

> https://github.com/oshalygin/gqlgen-pg-todo-example
