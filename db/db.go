package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" //gorm sqlite
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

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

func NewStudentHandler(db *gorm.DB) *StudentHandler{
	return &StudentHandler{DB : db}
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

func (s * StudentHandler) AddStudent(student Student) error{


	result := s.DB.Create(&student)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student!")
	return nil
}

func (s * StudentHandler) GetStudents() ([]Student, error){
	students := []Student {}
	
	err := s.DB.Find(&students).Error
	return students, err
}
