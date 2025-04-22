package routes

import (
	"book-store-server/config"
	"book-store-server/controllers"
	"book-store-server/internal/middleware"
	"book-store-server/internal/storage/DatabaseService"
	"book-store-server/services"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(config *config.Config, db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	databaseService := DatabaseService.NewDBService(db)

	appService := services.NewAppService(databaseService)
	appController := controllers.NewController(appService)

	router.Static("uploads", "./uploads")

	bookGroup := router.Group("/books")
	{
		bookGroup.POST("", appController.CreateBook)
		bookGroup.GET("/:id", appController.GetBook)
		bookGroup.GET("", appController.GetBooks)
		bookGroup.DELETE("/:id", appController.DeleteBook)
		bookGroup.PUT("/:id", appController.UpdateBook)
	}

	return router
}
