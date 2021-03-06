package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/global"
	"github.com/go4digital/booknow-api/graph/generated"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/middleware"
	"github.com/go4digital/booknow-api/resolvers"
	"github.com/go4digital/booknow-api/services"
)

var db *pg.DB
var port string

func init() {
	global.LoadEnvFile()
	port = os.Getenv("APPLICATION_PORT")

	db = database.Connect()
	database.CreateSchema(db)
}

func main() {
	defer db.Close()
	leadDao := dao.NewLeads(db)

	leadsService := services.NewLeads(leadDao)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			Service: leadsService,
		},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.CorsMiddleware(middleware.VerifyCaptcha(server)))

	log.Info(fmt.Sprintf("Server running on localhost:%s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
