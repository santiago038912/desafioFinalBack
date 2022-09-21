package handler

import (
	// "github.com/desafioFinalBack/internal/domain"
	"errors"
	"strconv"

	"github.com/desafioFinalBack/internal/dentist"
	"github.com/desafioFinalBack/internal/domain"
	"github.com/desafioFinalBack/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// GetByID obtiene un producto por su id
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

// StoreDentists godoc
// @Summary Store dentists
// @Tags Dentists
// @Description store dentists
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body domain.Dentist true "Dentist to store"
// @Success 200 {object} web.response
// @Router /products [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Put actualiza un producto
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch update selected fields of a product WIP
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name           string `json:"name,omitempty"`
		LastName       string `json:"last_name,omitempty"`
		RegisterNumber string `json:"register_number,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			Name:           r.Name,
			LastName:       r.LastName,
			RegisterNumber: r.RegisterNumber,
		}

		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Delete elimina un producto
func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}




// validateEmptys valida que los campos no esten vacios
func validateEmptys(dentist *domain.Dentist) (bool, error) {
	if dentist.Name == "" || dentist.LastName == "" || dentist.RegisterNumber == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
