package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/afifialaa/employee-management/models"
)

// Insert new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{
		FirstName:     r.FormValue("firstName"),
		LastName:      r.FormValue("lastName"),
		Email:         r.FormValue("email"),
		PhoneNumber:   r.FormValue("phoneNumber"),
		Gender:        r.FormValue("gender"),
		JobTitle:      r.FormValue("jobTitle"),
		Department:    r.FormValue("department"),
		Country:       r.FormValue("country"),
		City:          r.FormValue("city"),
		StreetAddress: r.FormValue("streetAddress"),
		University:    r.FormValue("university"),
		Salary:        r.FormValue("salary"),
	}

	result := employee.CreateEmployee()
	if !result {
		res := map[string]string{"err": "Failed to create employee"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Employee created successfully"}
	json.NewEncoder(w).Encode(res)
	return
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{
		FirstName:     r.FormValue("firstName"),
		LastName:      r.FormValue("lastName"),
		Email:         r.FormValue("email"),
		PhoneNumber:   r.FormValue("phoneNumber"),
		Gender:        r.FormValue("gender"),
		JobTitle:      r.FormValue("jobTitle"),
		Department:    r.FormValue("department"),
		Country:       r.FormValue("country"),
		City:          r.FormValue("city"),
		StreetAddress: r.FormValue("streetAddress"),
		University:    r.FormValue("university"),
		Salary:        r.FormValue("salary"),
	}

	result := employee.DeleteEmployee()
	if !result {
		res := map[string]string{"err": "Failed to delete employee"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Employee deleted successfully"}
	json.NewEncoder(w).Encode(res)
	return
}
