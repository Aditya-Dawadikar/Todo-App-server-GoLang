package main

import (
	"log"
	"net/http"
	"todo_service/version_0.0.1/controllers"

	"github.com/gorilla/mux"
)

func requestHandler() {

	Router := mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/todo", controllers.CreateNewTodo).Methods("POST")
	Router.HandleFunc("/todo/all", controllers.GetAllTodo).Methods("GET")
	Router.HandleFunc("/todo/all/{user_id}", controllers.GetAllTodoByUserId).Methods("GET")

	Router.HandleFunc("/todo/{id}", controllers.GetTodoById).Methods("GET")
	Router.HandleFunc("/todo/{id}", controllers.AddItemByTodoId).Methods("PATCH")
	Router.HandleFunc("/todo/{id}", controllers.DeleteTodoById).Methods("DELETE")
	Router.HandleFunc("/todo/item/{item_id}", controllers.RemoveItemById).Methods("DELETE")
	Router.HandleFunc("/todo/item/{item_id}", controllers.MarkItemById).Methods("PATCH")

	Router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	Router.HandleFunc("/user/signup", controllers.RegisterUser).Methods("POST")
	Router.HandleFunc("/user/delete", controllers.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/user/all", controllers.GetAllUsers).Methods("GET")
	Router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", Router))
}

func main() {
	requestHandler()
}
