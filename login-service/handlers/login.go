package handlers

import (
	"net/http"

	"github.com/PaulTaranu/CarTrack/login-service/auth"
	"github.com/PaulTaranu/CarTrack/login-service/models"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Request"})
	}

	user, err := models.FindUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal error"})
	}

	if user == nil || !auth.CheckPassword(req.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	token, err := auth.GenerateToken(user.ID, user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "JWT generation failed"})
	}
	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}
