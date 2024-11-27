package handlers

import (
	"PruebaGoFinanzauto/internal/domain/models"
	"PruebaGoFinanzauto/internal/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	useCase *usecases.StudentUseCase
}

func NewStudentHandler(useCase *usecases.StudentUseCase) *StudentHandler {
	return &StudentHandler{
		useCase: useCase,
	}
}

func (h *StudentHandler) Create(c echo.Context) error {
	student := new(models.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.CreateStudent(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) GetAll(c echo.Context) error {
	filters := make(map[string]interface{})
	if name := c.QueryParam("name"); name != "" {
		filters["first_name"] = name
	}

	students, err := h.useCase.GetStudents(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}

	student, err := h.useCase.GetStudentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}

	student := new(models.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.UpdateStudent(uint(id), student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}

	if err := h.useCase.DeleteStudent(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
