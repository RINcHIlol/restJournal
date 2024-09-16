package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userIdCtx   = "userId"
	userRoleCtx = "userRole"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid header")
		return
	}

	userId, userRole, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtx, userId)
	c.Set(userRoleCtx, userRole)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "could not get user id")
		return 0, errors.New("user id not found")
	}

	return id.(int), nil
}

func getUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(userRoleCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "could not get user role")
		return "", errors.New("user role not found")
	}

	return role.(string), nil
}
