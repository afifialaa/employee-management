package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/afifialaa/employee-management/config"
)

type Employee struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Gender        string `json:"gender"`
	JobTitle      string `json:"job_title"`
	Department    string `json:"department,omitempty"`
	Country       string `json:"country"`
	City          string `json:"city,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	University    string `json:"university,omitempty"`
	Salary        string `json:"salary,omitempty"`
}

// Create new employee
func (emp Employee) CreateEmployee() bool {
	_, err := config.EmployeeCollection.InsertOne(context.TODO(), emp)
	if err != nil {
		fmt.Println("mongodb error ", err.Error())
		return false
	}

	fmt.Println("Employee was created")
	return true
}

func (emp Employee) DeleteEmployee() bool {
	filter := bson.D{{"email", emp.Email}}

	_, err := config.EmployeeCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to delete employee")
		return false
	}

	return true
}
