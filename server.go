package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/directives"
	"github.com/PrameshKarki/event-management-golang/graph"
	"github.com/PrameshKarki/event-management-golang/graph/middlewares/auth"
	resolver "github.com/PrameshKarki/event-management-golang/graph/resolvers"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
	configs.GetDatabaseConnection()
	router := gin.Default()
	// Initialize middlewares
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	router.Use(auth.Middleware())

	router.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}, Directives: graph.DirectiveRoot{
		ShouldBeAuthenticated: directives.ShouldBeAuthenticated(),
	}}))

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
