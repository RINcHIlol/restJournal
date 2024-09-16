package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func (h *Handler) getAllStudents(c *gin.Context) {
//	userRole, err := getUserRole(c)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if userRole == "student" {
//		newErrorResponse(c, http.StatusBadRequest, "You can't get all students")
//		return
//	}
//
//	students, err := h.services.Students.GetAll()
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, students)
//}
//
//func (h *Handler) getStudentById(c *gin.Context) {
//	var studentId int
//
//	userRole, err := getUserRole(c)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	if userRole == "student" {
//		studentId, err = getUserId(c)
//		if err != nil {
//			newErrorResponse(c, http.StatusInternalServerError, err.Error())
//			return
//		}
//	} else {
//		studentId, err = strconv.Atoi(c.Param("id"))
//		if err != nil {
//			newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
//			return
//		}
//	}
//
//	student, err := h.services.Students.GetById(studentId)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, student)
//}

func (h *Handler) getAllStudents(c *gin.Context) {
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole == "student" {
		studentId, err := getUserId(c)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		student, err := h.services.Students.GetById(studentId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, student)
		return
	}

	students, err := h.services.Students.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *Handler) getStudentById(c *gin.Context) {
	var studentId int

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if userRole == "student" {
		newErrorResponse(c, http.StatusInternalServerError, "idi nahuy shket")
		return
	}

	studentId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	student, err := h.services.Students.GetById(studentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, student)
}
