package daos

import (
	"encoding/json"
	"rahulProj/student/beans"
	"rahulProj/student/resources/database"
)

func FetchStudentDetails() []beans.Student {
	var StudentInfo []beans.Student

	database.DB.Find(&StudentInfo)
	// fmt.Println("Get All rows from db")
	// fmt.Println(StudentInfo)
	return StudentInfo
}

func FetchStudentById(id int) (beans.Student, error) {

	var StudentInfo beans.Student

	result := database.DB.Where("id = ?", id).First(&StudentInfo)

	return StudentInfo, result.Error
}

func InsertNewStudent(requestBody []byte) (beans.Student, error) {

	var StudentInfo beans.Student

	json.Unmarshal(requestBody, &StudentInfo)

	result := database.DB.Create(&StudentInfo)

	return StudentInfo, result.Error
}

func DeleteStudent(id int) (beans.Student, error) {

	var StudentInfo beans.Student

	result := database.DB.Where("id = ?", id).Delete(&StudentInfo)

	return StudentInfo, result.Error

}
