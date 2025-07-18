package controller

import (
	"filmoteka/internal/controller/dto"
	"filmoteka/internal/model"
	"filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	svc *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{svc: s}
}

func (h *MovieHandler) Create(c *gin.Context) {
	var movieDTO dto.MovieDTO
	if err := c.ShouldBindJSON(&movieDTO); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := dto.ValidateMovieDTO(&movieDTO); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.svc.Add(dto.MovieDTOToModel(&movieDTO))
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *MovieHandler) List(c *gin.Context) {
	sort := c.Query("sort")
	res, err := h.svc.List(sort)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *MovieHandler) Search(c *gin.Context) {
	q := c.Query("q")
	res, err := h.svc.Search(q)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *MovieHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var m model.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	m.ID = id
	if err := h.svc.Update(&m); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MovieHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MovieHandler) AddActorToMovie(c *gin.Context) {
	movieIDStr := c.Param("movie_id")
	actorIDStr := c.Param("actor_id")

	movieID, err := strconv.ParseInt(movieIDStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid movie id"})
		return
	}

	actorID, err := strconv.ParseInt(actorIDStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid actor id"})
		return
	}

	if err := h.svc.AddActor(movieID, actorID); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
