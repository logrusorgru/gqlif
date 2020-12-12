package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/logrusorgru/gqlif/internal/directives"
	"github.com/logrusorgru/gqlif/internal/graph/server"
	"github.com/logrusorgru/gqlif/internal/resolver"
)

func main() {

	var conf server.Config
	conf.Resolvers = resolver.NewResolver()
	conf.Directives.AllowValueType = directives.AllowValueType

	var (
		schema = server.NewExecutableSchema(conf)
		srv    = handler.New(schema)
	)

	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:3000/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
