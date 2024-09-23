package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getGradesGroup(c *gin.Context) {
	role, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if role != "teacher" {
		newErrorResponse(c, http.StatusForbidden, "You are not the teacher")
		return
	}

	id, err := strconv.Atoi(c.Param("id_group"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	subject, err := strconv.Atoi(c.Param("subject"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	students, err := h.services.JournalGroups.GetAll(id, subject)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, students)
}
