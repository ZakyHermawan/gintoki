package main

import (
	"gintoki/controller"
	"gintoki/middlewares"
	"gintoki/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

var (
	VideoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(VideoService)

	LoginService    service.LoginService       = service.NewLoginService()
	JwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(LoginService, JwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "input valid"})
			}
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run(":" + port)
}
