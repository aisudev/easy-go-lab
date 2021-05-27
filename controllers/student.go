package controllers

import (
	"echo-lab-go/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateStudentScore ..
func CreateStudentScore(c echo.Context) error {

	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	if err := models.SQLiteDB.Create(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Create StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetStudentScore ..
func GetStudentScore(c echo.Context) error {

	name := c.Param("name")
	fmt.Println(name)
	studentModel := models.StudentScore{}
	if err := models.SQLiteDB.Debug().Where("name = ?", name).First(&studentModel).Error; err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get StudentScore error"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": studentModel})
}

// GetAllStudentScore ..
func GetAllStudentScore(c echo.Context) error {
	studentModel := []models.StudentScore{}
	if err := models.SQLiteDB.Find(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get All StudentScore error"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": studentModel})
}

// UpdateStudentScore ..
func UpdateStudentScore(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)

	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	if err := models.SQLiteDB.Debug().Where("id = ?", id).Updates(studentModel).Error; err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("studentModel: ", studentModel)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// DeleteStudentScore ..
func DeleteStudentScore(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	if err := models.SQLiteDB.Debug().Where("id = ?", id).Delete(&studentModel).Error; err != nil {
		fmt.Println(err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// RecoverStudentScore ..
func RecoverStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := models.SQLiteDB.Model(&studentModel).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		fmt.Println(err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

type grade struct {
	Name    string  `json:"name"`
	Subject string  `json:"subject"`
	Score   float64 `json:"-"`
	Grade   string  `json:"grade"`
}

// GetSubjectGrade ..
func GetSubjectGrade(c echo.Context) error {
	gradeModel := []grade{}
	// data := map[string]interface{}{}
	if err := models.SQLiteDB.Debug().Table("student_scores").Select("name", "subject", "AVG(score) as score").Group("name").Order("subject").Scan(&gradeModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get All StudentScore error"})
	}

	for i, v := range gradeModel {
		var g string
		if v.Score <= 50 {
			g = "F"
		} else if v.Score <= 60 {
			g = "D"
		} else if v.Score <= 70 {
			g = "C"
		} else if v.Score <= 80 {
			g = "B"
		} else {
			g = "A"
		}
		gradeModel[i].Grade = g
	}

	fmt.Println(gradeModel)

	// Fix grade: A: 81-100, B: 71-80, C: 61-70, D: 51-60, F: 0-50
	/* expected response:
	[
		{
			"name": "Tong",
			"subject": "eng",
			"grade": "A"
		},
		{
			"name": "Tong",
			"subject": "math",
			"grade": "B"
		},
		{
			"name": "Tong2",
			"subject": "eng",
			"grade": "C"
		}
	]
	*/

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": gradeModel})
}

// GetSubjectGradeByStudentName ..
func GetSubjectGradeByStudentName(c echo.Context) error {
	name := c.Param("name")
	gradeModel := []grade{}
	// data := map[string]interface{}{}
	if err := models.SQLiteDB.Debug().Table("student_scores").Select("name", "subject", "AVG(score) as score").Where("name = ?", name).Group("subject").Scan(&gradeModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get All StudentScore error"})
	}

	resMap := map[string]string{}

	for _, v := range gradeModel {
		var g string
		if v.Score <= 50 {
			g = "F"
		} else if v.Score <= 60 {
			g = "D"
		} else if v.Score <= 70 {
			g = "C"
		} else if v.Score <= 80 {
			g = "B"
		} else {
			g = "A"
		}
		resMap[v.Subject] = g
	}

	/* expected response:
	{
		"eng": "A",
		"math": "B",
		"social": "C"
	}
	*/

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": resMap})
}
