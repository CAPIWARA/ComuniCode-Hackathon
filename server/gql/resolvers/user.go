package resolvers

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	fmt.Println("NICE")
	return nil, nil
}
