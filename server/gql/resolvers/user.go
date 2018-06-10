package resolvers

import (
	"comunicode/server/users"
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("email").(string)
	user, err := users.FindByEmail(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Checkout(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("email").(string)
	value := params.Args["value"].(string)
	user, _ := users.FindByEmail(token)
	if err := user.Checkout(value); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println("Checkout completo, saving in database")
	if err := user.Alter(); err != nil {
		return nil, err
	}
	fmt.Println("Update at mongodb")
	return user, nil
}
