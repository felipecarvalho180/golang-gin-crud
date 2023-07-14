package controllers

import (
	"fmt"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func FetchStudent(c *gin.Context) {
	ID := c.Params.ByName("ID")

	var student models.Student

	database.DB.First(&student, ID)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	ID := c.Params.ByName("ID")

	database.DB.Delete(&student, ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func EditStudent(c *gin.Context) {
	var student models.Student
	ID := c.Params.ByName("ID")

	database.DB.First(&student, ID)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func FetchStudentByName(c *gin.Context) {
	var student models.Student
	name := c.Param("name")
	fmt.Println(name)
	database.DB.Where(&models.Student{Name: name}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
