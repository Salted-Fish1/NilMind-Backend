package router

import (
	"golesson/controller/files"
	"golesson/controller/graphql"
	"golesson/controller/hello"
	signupin "golesson/controller/signUp-In"
	"golesson/controller/token"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.Use(cors.Default())
	Router.GET("/news/:id", hello.Get)
	// Router.GET("/news", Hello.List)
	Router.POST("/sign-up", signupin.SignUp)
	Router.POST("/sign-in", signupin.SignIn)
	Router.POST("/sign-test", token.AuthToken(), signupin.SignTest)
	Router.POST("/verify-code", signupin.VerifyCode)

	Router.POST("/files", files.Post)
	Router.GET("/all-files", files.FindAll)
	Router.DELETE("/files/:id", files.DELETE)
	Router.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		c.File("files/" + filename)
	})

	Router.PUT("/news/:id", hello.Put)
	Router.DELETE("/news", hello.Destroy)
	Router.POST("/news", hello.Post)

	Router.POST("/graphql", graphql.GraphqlHandler())
	Router.GET("/graphql", graphql.GraphqlHandler())
	Router.Static("/files", "./files")
}
