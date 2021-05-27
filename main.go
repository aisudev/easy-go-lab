package main

import (
	"net/http"

	"echo-lab-go/controllers"
	"echo-lab-go/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := models.InitialDatabase(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", hello)

	e.POST("/student", controllers.CreateStudentScore)
	e.GET("/student/:name", controllers.GetStudentScore)
	e.GET("/student", controllers.GetAllStudentScore)
	e.PUT("/student/update/:id", controllers.UpdateStudentScore)
	e.DELETE("/student/delete/:id", controllers.DeleteStudentScore)
	e.PATCH("/student/recover/:id", controllers.RecoverStudentScore)

	e.GET("/subject/grade/", controllers.GetSubjectGrade)
	e.GET("/subject/grade/:name", controllers.GetSubjectGradeByStudentName)

	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "hello world!"})
}
