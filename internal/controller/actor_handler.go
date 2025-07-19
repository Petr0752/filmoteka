package controller

import (
	"filmoteka/internal/controller/dto"
	"filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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

// Create актера
// @Summary Создать нового актёра
// @Tags Актёры
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param actor body dto.ActorDTO true "Актёр"
// @Success 201 {object} map[string]int64
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /actors [post]
func (h *ActorHandler) Create(c *gin.Context) {
	var actorDTO dto.ActorDTO
	if err := c.ShouldBindJSON(&actorDTO); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := dto.ValidateActorDTO(&actorDTO); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.svc.Add(dto.ActorDTOToModel(&actorDTO))
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// List актеров
// @Summary Получить список актёров
// @Tags Актёры
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Actor
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /actors/list [get]
func (h *ActorHandler) List(c *gin.Context) {
	res, err := h.svc.ListWithMovies(h.mr)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get actors with movies"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetByID актера
// @Summary Получить актёра по ID
// @Tags Актёры
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID актёра"
// @Success 200 {object} model.Actor
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /actors/{id} [get]
func (h *ActorHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	actor, err := h.svc.Get(id, h.mr)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	c.JSON(http.StatusOK, actor)
}

// Update актёра
// @Summary Обновить данные актёра
// @Tags Актёры
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID актёра"
// @Param actor body dto.ActorDTO true "Актёр"
// @Success 204 {string} string "no content"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /actors/{id} [patch]
func (h *ActorHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var actorDTO dto.ActorDTO
	if err := c.ShouldBindJSON(&actorDTO); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	actorDTO.ID = id
	actor := dto.ActorDTOToModel(&actorDTO)

	if err := h.svc.Update(actor); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete актера
// @Summary Удалить актёра
// @Tags Актёры
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID актёра"
// @Success 204 {string} string "no content"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /actors/{id} [delete]
func (h *ActorHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
