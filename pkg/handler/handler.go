package handler

import (
	"github.com/gin-gonic/gin"
	"rest_journal/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			students := lists.Group("/students")
			{
				students.GET("/", h.getAllStudents)
				students.GET("/:id", h.getStudentById)
			}

			teachers := lists.Group("/teachers")
			{
				teachers.GET("/", h.getAllTeachers)
				teachers.GET("/:id", h.getTeacherById)
			}

			groups := lists.Group("/groups")
			{
				groups.GET("/", h.getAllGroups)
				groups.GET("/:id", h.getGroupById)
				groups.GET("/:id/students", h.getAllStudentsByGroup)
			}

			specialties := lists.Group("/specialties")
			{
				specialties.GET("/", h.getAllSpecialties)
				specialties.GET("/:id", h.getSpecialtyById)
			}

			subjects := lists.Group("/subjects")
			{
				subjects.GET("/", h.getAllSubjects)
				subjects.GET("/:id", h.getSubjectById)
				subjects.GET("/specialties/:id", h.getSubjectsBySpecialty)
			}
		}
	}

	//journal := router.Group("/journal"){
	//	journal.
	//}

	return router
}
