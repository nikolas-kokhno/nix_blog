package handlers

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/nikolas-kokhno/nix_blog/models"
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

// @Summary Create new comments
// @Tags Comments
// @Security userLogin
// @Description created new comments
// @ModuleID createNewComments
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments [post]
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

// @Summary Update comment by ID
// @Tags Comments
// @Security userLogin
// @Description updated comment data
// @ModuleID updateCommentByID
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /comments/{id} [put]
func UpdateCommentByID(c echo.Context) error {
	commentModel := new(models.Comments)
	if err := c.Bind(commentModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if commentModel.Name == "" || commentModel.Email == "" || commentModel.Body == "" || commentModel.PostId <= 0 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <name>, <email>, <body>, <user_id> are required",
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

	if err := models.DB.Model(&commentModel).Update(models.Comments{Name: commentModel.Name, Email: commentModel.Email, Body: commentModel.Body}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commentModel)
}

// @Summary Delete comment by ID
// @Tags Comments
// @Security userLogin
// @Description deleted comment data
// @ModuleID deleteCommentByID
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
