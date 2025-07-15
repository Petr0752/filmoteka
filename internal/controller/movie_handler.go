package controller

import (
	"encoding/json"
	"filmoteka/internal/model"
	"filmoteka/internal/service"
	"net/http"
)

type MovieHandler struct {
	svc *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{svc: s}
}

func (h *MovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	var m model.Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	id, err := h.svc.Add(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *MovieHandler) List(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	res, err := h.svc.List(sort)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (h *MovieHandler) Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	res, err := h.svc.Search(q)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (h *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/movies/")
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}

	var a model.Movie
	if err = json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	a.ID = id

	if err = h.svc.Update(&a); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(204)
}

func (h *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/movies/")
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}

	if err = h.svc.Delete(id); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(204)
}
