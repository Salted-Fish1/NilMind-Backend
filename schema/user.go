package schema

import (
	"fmt"
	"golesson/model"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "User Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: graphql.NewList(workspaceType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// 实现查询逻辑
				fmt.Println("------------")
				// id, _ := params.Source.(map[string]interface{})["id"].(primitive.ObjectID)
				user := params.Source.(model.User)
				return (&model.File{}).Query(user.Id)
				// model.File{}.
			},
		},
	},
})

var userHello = graphql.Field{
	Name:        "QueryUser",
	Description: "Query User",
	Type:        graphql.NewList(userType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"file": &graphql.ArgumentConfig{
			Type: graphql.NewList(workspaceType),
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, _ := p.Args["id"].(string)
		username, _ := p.Args["username"].(string)
		fmt.Println("userHello")

		return (&model.User{}).Query(id, username)
	},
}

var UserInput = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "User Model",
	Fields: graphql.Fields{
		// "id": &graphql.Field{
		// 	Type: graphql.ID,
		// },
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var userMutation = graphql.Field{
	Name:        "userMutation",
	Description: "userMutation",
	Type:        UserInput,
	Args: graphql.FieldConfigArgument{
		// "id": &graphql.ArgumentConfig{
		// 	Type: graphql.ID,
		// },
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		fmt.Println(p.Info.Path.AsArray()...)
		id, _ := p.Args["id"].(string)
		// username, _ := p.Args["username"].(string)

		// return (&model.User{}).Query(id, username)
		return id, nil
	},
}
