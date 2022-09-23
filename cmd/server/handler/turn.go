package handler

import (
	"errors"
	"strconv"

	"github.com/desafioFinalBack/internal/domain"
	"github.com/desafioFinalBack/internal/turn"
	"github.com/desafioFinalBack/pkg/web"

	"github.com/gin-gonic/gin"
)

type turnHandler struct {
	t turn.Service
}

// NewTurnHandler crea un nuevo controller de turnos
func NewTurnHandler(t turn.Service) *turnHandler {
	return &turnHandler{
		t: t,
	}
}

// GetByID obtiene un turno por su id
func (h *turnHandler) GetTurnByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		turn, err := h.t.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		web.Success(c, 200, turn)
	}
}

// GetByID obtiene un turno por su id
func (h *turnHandler) GetTurnByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		turn, err := h.t.GetByDNI(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		web.Success(c, 200, turn)
	}
}

// PostTurn crea un nuevo turno
func (h *turnHandler) PostTurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turn domain.Turn
		err := c.ShouldBindJSON(&turn)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysTurn(&turn)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.t.Create(turn)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// PutTurn actualiza un turno
func (h *turnHandler) PutTurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var turn domain.Turn
		err = c.ShouldBindJSON(&turn)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysTurn(&turn)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		t, err := h.t.Update(id, turn)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t)
	}
}

// PatchTurn actualiza un turno
func (h *turnHandler) PatchTurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var TurnNew domain.TurnDTO
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if err := c.ShouldBindJSON(&TurnNew); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		update := domain.Turn{
			Date: TurnNew.Date,
			Time: TurnNew.Date,
			Description: TurnNew.Description,
		}

		t, err := h.t.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t)
	}
}

// DeleteTurn elimina un turno
func (h *turnHandler) DeleteTurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.t.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}