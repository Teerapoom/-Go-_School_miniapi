package model

import (
	"github.com/teerapoom/School_MiniApi/Server/database"
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Fullname  string `json:"fullname"`
	StudentId uint   `json:"student_id"`
	ClassRoom string `json:"class_room"`
	Gender    string `json:"gender"`
}

// * value , & address
func (school *School) Save() (*School, error) {
	err := database.Db.Create(&school).Error
	if err != nil {
		return &School{}, err
	}
	return school, nil
}

func GetAllStuder(School *[]School) (err error) {
	err = database.Db.Find(School).Error
	if err != nil {
		return err
	}
	return nil
}

func GetByIdStuder(student *School, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(student).Error
	if err != nil {
		return err
	}
	return nil
}

func Update(studer *School) (err error) {
	err = database.Db.Updates(studer).Error
	if err != nil {
		return err
	}
	return nil
}

func DleStuder(studer *School) error {
	if err := database.Db.Delete(&studer).Error; err != nil {
		return err
	}
	return nil
}
