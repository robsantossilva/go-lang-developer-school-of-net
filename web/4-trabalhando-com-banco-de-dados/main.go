package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func main() {

	stmt, err := db.Prepare("INSERT INTO posts(title, body) values(?,?)")
	checkErr(err)

	_, err = stmt.Exec("My First Post", "My First content")
	checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		post := Post{Id: 1, Title: "Unnamed Post", Body: "No content"}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		if body := r.FormValue("body"); body != "" {
			post.Body = body
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8081", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
