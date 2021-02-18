package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/nikolas-kokhno/nix_blog/models"
	"github.com/spf13/viper"
)

func generateToken(username string) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token
}

// @Summary User login
// @Tags Auth
// @Description user sign in
// @ModuleID userLogin
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/login [post]
func Login(c echo.Context) error {
	userModel := new(models.Users)
	if err := c.Bind(userModel); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if userModel.Username == "" || userModel.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <username>, <password> are required",
		})
	}

	/* Validate characters request field */
	if len(userModel.Username) < 3 || len(userModel.Password) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <username>, <password> must be more than 2 characters",
		})
	}

	if !models.DB.Where("username = ? AND password = ?", userModel.Username, userModel.Password).Find(&models.Users{}).RecordNotFound() {
		token := generateToken(userModel.Username)

		t, err := token.SignedString([]byte(viper.GetString("secretJWT")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// @Summary User signup
// @Tags Auth
// @Description user sign up
// @ModuleID userSignup
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/signup [post]
func SignUp(c echo.Context) error {
	userModel := new(models.Users)
	if err := c.Bind(userModel); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if userModel.Username == "" || userModel.Password == "" || userModel.Name == "" || userModel.Email == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <username>, <password>, <name>, <email> are required",
		})
	}

	/* Validate characters request field */
	if len(userModel.Username) < 3 || len(userModel.Password) < 3 || len(userModel.Name) < 3 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Fields: <username>, <password>, <name> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(userModel.Email) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Field <email> is not valid. For example: test@test.com",
		})
	}

	if err := models.DB.Create(&userModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	token := generateToken(userModel.Username)
	t, err := token.SignedString([]byte(viper.GetString("secretJWT")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
