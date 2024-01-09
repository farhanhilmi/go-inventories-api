package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"sbm-itb/config"
	"sbm-itb/httpserver/handler"
	"sbm-itb/httpserver/logger"
	"sbm-itb/httpserver/middleware"
	"sbm-itb/httpserver/router"
	"sbm-itb/repository"
	"sbm-itb/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(gin *gin.Engine, db *gorm.DB) {
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true

	gin.Use(middleware.Logger(logger.NewLogger()))
	gin.Use(middleware.ErrorHandler())
	gin.Use(cors.New(configCors))

	inventoriesRepo := repository.NewInventoriesRepository(db)

	inventoriesUsecase := usecase.NewInventoriesUsecase(inventoriesRepo)

	inventoriesHandler := handler.NewInventoriesHandler(inventoriesUsecase)

	apiGroup := gin.Group("api")

	router.NewInventoriesRouter(inventoriesHandler, gin, apiGroup)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetEnv("PORT")),
		Handler: gin,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
	}
}
