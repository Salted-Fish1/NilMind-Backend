package schema

import (
	"golesson/model"

	"github.com/graphql-go/graphql"
)

var workspaceType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "workspace",
	Description: "workspace Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"user_id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"last_update_time": &graphql.Field{
			Type: graphql.DateTime,
		},
		"path": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var workspaceHello = graphql.Field{
	Name:        "Queryworkspace",
	Description: "Query workspace",
	Type:        graphql.NewList(workspaceType),
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
		"last_update_time": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
		"path": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		userId, _ := p.Args["user_id"].(string)

		return (&model.File{}).Query(userId)
	},
}
