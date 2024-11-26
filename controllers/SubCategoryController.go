package controllers

import (
	"net/http"
	"strconv"
	"github.com/pt-sinan-akbar/entities"

	"github.com/pt-sinan-akbar/helpers"
	"github.com/pt-sinan-akbar/services"

	"github.com/gin-gonic/gin"
)

type SubCategoryController struct {
	scs *services.SubCategoryService
}

func NewSubCategoryController(subCategoryService *services.SubCategoryService) SubCategoryController {
	return SubCategoryController{subCategoryService}
}

// GetAllSubCategory godoc
// @Summary Get All Sub Category Data
// @Description Get All Sub Category Data
// @Tags subcategories
// @Produce json
// @Success 200 {array} entities.SubCategory
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /subcategories [get]
func (scc SubCategoryController) GetAll(c *gin.Context) {
	obj, err := scc.scs.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// GetSubCategoryByID godoc
// @Summary Get Sub Category By ID
// @Description Get Sub Category By ID
// @Tags subcategories
// @Produce json
// @Param id path int true "data"
// @Success 200 {object} entities.SubCategory
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /subcategories/{id} [get]
func (scc SubCategoryController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	obj, err := scc.scs.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// CreateSubCategory godoc
// @Summary Create Category
// @Description Create Category
// @Tags subcategories
// @Accept json
// @Produce json
// @Param data body entities.SubCategory true "data"
// @Success 200 {object} entities.SubCategory
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /subcategories [post]
func (scc SubCategoryController) CreateAsync(c *gin.Context) {
	var obj entities.SubCategory
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err := scc.scs.CreateAsync(obj); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

// DeleteSubCategory godoc
// @Summary Delete Sub Category
// @Description Delete Sub Category
// @Tags subcategories
// @produce json
// @Param id path int true "data"
// @Success 200 {object} entities.SubCategory
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /subcategories/{id} [delete]
func (scc SubCategoryController) DeleteAsync(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err = scc.scs.DeleteAsync(id); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, helpers.ErrResponse{Message: "Successfully deleted record"})
}

// EditSubCategory godoc
// @Summary Edit Sub Category
// @Description Edit Sub Category
// @tags subcategories
// @Accept json
// @Produe json
// @Param id path int true "data"
// @Param data body entities.SubCategory true "data"
// @Success 200 {object} entities.SubCategory
// @Failure 404 {object} helpers.ErrResponse "Page not found"
// @Failure 500 {object} helpers.ErrResponse "Internal Server Error"
// @Router /subcategories/{id} [put]
func (scc SubCategoryController) EditAsync(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	var updatedObj entities.SubCategory
	if err := c.ShouldBindJSON(&updatedObj); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrResponse{Message: err.Error()})
		return
	}

	if err := scc.scs.EditAsync(id, &updatedObj); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedObj)
}