package handlers

import (
	"net/http"
	"time"

	"github.com/nikolas-kokhno/nix_blog/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func generateToken(id int64, username, email string) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token
}

func decodeToken(jwtToken string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwtToken[7:], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secretJWT")), nil
	})

	return token, err
}

// @Summary User login
// @Tags Auth
// @Description user sign in
// @ModuleID userLogin
// @Accept json
// @Produce json
// @Param data body models.UserLogin true "Enter your username and password"
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/login [post]
func Login(c echo.Context) error {
	userModel := new(models.Users)
	if err := c.Bind(userModel); err != nil {
		return c.JSON(http.StatusInternalServerError, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if userModel.Username == "" || userModel.Password == "" {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <username>, <password> are required",
		})
	}

	/* Validate characters request field */
	if len(userModel.Username) < 3 || len(userModel.Password) < 3 {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <username>, <password> must be more than 2 characters",
		})
	}

	if !models.DB.Where("username = ? AND password = ?", userModel.Username, userModel.Password).Find(&userModel).RecordNotFound() {
		token := generateToken(userModel.ID, userModel.Username, userModel.Email)

		t, err := token.SignedString([]byte(viper.GetString("secretJWT")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, MessageResponse{
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
// @Param data body models.Users true "Enter your registration details"
// @Success 200 {object} SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/signup [post]
func SignUp(c echo.Context) error {
	userModel := new(models.Users)
	if err := c.Bind(userModel); err != nil {
		return c.JSON(http.StatusInternalServerError, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	/* Validate required request field */
	if userModel.Username == "" || userModel.Password == "" || userModel.Name == "" || userModel.Email == "" {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <username>, <password>, <name>, <email> are required",
		})
	}

	/* Validate characters request field */
	if len(userModel.Username) < 3 || len(userModel.Password) < 3 || len(userModel.Name) < 3 {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Fields: <username>, <password>, <name> must be more than 2 characters",
		})
	}

	/* Check for valid email */
	if !isEmailValid(userModel.Email) {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Status:  "error",
			Message: "Field <email> is not valid. For example: test@test.com",
		})
	}

	if err := models.DB.Create(&userModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	token := generateToken(userModel.ID, userModel.Username, userModel.Email)
	t, err := token.SignedString([]byte(viper.GetString("secretJWT")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, MessageResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
