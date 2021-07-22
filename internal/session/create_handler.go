package session

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/trackit/database/definitions"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func CreateHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		renderError(http.StatusUnauthorized, ErrInvalidCredentials)
		return
	}

	user, err := userLogin(username, password)
	if err != nil {
		renderError(http.StatusUnauthorized, ErrInvalidCredentials)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func renderError(statusCode int, err error) {
}

func userLogin(username string, password string) (*definitions.User, error) {
}
