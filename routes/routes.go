package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-htmx-web/model"
	"html/template"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {

    todos, err := model.GetAllTodos()
    if err != nil {
        fmt.Println("Could not get all todos from db", err)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    err = tmpl.ExecuteTemplate(w, "Todos", todos)
    if err != nil {
        fmt.Println("Could not execute template", err)
    }
}


func index(w http.ResponseWriter, r *http.Request) {
    
    todos, err := model.GetAllTodos()
    if err != nil {
        fmt.Println("Could not get all todos from db", err)
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    err = tmpl.Execute(w, todos)
    if err != nil {
        fmt.Println("Could not execute template", err)
    }

}


func markTodo(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
    if err != nil {
        fmt.Println("Could not parse id", err)
    }

    err = model.MarkDone(id)
    if err != nil {
        fmt.Println("Could not update todo", err)
    }



    sendTodos(w)



}

func createTodo(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm()
    if err != nil {
        fmt.Println("Error parsing form", err)
    }

    err = model.CreateTodo(r.FormValue("todo"))
    if err != nil {
        fmt.Println("Could not create todo", err)
    }

    sendTodos(w)

}


func deleteTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
    if err != nil {
        fmt.Println("Could not parse id", err)
    }

    err = model.Delete(id)
    if err != nil {
        fmt.Println("Could not delete", err)
    }

    sendTodos(w)

}


func SetupAndRun() {
    fmt.Println("Create NewRouter")
    mux := mux.NewRouter()
    fmt.Println("Created NewRouter!")

    mux.HandleFunc("/", index)
    mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
    mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
    mux.HandleFunc("/create", createTodo).Methods("POST")
    fmt.Println("It's Ready! on http://localhost:6000")
    log.Println("It's Ready! on http://localhost:6000")
    log.Fatal(http.ListenAndServe("localhost:6000", mux))
    
}












