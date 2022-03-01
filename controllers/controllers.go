package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"rahulProj/student/daos"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StudentDetails(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	StudentInfo := daos.FetchStudentDetails()

	// json.NewEncoder(c.Writer).Encode(StudentInfo)
	c.IndentedJSON(200, StudentInfo)
}

func StudentByID(c *gin.Context) {

	id := c.Param("id")

	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("Eroor while converting id string to int")
	}

	StudentInfo, err := daos.FetchStudentById(int(id_int))

	if err != nil {
		fmt.Println("error getting value by ID")
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": " cannot find student by ID " + id,
		})
	} else {
		c.IndentedJSON(200, StudentInfo)
	}

}

func RegisterStudent(c *gin.Context) {

	requestBody, _ := ioutil.ReadAll(c.Request.Body)

	StudentInfo, err := daos.InsertNewStudent(requestBody)

	if err != nil {
		fmt.Println("db connection error while registering student ")
		fmt.Println(err)
		c.IndentedJSON(400, "Registration unsuccesful")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": StudentInfo.Name + " Registered Succesfully",
		})
	}

}

func UnregisterStudent(c *gin.Context) {

	id := c.Param("id")

	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("Eroor while converting id string to int")
	}

	StudentInfo, err := daos.DeleteStudent(int(id_int))

	if err != nil {
		fmt.Println("Error while unregistering student")
		fmt.Println(err)
		c.IndentedJSON(400, "Unregistration unsuccesful")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": StudentInfo.Name + " Unregistered Succesfully",
		})
	}

}
