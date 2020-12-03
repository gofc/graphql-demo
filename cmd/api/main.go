package main

import (
	"github.com/gofc/graphql-demo/graph"
	"github.com/gofc/graphql-demo/internal/repository"
	"github.com/gofc/graphql-demo/pkg/logger"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofc/graphql-demo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	defer logger.Close()

	resolver := graph.NewResolver(repository.NewRepository())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
