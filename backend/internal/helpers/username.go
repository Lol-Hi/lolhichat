package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUsername = "";

func GetContextUsername(c *gin.Context) (string, error) {
	user, exists := c.Get(ContextUsername);
	if !exists {
		return "", errors.New("No username in context")
	}
	username, ok := user.(string)
	if !ok {
		return "", errors.New("Context username of invalid type")
	}

	return username, nil;
}

