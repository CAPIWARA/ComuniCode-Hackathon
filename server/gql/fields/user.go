package fields

import (
	"github.com/VitorLuizC/ComuniCode-Hackathon/server/gql/resolvers"
	"github.com/VitorLuizC/ComuniCode-Hackathon/server/gql/types"

	"github.com/graphql-go/graphql"
)

var UserQuery = &graphql.Field{
	Type:        types.User,
	Description: "get user",
	Resolve:     resolvers.GetUser,
}
