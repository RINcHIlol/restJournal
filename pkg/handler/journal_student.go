package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getStudentGrades(c *gin.Context) {
	var id int

	teacherId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	role, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if role != "teacher" {
		id, err = getUserId(c)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		id, err = strconv.Atoi(c.Param("id_student"))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	grades, err := h.services.JournalStudents.GetStudentGrades(id, teacherId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, grades)
}

func (h *Handler) updateStudentGrade(c *gin.Context) {

}

func (h *Handler) addStudentGrade(c *gin.Context) {

}

func (h *Handler) deleteStudentGrade(c *gin.Context) {

}
