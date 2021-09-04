package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/legato/graph"
	"github.com/legato/graph/generated"
	"github.com/legato/interface/controller"
	"github.com/legato/registry"
	"net/http"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	graphQLServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			TrackRepository:       registry.NewTrackRepository(),
			AlbumRepository:       registry.NewAlbumRepository(),
			AlbumArtistRepository: registry.NewAlbumArtistRepository(),
			GenreRepository:       registry.NewGenreRepository(),
		},
	}))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	r.GET("/graphql", func(ctx *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.POST("/graphql", func(ctx *gin.Context) {
		graphQLServer.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.GET("/music/:path_hash/download", controller.MusicController.GetDownloadMusic)

	Router = r
}
