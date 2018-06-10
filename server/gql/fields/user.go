package fields

import (
	"comunicode/server/gql/resolvers"
	"comunicode/server/gql/types"

	"github.com/graphql-go/graphql"
)

var UserQuery = &graphql.Field{
	Type:        types.User,
	Description: "get user",
	Resolve:     resolvers.GetUser,
}

var CheckoutMutation = &graphql.Field{
	Type:        types.User,
	Description: "checkout",
	Resolve:     resolvers.Checkout,
	Args: graphql.FieldConfigArgument{
		"value": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
}
