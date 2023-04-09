package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jorgemontess/go-gorm-restapi/models"
	"github.com/jorgemontess/go-gorm-restapi/db"
	"github.com/jorgemontess/go-gorm-restapi/routes"
)

//Installing go air go install github.com/cosmtrek/air@latest
//Para no tener que cancelar la conexion cada vez que hagamos cambios en el archivo
// se ejecuta con air init


func main(){
    fmt.Println("Hola")


    db.DBConnection()
    db.DB.AutoMigrate(models.Task{})
    db.DB.AutoMigrate(models.User{})



    r := mux.NewRouter()

    //---- HANDLERS ----

    r.HandleFunc("/", routes.HomeHandler)


    //{id} reciben el parametro
    r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
    r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
    r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
    r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
    


    http.ListenAndServe(":3000", r)
}
