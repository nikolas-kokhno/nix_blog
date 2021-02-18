package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikolas-kokhno/nix_blog/models"
)

// @Summary Get post by ID
// @Tags Posts
// @Description returning post data by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Posts
// @Router /posts/{id} [get]
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

// @Summary Get all posts
// @Tags Posts
// @Description returning all post
// @ModuleID getAllPosts
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /posts [get]
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

// @Summary Create new post
// @Tags Posts
// @Security userLogin
// @Description created new post
// @ModuleID createNewPost
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /posts [post]
func CreateNewPost(c echo.Context) error {
	postModel := new(models.Posts)
	if err := c.Bind(postModel); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: err.Error()})
	}

	/* Validate required request field */
	if postModel.Title == "" || postModel.Body == "" || postModel.UserID <= 0 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <title>, <body>, <user_id> are required",
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

// @Summary Update post by ID
// @Tags Posts
// @Security userLogin
// @Description updated post data
// @ModuleID updatePostByID
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /posts/{id} [put]
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

// @Summary Delete post by ID
// @Tags Posts
// @Security userLogin
// @Description deleted post data
// @ModuleID deletePostByID
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /posts/{id} [delete]
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
