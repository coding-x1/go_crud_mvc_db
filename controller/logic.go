package controller

import (
	"fmt"
	"log"
	"net/http"
	"simpleapp/model"
	"strconv"
	"text/template"
)

func Master() {
	//You can think of handlers as being a bit like controllers. Generally speaking, they're responsible for carrying out your application logic and writing response headers and bodies.
	http.HandleFunc("/index", index())
	http.HandleFunc("/create", create())
	http.HandleFunc("/delete", delete())
	http.HandleFunc("/update", update())
}
func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create Request Received")
		val := r.FormValue("val")
		model.Create(val)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete Request Received")
		id1, _ := strconv.Atoi(r.FormValue("id1"))
		model.Delete(id1)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
func update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update Request Received")
		val := r.FormValue("val")
		id1, _ := strconv.Atoi(r.FormValue("id1"))
		model.Update(val, id1)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Index Request Received")
		var id int
		var en string
		type d struct {
			i []int
			e []string
		}
		p := make(map[int]string)
		rows := model.Index()
		for rows.Next() {
			rows.Scan(&id, &en)
			//fmt.Println(id, en)
			p[id] = en

		}
		fmt.Println(p)

		tmp, err := template.ParseGlob("./view/*.gohtml")
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}

		tmp.ExecuteTemplate(w, "index.gohtml", p)

	}
}
