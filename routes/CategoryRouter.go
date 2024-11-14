package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-sinan-akbar/controllers"
)

type CategoryRouterControler struct {
	categoryController controllers.CategoryController
}

func NewCategoryRouterController(cc controllers.CategoryController) CategoryRouterControler {
	return CategoryRouterControler{cc}
}

func (brc *CategoryRouterControler) CategoryRouter(rg *gin.RouterGroup) {
	rg.GET("/categories", brc.categoryController.GetAll)
	rg.GET("/categories/:id", brc.categoryController.GetByID)
	rg.POST("/categories", brc.categoryController.CreateAsync)
	rg.PUT("/categories/:id", brc.categoryController.EditAsync)
	rg.DELETE("/categories/:id", brc.categoryController.DeleteAsync)
}

