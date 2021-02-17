package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nikolas-kokhno/nix_blog/models"
)

func GetPostByID(c echo.Context) error {
	postModel := new(models.Posts)

	if err := models.DB.Where("id = ?", c.Param("id")).First(&postModel).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   postModel,
	})
}

func GetAllPosts(c echo.Context) error {
	postModel := new([]models.Posts)
	user_id := c.QueryParam("user_id")

	/* Filter by user_id */
	if user_id != "" {
		if err := models.DB.First(&postModel, "user_id = ?", user_id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, SuccessResponse{
			Status: "success",
			Data:   postModel,
		})
	}

	models.DB.Find(&postModel)

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   postModel,
	})
}
