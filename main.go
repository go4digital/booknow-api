package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/global"
	"github.com/go4digital/booknow-api/graph/generated"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/resolvers"
)

func main() {
	port := global.Getenv("APPLICATION_PORT")

	db := database.Connect()

	database.CreateSchema(db)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			DB: db,
		},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Info(fmt.Sprintf("Server running on localhost:%s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	defer db.Close()

}
