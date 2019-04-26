package gql

import (
	"github.com/graphql-go/graphql"
	postgres "github.com/jbabin91/go-graphql-api/db"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type.
// This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	// Create a resolver holding the database.
	// Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up.
	// In this example we have a user query that takes one argument
	// called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}

	return &root
}
