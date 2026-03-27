package main

import (
	"fmt"
	"go-student-manager/handlers"
	"go-student-manager/services"
)

func main() {
	fmt.Println("CHO AN CUT")
	services.ShowStudent()
	handlers.Menu()
	services.Addstudent()
}
