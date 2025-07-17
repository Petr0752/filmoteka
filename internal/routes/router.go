package routes

import (
	"filmoteka/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ah *controller.ActorHandler, mh *controller.MovieHandler) *gin.Engine {
	r := gin.Default()

	// Актёры
	r.POST("/actors", ah.Create)
	r.GET("/actors/list", ah.List)
	r.GET("/actors/:id", ah.GetByID)
	r.PATCH("/actors/:id", ah.Update)
	r.DELETE("/actors/:id", ah.Delete)

	// Фильмы
	r.POST("/movies", mh.Create)
	r.GET("/movies/list", mh.List)
	r.GET("/movies/search", mh.Search)
	r.PATCH("/movies/:id", mh.Update)
	r.DELETE("/movies/:id", mh.Delete)
	r.POST("/movies/:movie_id/actors/:actor_id", mh.AddActorToMovie)

	return r
}
