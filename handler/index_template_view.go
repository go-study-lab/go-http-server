package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type Teacher struct {
	Name    string
	Subject string
}
type Student struct {
	Id      int
	Name    string
	Country string
}

type Rooster struct {
	Teacher Teacher
	Students []Student
}
func ShowIndexView(response http.ResponseWriter, request *http.Request) {

	teacher := Teacher{
		Name:    "Alex",
		Subject: "Physics",
	}
	students := []Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	rooster := Rooster{
		Teacher:  teacher,
		Students: students,
	}

	tmpl, err := template.ParseFiles("./views/layout.gohtml", "./views/nav.gohtml", "./views/content.gohtml")
	if err != nil {
		fmt.Println("Error " +  err.Error())
	}
	tmpl.Execute(response, rooster)
}
