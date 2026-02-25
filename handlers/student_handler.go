package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format or validation failed"})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.UpdateStudent(id, student)
	if err != nil {
		if err.Error() == "student not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		if err.Error() == "student not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
