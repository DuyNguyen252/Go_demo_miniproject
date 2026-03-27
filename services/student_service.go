package services

import (
	"fmt"
	"go-student-manager/data"
	"go-student-manager/models"
	"go-student-manager/utils"
)

const StorageDir = "output_data"

func Addstudent() {
	fullName := utils.InputString("Enter Full Name: ")
	age := utils.InputInt("Enter Age: ")
	class := utils.InputString("Enter Class: ")

	students := models.Students{
		FullName: fullName,
		Age:      age,
		Class:    class,
	}
	content := fmt.Sprintf("Name: %s, Age: %d, Class: %s",
		students.FullName, students.Age, students.Class)

	err := data.WriteInfoStudent(StorageDir, "student.txt", content)
	if err != nil {
		fmt.Println(err)
	}

}

func ShowStudent() {
	fmt.Println("Info Student: ")

	content, err := data.ReadInfoStudent(StorageDir, "student.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
