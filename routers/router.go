package routers

import (
	_ "github.com/nikolas-kokhno/nix_blog/docs"
	"github.com/nikolas-kokhno/nix_blog/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {
	/* Swagger documentation */
	e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)

	/* Create public group */
	p := e.Group("/api/v1")

	/* Auth routes */
	p.POST("/auth/login", handlers.Login)
	p.POST("/auth/signup", handlers.SignUp)

	/* Post routes */
	p.GET("/posts", handlers.GetAllPosts)
	p.GET("/posts/:id", handlers.GetPostByID)
	p.POST("/posts", handlers.CreateNewPost)
	p.PUT("/posts/:id", handlers.UpdatePostByID)
	p.DELETE("/posts/:id", handlers.DeletePostByID)

	/* Comment routers */
	p.GET("/comments", handlers.GetAllComments)
	p.GET("/comments/:id", handlers.GetCommentByID)
	p.POST("/comments", handlers.CreateNewComment)
	p.PUT("/comments/:id", handlers.UpdateCommentByID)
	p.DELETE("/comments/:id", handlers.DeleteCommentByID)
}
