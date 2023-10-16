package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph"
	"github.com/PrameshKarki/event-management-golang/graph/middlewares/auth"
	resolver "github.com/PrameshKarki/event-management-golang/graph/resolvers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	initLogger()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := gin.Default()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	configs.GetDatabaseConnection()
	router.Use(auth.Middleware())
	router.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	logrus.Info("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Fatal(http.ListenAndServe(":"+port, router))
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}
