package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorisationHeader = "Authorisation"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}
	idString, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid type of id")
		return "", errors.New("user id not found")
	}
	return idString, nil
}

func SendJSONResponse(c *gin.Context, key string, value interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		key: value,
	})
}
