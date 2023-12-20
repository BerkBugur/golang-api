package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Tags Users
// @Summary Sign up a new user
// @Produce json
// @Param email formData string true "Email address"
// @Param password formData string true "Password"
// @Success 200 {object} gin.H
// @Router /users/signup [post]
func SignUp(c *gin.Context) {
	// Get the email/pass req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return

	}
	// Create user
	user := models.Users{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{})

}

// @Tags Users
// @Summary Log in a user
// @Produce json
// @Param email formData string true "Email address"
// @Param password formData string true "Password"
// @Success 200 {object} gin.H
// @Router /users/login [post]
func Login(c *gin.Context) {
	// Get the email/pass req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	// Look in db
	var user models.Users
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return

	}
	// Compare pass with hashed
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return

	}
	// Cookies for token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

// @Tags Users
// @Summary Validate the logged-in user
// @Produce json
// @Success 200 {object} gin.H
// @Security jwt
// @SecurityDefinitions jwt
// @Router /users/validate [get]
func Validate(c *gin.Context) {
	//user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in",
	})
}
