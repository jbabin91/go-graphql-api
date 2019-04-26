package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	postgres "github.com/jbabin91/go-graphql-api/db"
	"github.com/jbabin91/go-graphql-api/gql"
	"github.com/jbabin91/go-graphql-api/server"
)

func main() {
	// Initialize the api and return a pointer to the router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	router, db := initializeAPI()
	defer db.Close()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	// Create a new connection to our postgres database
	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "postgres", "docker", "go_graphql_db"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)

	// Create a new graphql schema, passing in the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // Match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // Recover from panics without crashing the server
	)

	// Create the graphql route with a Server method to handle it
	router.Get("/graphql", s.GraphQL())

	return router, db
}
