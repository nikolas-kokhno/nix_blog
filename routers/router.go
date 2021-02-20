package routers

import (
	_ "github.com/nikolas-kokhno/nix_blog/docs"
	"github.com/nikolas-kokhno/nix_blog/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {
	/* Create middleware */
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	/* Swagger documentation */
	e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)

	/* Create protected group */
	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(viper.GetString("secretJWT"))))

	/* Create public group */
	p := e.Group("/api/v1")

	/* Auth routes */
	p.POST("/auth/login", handlers.Login)
	p.POST("/auth/signup", handlers.SignUp)

	/* Post routes */
	p.GET("/posts", handlers.GetAllPosts)
	p.GET("/posts/:id", handlers.GetPostByID)
	r.POST("/posts", handlers.CreateNewPost)
	r.PUT("/posts/:id", handlers.UpdatePostByID)
	r.DELETE("/posts/:id", handlers.DeletePostByID)

	/* Comment routers */
	p.GET("/comments", handlers.GetAllComments)
	p.GET("/comments/:id", handlers.GetCommentByID)
	r.POST("/comments", handlers.CreateNewComment)
	r.PUT("/comments/:id", handlers.UpdateCommentByID)
	r.DELETE("/comments/:id", handlers.DeleteCommentByID)
}
