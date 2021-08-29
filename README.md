# Introduction

Booking management system GraphQL API

# Tools & Technologies

1. Go
2. Postgres
3. GraphQL
4. GO-PG ORM
5. Logrus


# Setup

1. **Install [Go 1.13 or greater](https://dl.google.com/go/go1.13.darwin-amd64.pkg)**.
   - The recommended approach is to use the installer to get started.
2. **Install Postgres**
   - https://www.postgresql.org/download/
3. **Ensure that you have a `BookNow`** database created
4. ```diff
   - Update `CONNECTION_STRING` inside .env file
   ```
5. **Start the magic by running the following command**
   - `go run main.go`


# Sample Queries

```graphql
query {
  {
  leads {
    firstName
    lastName
    email
    phone
    description
  }
}
}
```

```graphql
mutation {
  createLead(
    input: {
      firstName: "Test"
      lastName: "User"
      email: "test.user@example.com"
      phone: "9874561236"
      description: "need cleaning service"
    }
  ) {
    id
    firstName
    lastName
    email
    phone
    description
    createdAt
  }
}
```


**Follow Example Link**
> https://github.com/oshalygin/gqlgen-pg-todo-example
