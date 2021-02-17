package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nikolas-kokhno/nix_blog/models"
)

func GetCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)

	if err := models.DB.Where("id = ?", c.Param("id")).First(&commentModel).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   commentModel,
	})
}

func GetAllComments(c echo.Context) error {
	commentsModel := new([]models.Comments)
	post_id := c.QueryParam("post_id")

	/* Filter by post_id */
	if post_id != "" {
		if err := models.DB.First(&commentsModel, "post_id = ?", post_id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, SuccessResponse{
			Status: "success",
			Data:   commentsModel,
		})
	}

	models.DB.Find(&commentsModel)

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   commentsModel,
	})
}
