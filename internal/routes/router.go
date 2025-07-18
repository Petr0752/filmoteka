package routes

import (
	"filmoteka/internal/controller"
	"filmoteka/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ah *controller.ActorHandler, mh *controller.MovieHandler, uh *controller.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Авторизация
	r.POST("/login", uh.Login)

	auth := r.Group("/", middleware.AuthMiddleware())
	admin := auth.Group("/", middleware.RoleMiddleware("admin"))

	// Актёры public
	auth.GET("/actors/list", ah.List)
	auth.GET("/actors/:id", ah.GetByID)

	// Актёры private
	admin.POST("/actors", ah.Create)
	admin.PATCH("/actors/:id", ah.Update)
	admin.DELETE("/actors/:id", ah.Delete)

	// Фильмы public
	auth.GET("/movies/list", mh.List)
	auth.GET("/movies/search", mh.Search)

	// Фильмы private
	admin.POST("/movies", mh.Create)
	admin.PATCH("/movies/:id", mh.Update)
	admin.DELETE("/movies/:id", mh.Delete)
	admin.POST("/movies/:movie_id/actors/:actor_id", mh.AddActorToMovie)

	return r
}
