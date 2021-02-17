package routers

import (
	"github.com/labstack/echo"
	"github.com/nikolas-kokhno/nix_blog/handlers"
)

func InitRoutes(e *echo.Echo) {
	/* Create public group */
	p := e.Group("/api/v1")

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
