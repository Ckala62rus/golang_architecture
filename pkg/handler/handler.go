package handler

import (
	_ "github.com/Ckala62rus/go/docs" // docs folder
	"github.com/Ckala62rus/go/pkg/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		users := api.Group("/users", h.userIdentity)
		{
			users.GET("/", h.GetAllUsers)
			users.GET("/user/:name", h.GetUserByName)
			users.GET("/:id", h.GetById)
			users.POST("/", h.CreateUser)
			users.DELETE(":id", h.DeleteUserById)
			users.PUT(":id", h.UpdateUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
