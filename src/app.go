package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api-4/config/db"
	"github.com/rudikurniawan99/go-api-4/src/registry"
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
	s.httpServer.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET"},
		AllowCredentials: true,
	}))

	registry.Registry(s.httpServer, s.database)

	s.httpServer.Run(":8082")
}
