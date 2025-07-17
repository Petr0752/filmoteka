package controller

import (
	"filmoteka/internal/controller/dto"
	"filmoteka/internal/model"
	"filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ActorHandler struct {
	svc *service.ActorService
	mr  service.MovieRepository
}

func NewActorHandler(s *service.ActorService, mr service.MovieRepository) *ActorHandler {
	return &ActorHandler{svc: s, mr: mr}
}

func (h *ActorHandler) Create(c *gin.Context) {
	var actorDTO dto.ActorDTO
	if err := c.ShouldBindJSON(&actorDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := dto.ValidateActorDTO(&actorDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.svc.Add(dto.ActorDTOToModel(&actorDTO))
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
}
