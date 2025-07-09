package main

import (
	"database/sql"
	"filmoteka/internal/controller"
	"filmoteka/internal/repository"
	"filmoteka/internal/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres",
		"postgres://film_admin:secret@localhost:5432/filmoteka?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// --- репозитории
	ar := repository.NewActorPG(db)
	mr := repository.NewMoviePG(db)

	// --- сервисы
	as := service.NewActorService(ar)
	ms := service.NewMovieService(mr)

	// --- контроллеры
	ah := controller.NewActorHandler(as, mr)
	mh := controller.NewMovieHandler(ms)

	// простейший роутинг
	http.HandleFunc("/actors", ah.Create)        // POST
	http.HandleFunc("/actors/list", ah.List)     // GET
	http.HandleFunc("/movies", mh.Create)        // POST
	http.HandleFunc("/movies/list", mh.List)     // GET ?sort=
	http.HandleFunc("/movies/search", mh.Search) // GET ?q=

	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
