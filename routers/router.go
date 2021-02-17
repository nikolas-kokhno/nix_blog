package routers

import (
	"github.com/labstack/echo"
	"github.com/nikolas-kokhno/nix_blog/handlers"
)

func InitRoutes(e *echo.Echo) {
	/* Create public group */
	p := e.Group("/api/v1")

	/* Auth routes */

	/* Post routes */
	p.GET("/posts", handlers.GetAllPosts)
	p.GET("/posts/:id", handlers.GetPostByID)

	/* Comment routers */
	p.GET("/comments", handlers.GetAllComments)
}
