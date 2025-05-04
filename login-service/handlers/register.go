package handlers

import (
	"log"
	"net/http"

	"github.com/PaulTaranu/CarTrack/login-service/auth"
	"github.com/PaulTaranu/CarTrack/login-service/models"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

func RegisterHandler(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		log.Fatal("Invalid registration request")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid registration request"})
	}

	if req.Email == "" || req.Password == "" {
		log.Fatal("Empty email or password")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Empty email or password"})
	}

	if user, _ := models.FindUserByEmail(req.Email); user != nil {
		log.Fatal("User already exists")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "User already exists"})
	}

	safePassword, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Fatal("Password hashing failed")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Password hashing failed"})
	}

	newUser := models.User{
		Email:    req.Email,
		Password: safePassword,
	}

	if err := models.CreateUser(&newUser); err != nil {
		log.Fatal("Creation failed")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Creation failed"})
	}

	return c.JSON(http.StatusOK, RegisterResponse{ID: newUser.ID})
}
