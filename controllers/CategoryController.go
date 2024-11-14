package controllers

import (
	"github.com/pt-sinan-akbar/entities"
	"github.com/pt-sinan-akbar/Helpers"
	"github.com/pt-sinan-akbar/Services"

	// "github.com/pt-sinan-akbar/Entitites"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	cs *services.CategoryService
}

func NewCategoryController(categoryManager *services.CategoryService) CategoryController {
	return CategoryController{categoryManager}
}

// GetAllCategory godoc
// @Summary Get All Category Data
// @Description Get All Category Data
// @Tags categories
// @Produce json
// @Success 200 {array} entities.Category
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /categories [get]
func (cc CategoryController) GetAll(c *gin.Context) {
	obj, err := cc.cs.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// GetCategoryByID godoc
// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags categories
// @Produce json
// @Param id path int true "data"
// @Success 200 {object} entities.Category
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /categories/{id} [get]
func (cc CategoryController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	obj, err := cc.cs.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// CreateCategory godoc
// @Summary Create Category
// @Description Create Category
// @Tags categories
// @Accept json
// @Produce json
// @Param data body entities.Category true "data"
// @Success 200 {object} entities.Category
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /categories [post]
func (cc CategoryController) CreateAsync(c *gin.Context) {
	var obj entities.Category
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err := cc.cs.CreateAsync(obj); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// DeleteCategory godoc
// @Summary Delete Category
// @Description Delete Category
// @Tags categories
// @Produce json
// @Param id path int true "data"
// @Success 200 {object} entities.Category
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /categories/{id} [delete]
func (cc CategoryController) DeleteAsync(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err = cc.cs.DeleteAsync(id); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, helpers.ErrResponse{Message: "Successfully deleted record"})
}

// EditCategory godoc
// @Summary Edit Category
// @Description Edit Category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "data"
// @Param data body entities.Category true "data"
// @Success 200 {object} entities.Category
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /categories/{id} [put]
func (cc CategoryController) EditAsync(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	var updatedObj entities.Category
	if err := c.ShouldBindJSON(&updatedObj); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err := cc.cs.EditAsync(id, &updatedObj); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedObj)
}
