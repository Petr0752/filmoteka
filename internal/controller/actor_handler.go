package controller

import (
	"encoding/json"
	"filmoteka/internal/model"
	"filmoteka/internal/repository"
	"filmoteka/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type ActorHandler struct {
	svc *service.ActorService
	mr  repository.MovieRepository
}

func NewActorHandler(s *service.ActorService, mr repository.MovieRepository) *ActorHandler {
	return &ActorHandler{svc: s, mr: mr}
}

func (h *ActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var a model.Actor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	id, err := h.svc.Add(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *ActorHandler) List(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.ListWithMovies(h.mr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}

func parseID(path, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.ParseInt(idStr, 10, 64)
}

func (h *ActorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/actors/")
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}

	a, err := h.svc.Get(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	_ = json.NewEncoder(w).Encode(a)
}

func (h *ActorHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/actors/")
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}

	var a model.Actor
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

func (h *ActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/actors/")
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
