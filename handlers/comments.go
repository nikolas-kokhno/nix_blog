package handlers

import (
	"net/http"
	"regexp"

	"github.com/nikolas-kokhno/nix_blog/models"

	"github.com/labstack/echo/v4"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// @Summary Get comment by ID
// @Tags Comments
// @Description returning comment data by ID
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} models.Comments
// @Router /comments/{id} [get]
func GetCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)

	if err := models.DB.Where("id = ?", c.Param("id")).First(&commentModel).Error; err != nil {
		return c.JSON(http.StatusNotFound, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   commentModel,
	})
}

// @Summary Get all comments
// @Tags Comments
// @Description returning all comments
// @ModuleID getAllComments
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments [get]
func GetAllComments(c echo.Context) error {
	commentsModel := new([]models.Comments)
	post_id := c.QueryParam("post_id")

	/* Filter by post_id */
	if post_id != "" {
		if err := models.DB.First(&commentsModel, "post_id = ?", post_id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, MessageResponse{
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

// @Summary Create new comments
// @Tags Comments
// @Security ApiKeyAuth
// @Description created new comments
// @ModuleID createNewComments
// @Accept  json
// @Produce  json
// @Param data body models.Comments true "Enter comment data to create a comment"
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments [post]
func CreateNewComment(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := c.Bind(commentModel); err != nil {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if commentModel.Name == "" || commentModel.Email == "" || commentModel.Body == "" {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> are required",
		})
	}

	/* Validate characters request field */
	if len(commentModel.Name) < 3 || len(commentModel.Email) < 3 || len(commentModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(commentModel.Email) {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "email is not valid",
		})
	}

	models.DB.Create(&commentModel)
	return c.JSON(http.StatusCreated, commentModel)
}

// @Summary Update comment by ID
// @Tags Comments
// @Security ApiKeyAuth
// @Description updated comment data
// @ModuleID updateCommentByID
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param data body models.Comments true "Enter comment data to create a comment"
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments/{id} [put]
func UpdateCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)
	commentsID := c.Param("id")
	if err := c.Bind(commentModel); err != nil {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if commentModel.Name == "" || commentModel.Email == "" || commentModel.Body == "" || commentModel.PostId <= 0 {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body>, <user_id> are required",
		})
	}

	/* Validate characters request field */
	if len(commentModel.Name) < 3 || len(commentModel.Email) < 3 || len(commentModel.Body) < 3 {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(commentModel.Email) {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "email is not valid",
		})
	}

	/* Checking if id exists in the db  */
	commetsData := models.Comments{}
	if err := models.DB.Where("id = ?", commentsID).First(&commetsData).Error; err != nil {
		return c.JSON(http.StatusNotFound, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	models.DB.Model(&models.Comments{}).Where("id = ?", commentsID).Updates(models.Comments{Body: commentModel.Body})

	return c.JSON(http.StatusOK, MessageResponse{
		Status:  "success",
		Message: "Comment updated successfully!",
	})
}

// @Summary Delete comment by ID
// @Tags Comments
// @Security ApiKeyAuth
// @Description deleted comment data
// @ModuleID deleteCommentByID
// @Param id path int true "Comment ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments/{id} [delete]
func DeleteCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&commentModel).Error; err != nil {
		return c.JSON(http.StatusNotFound, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	models.DB.Delete(&commentModel)

	return c.JSON(http.StatusOK, MessageResponse{
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
