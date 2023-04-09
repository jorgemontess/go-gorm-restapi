package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



//Conexion a Postgres en docker
//me aseguro de que el contenedor este iniciado
// si está en docker ps -a lo puedo inicar con "docker start <name_container>"
//Para ejecutar el contenedor hago "docker exec -it <name_container> bash"
//Para entrar a la base de datos se utiliza el comando "psql -U <user>"
//\l para ver las bases de datos
//\c <data base Name> para entrar a la base de datos


//--- VARIABLES LOCALES ---

//DATA BASE STREAM
var DSN = "host=localhost user=jorge password=mysecretpassword dbname=gorm port=5432"

var DB *gorm.DB


func DBConnection(){

    //CONEXION

    var err error

    DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("DB connected")

    }
}
