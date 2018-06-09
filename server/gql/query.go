package gql

import (
	"capiwara-boilerplate/gql/fields"

	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "query",
		Description: "query",
		Fields: graphql.Fields{
			"user": fields.UserQuery,
		},
	},
)
