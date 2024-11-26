package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-sinan-akbar/controllers"
)

type SubCategoryRouterController struct {
	subCategoryRouterController controllers.SubCategoryController
}

func NewSubCategoryRouterController(scc controllers.SubCategoryController) SubCategoryRouterController {
	return SubCategoryRouterController{scc}
}

func (scc *SubCategoryRouterController) SubCategoryRouter(rg *gin.RouterGroup) {
	rg.GET("/subcategories", scc.subCategoryRouterController.GetAll)
	rg.GET("/subcategories/:id", scc.subCategoryRouterController.GetByID)
	rg.POST("/subcategories", scc.subCategoryRouterController.CreateAsync)
	rg.PUT("/subcategories/:id", scc.subCategoryRouterController.EditAsync)
	rg.DELETE("/subcategories/:id", scc.subCategoryRouterController.DeleteAsync)
}