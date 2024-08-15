package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Age   int
	Meta  UserMeta
	Big   string
	Small template.HTML
}

type UserMeta struct {
	Visit int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Taniya",
		Age:  12,
		Meta: UserMeta{
			Visit: 12,
		},
		Big:   "<script>This is way using XSS </script>",
		Small: "<script>This is way using XSS if we need to avoid encoding </script>",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
