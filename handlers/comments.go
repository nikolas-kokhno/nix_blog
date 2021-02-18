package handlers

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/nikolas-kokhno/nix_blog/models"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

func CreateNewComment(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := c.Bind(commentModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if commentModel.Name == "" || commentModel.Email == "" || commentModel.Body == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> are required",
		})
	}

	/* Validate characters request field */
	if len(commentModel.Name) < 3 || len(commentModel.Email) < 3 || len(commentModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(commentModel.Email) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "email is not valid",
		})
	}

	models.DB.Create(&commentModel)
	return c.JSON(http.StatusOK, commentModel)
}

func UpdateCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := c.Bind(commentModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if commentModel.Name == "" || commentModel.Email == "" || commentModel.Body == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> are required",
		})
	}

	/* Validate characters request field */
	if len(commentModel.Name) < 3 || len(commentModel.Email) < 3 || len(commentModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(commentModel.Email) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "email is not valid",
		})
	}

	models.DB.Model(&commentModel).Update(models.Comments{Name: commentModel.Name, Email: commentModel.Email, Body: commentModel.Body})
	return c.JSON(http.StatusOK, commentModel)
}

func DeleteCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&commentModel).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	models.DB.Delete(&commentModel)

	return c.JSON(http.StatusOK, ErrorResponse{
		Status:  "success",
		Message: "Comment deleted successfully!",
	})
}

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
