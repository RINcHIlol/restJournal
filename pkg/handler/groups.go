package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllGroups(c *gin.Context) {
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole == "student" {
		newErrorResponse(c, http.StatusBadRequest, "idi nahuY <3")
		return
	}

	groups, err := h.services.Groups.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *Handler) getGroupById(c *gin.Context) {
	var groupId int

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole == "student" {
		newErrorResponse(c, http.StatusBadRequest, "idi nahuY <3")
		return
	}

	groupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	group, err := h.services.Groups.GetById(groupId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) getAllStudentsByGroup(c *gin.Context) {
	var groupId int

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole == "student" {
		newErrorResponse(c, http.StatusBadRequest, "idi nahuY <3")
		return
	}

	groupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	students, err := h.services.Groups.GetAllStudents(groupId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, students)
}
