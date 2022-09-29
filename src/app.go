package src

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/config/db"
	"gorm.io/gorm"
)

type server struct {
	httpServer *gin.Engine
	database   *gorm.DB
}

func InitServer() *server {
	e := gin.Default()
	db := db.InitGorm()

	return &server{
		httpServer: e,
		database:   db,
	}
}

func (s *server) Run() {
	s.httpServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "api test success",
		})
	})
	s.httpServer.Run(":8082")
}
