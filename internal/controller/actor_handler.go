package controller

import (
<<<<<<< HEAD
	"encoding/json"
	"filmoteka/internal/model"
	"filmoteka/internal/repository"
	"filmoteka/internal/service"
	"net/http"
=======
	"filmoteka/internal/model"
	"filmoteka/internal/repository"
	"filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
>>>>>>> master
)

type ActorHandler struct {
	svc *service.ActorService
	mr  repository.MovieRepository
}

func NewActorHandler(s *service.ActorService, mr repository.MovieRepository) *ActorHandler {
	return &ActorHandler{svc: s, mr: mr}
}

<<<<<<< HEAD
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
=======
func (h *ActorHandler) Create(c *gin.Context) {
	var a model.Actor
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	id, err := h.svc.Add(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *ActorHandler) List(c *gin.Context) {
	res, err := h.svc.ListWithMovies(h.mr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get actors with movies"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ActorHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	actor, err := h.svc.Get(id, h.mr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func (h *ActorHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var a model.Actor
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	a.ID = id
	if err := h.svc.Update(&a); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ActorHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
>>>>>>> master
}
