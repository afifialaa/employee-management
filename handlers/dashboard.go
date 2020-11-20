package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/afifialaa/employee-management/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsersNum(w http.ResponseWriter, r *http.Request) {
	number, err := config.UserCollection.CountDocuments(context.TODO(), bson.M{"": ""})
	if err != nil {
		log.Fatal(err)
	}

	res := map[string]int64{"number": number}
	json.NewEncoder(w).Encode(res)
	return
}

func GetEmployeesNum(w http.ResponseWriter, r *http.Request) {
	number, err := config.EmployeeCollection.CountDocuments(context.TODO(), bson.M{"": ""})
	if err != nil {
		log.Fatal(err)
	}

	res := map[string]int64{"number": number}
	json.NewEncoder(w).Encode(res)
	return
}
