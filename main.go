package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/pt-sinan-akbar/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	// "github.com/pt-sinan-akbar/helpers"
	"log"
	"github.com/pt-sinan-akbar/services"
	"github.com/pt-sinan-akbar/controllers"
	"github.com/pt-sinan-akbar/initializers"
	"github.com/pt-sinan-akbar/routes"
	// "net/http"
)

//	@title			Go Finance API
//	@version		1.0
//	@description	Go Finance Swagger Documentation
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/


var (
	server *gin.Engine

	//Controllers
	CategoryController controllers.CategoryController
	SubCategoryController controllers.SubCategoryController

	//Routes
	CategoryRouterController routes.CategoryRouterControler
	SubCategoryRouterController routes.SubCategoryRouterController
	//Services
	CategoryService services.CategoryService
	SubCategoryService services.SubCategoryService
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables")
	}

	initializers.ConnectDB(&config)

	//Services
	CategoryService = services.NewCategoryService(initializers.DB)
	SubCategoryService = services.NewSubCategoryService(initializers.DB)
	//Controllers
	CategoryController = controllers.NewCategoryController(&CategoryService)
	SubCategoryController = controllers.NewSubCategoryController(&SubCategoryService)

	//Routes
	CategoryRouterController = routes.NewCategoryRouterController(CategoryController)
	SubCategoryRouterController = routes.NewSubCategoryRouterController(SubCategoryController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}


	server.Use(cors.New(corsConfig))
	router := server.Group("/api/v1")

	CategoryRouterController.CategoryRouter(router)
	SubCategoryRouterController.SubCategoryRouter(router)
	//branch init
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run("127.0.0.1:" + initializers.ConfigSetting.ServerPort))
}