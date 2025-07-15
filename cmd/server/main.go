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

	ar := repository.NewActorPG(db)
	mr := repository.NewMoviePG(db)

	as := service.NewActorService(ar)
	ms := service.NewMovieService(mr)

	ah := controller.NewActorHandler(as, mr)
	mh := controller.NewMovieHandler(ms)

	http.HandleFunc("/actors", ah.Create)
	http.HandleFunc("/actors/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.GetByID(w, r)
		case http.MethodPatch:
			ah.Update(w, r)
		case http.MethodDelete:
			ah.Delete(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
	http.HandleFunc("/actors/list", ah.List)

	http.HandleFunc("/movies", mh.Create)
	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			mh.Update(w, r)
		case http.MethodDelete:
			mh.Delete(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
	http.HandleFunc("/movies/list", mh.List)
	http.HandleFunc("/movies/search", mh.Search)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
