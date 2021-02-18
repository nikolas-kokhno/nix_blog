package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func CreateNewPost(c echo.Context) error {
	postModel := new(models.Posts)
	if err := c.Bind(postModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: err.Error()})
	}

	/* Validate required request field */
	if postModel.Title == "" || postModel.Body == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <title>, <body> are required",
		})
	}

	/* Validate characters request field */
	if len(postModel.Title) < 3 || len(postModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <title>, <body> must be more than 2 characters",
		})
	}

	models.DB.Create(&postModel)
	return c.JSON(http.StatusOK, postModel)
}

func UpdatePostByID(c echo.Context) error {
	postModel := new(models.Posts)
	if err := c.Bind(postModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if postModel.Title == "" || postModel.Body == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <title>, <body> are required",
		})
	}

	/* Validate characters request field */
	if len(postModel.Title) < 3 || len(postModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <title>, <body> must be more than 2 characters",
		})
	}

	/* Checking if id exists in the db  */
	if err := models.DB.Where("id = ?", c.Param("id")).First(&postModel).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	models.DB.Model(&postModel).Update(models.Posts{Title: postModel.Title, Body: postModel.Body})
	return c.JSON(http.StatusOK, postModel)
}

func DeletePostByID(c echo.Context) error {
	postModel := new(models.Posts)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&postModel).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	models.DB.Delete(&postModel)

	return c.JSON(http.StatusOK, ErrorResponse{
		Status:  "success",
		Message: "Post deleted successfully!",
	})
}
