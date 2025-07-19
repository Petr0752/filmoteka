// @title Filmoteka API
// @version 1.0
// @description REST API для управления базой данных фильмов и актёров
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"database/sql"
	_ "filmoteka/docs"
	"filmoteka/internal/controller"
	"filmoteka/internal/repository"
	"filmoteka/internal/routes"
	"filmoteka/internal/service"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://film_admin:secret@localhost:5432/filmoteka?sslmode=disable")
	if err != nil {
		log.Fatal("DB connect error:", err)
	}
	defer db.Close()

	actorRepo := repository.NewActorRepository(db)
	movieRepo := repository.NewMovieRepository(db)
	userRepo := repository.NewUserRepository(db)

	actorService := service.NewActorService(actorRepo)
	movieService := service.NewMovieService(movieRepo)
	userService := service.NewUserService(userRepo)

	actorHandler := controller.NewActorHandler(actorService, movieRepo)
	movieHandler := controller.NewMovieHandler(movieService)
	userHandler := controller.NewAuthHandler(userService)

	router := routes.SetupRouter(actorHandler, movieHandler, userHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
