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

	// stmt, err := db.Prepare("INSERT INTO posts(title, body) values(?,?)")
	// checkErr(err)

	// _, err = stmt.Exec("My First Post", "My First content")
	// checkErr(err)

	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}

	db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", items); err != nil {
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
