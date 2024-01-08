package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/School_MiniApi/Server/database"
	"github.com/teerapoom/School_MiniApi/Server/model"
)

// & คือ  address * คือ value
func CreateStudent(c *gin.Context) {
	var student model.School
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var duplicateStudent model.School
	database.Db.Where("fullname = ?", student.Fullname).First(&duplicateStudent)
	if duplicateStudent.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 400, "Messenger": "Student Duplicate"})
		return
	}

	newStudent := model.School{
		Fullname:  student.Fullname,
		StudentId: student.StudentId,
		ClassRoom: student.ClassRoom,
		Gender:    student.Gender,
	}

	_, err := newStudent.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": 201, "Messenger": "Method Post Successfully"})
}

func ViewAll(c *gin.Context) {
	var studentAll []model.School
	err := model.GetAllStuder(&studentAll)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Student": studentAll})
}

func ViewById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student model.School
	err := model.GetByIdStuder(&student, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStuder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student model.School
	err := model.GetByIdStuder(&student, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&student)
	err = model.Update(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func RemoveStuder(c *gin.Context) {
	var student model.School
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetByIdStuder(&student, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = model.DleStuder(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Messenger": "Successfully Remove"})
}
