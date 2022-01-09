package main

import (
	"log"
	"net/http"

	"Employee_project/controllers"
	"Employee_project/database"
	"Employee_project/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {

	router.HandleFunc("/createemployee", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/getemployees", controllers.GetAllEmployees).Methods("GET")
	router.HandleFunc("/getemployeebyid/{id}", controllers.GetEmployeeByID).Methods("GET")
	router.HandleFunc("/updateemployee/{id}", controllers.UpdateEmployeeByID).Methods("PUT")
	router.HandleFunc("/deleteemployee/{id}", controllers.DeletEmployeeByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User: "root",

			DB: "test",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Employee{})
}
