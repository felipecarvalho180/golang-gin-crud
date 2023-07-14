package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students/", controllers.FetchStudents)
	r.GET("/students/:ID", controllers.FetchStudent)
	r.GET("/students/name/:name", controllers.FetchStudentByName)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:ID", controllers.DeleteStudent)
	r.PATCH("/students/:ID", controllers.EditStudent)
	r.Run()
}
