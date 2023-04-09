package models

import "gorm.io/gorm"

type User struct {
    //Para que este struct se transforme en una tabla de postgres toca utilizar 
    gorm.Model


    FirstName string `gorm:"not null" json:"first_name"`
    LastName string `gorm:"not null" json:"last_name"`
    Email string `gorm:"not null;unique_index" json:"email"`
    Tasks []Task  `json:"tasks"`
}
