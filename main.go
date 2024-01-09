package main

import (
	"sbm-itb/config"
	"sbm-itb/database"
	"sbm-itb/httpserver"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := database.GetInstance()

	gin := gin.Default()
	httpserver.Start(gin, db)
}
