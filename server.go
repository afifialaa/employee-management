package main

import (
	"fmt"
	"net/http"

	"github.com/afifialaa/employee-management/config"
	"github.com/afifialaa/employee-management/handlers"
)

type Status struct {
	msg string
}

func main() {
	config.DBConnect()

	// routes
	http.HandleFunc("/employee/createEmployee", handlers.CreateEmployee)
	http.HandleFunc("/employee/deleteEmployee", handlers.DeleteEmployee)

	http.HandleFunc("/dashboard/getUsersNum", handlers.GetUsersNum)
	http.HandleFunc("/dashboard/getEmployeesNum", handlers.GetEmployeesNum)

	// listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)
}
