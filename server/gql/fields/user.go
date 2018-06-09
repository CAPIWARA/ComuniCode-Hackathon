package fields

import (
	"capiwara-boilerplate/gql/resolvers"
	"capiwara-boilerplate/gql/types"

	"github.com/graphql-go/graphql"
)

var UserQuery = &graphql.Field{
	Type:        types.User,
	Description: "get user",
	Resolve:     resolvers.GetUser,
}
