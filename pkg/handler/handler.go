package handler

import (
	"net/http"

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

	// for save images files
	router.Static("/images", "./images")

	router.LoadHTMLGlob("templates/*")
	router.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
		  "title": "Main website",
		})
	})

	// redirect on swagger ui dashboard
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.signIn)
			auth.GET("/me",  h.userIdentity, h.Me)
		}

		api.POST("/upload", h.userIdentity, h.UploadImage)
		api.GET("/mail", h.SendEmail)

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
