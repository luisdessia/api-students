package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" //gorm sqlite
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model        //struct do gorm
	Name       string `json: "name"`
	CPF        int    `json: "cpf"`
	Email      string `json: "email"`
	Age        int    `json: "age"`
	Active     bool   `json: "registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {

		//log fatal para parar a aplicação em caso de erro
		log.Fatalln(err)
	}

	//gerencia a struct Student
	db.AutoMigrate(&Student{})

	return db
}

//testar o banco de dados de froma manual
/*func AddStudent(){
db := Init()

student := Student{
	Name: "Bento",
	CPF: 12345,
	Email: "bento@gmail.com",
	Age: 20,
	Active: true,
}*/

func AddStudent(student Student) {
	db := Init()

	result := db.Create(&student)
	if result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student!")
}
