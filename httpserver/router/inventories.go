package router

import (
	"sbm-itb/httpserver/handler"

	"github.com/gin-gonic/gin"
)

func NewInventoriesRouter(h *handler.InventoriesHandler, gin *gin.Engine, group *gin.RouterGroup) *gin.Engine {
	group.POST("/inventories", h.Create)
	group.GET("/inventories", h.FindAll)
	group.GET("/inventories/:barangId", h.FindByID)
	group.PUT("/inventories/:barangId", h.UpdateByID)
	group.DELETE("/inventories/:barangId", h.DeleteByID)

	return gin
}
