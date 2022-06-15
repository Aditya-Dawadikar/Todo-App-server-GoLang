package main

import (
	"log"
	"net/http"
	"todo_service/version_0.0.1/auth"
	"todo_service/version_0.0.1/controllers"

	"github.com/gorilla/mux"
)

func requestHandler() {

	Router := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("./public/"))
	Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	Router.Handle("/todo", auth.IsAuthorized(controllers.CreateNewTodo)).Methods("POST")
	Router.Handle("/todo/all", auth.IsAuthorized(controllers.GetAllTodo)).Methods("GET")
	Router.Handle("/todo/all/{user_id}", auth.IsAuthorized(controllers.GetAllTodoByUserId)).Methods("GET")

	Router.Handle("/todo/{id}", auth.IsAuthorized(controllers.GetTodoById)).Methods("GET")
	Router.Handle("/todo/{id}", auth.IsAuthorized(controllers.AddItemByTodoId)).Methods("PATCH")
	Router.Handle("/todo/{id}", auth.IsAuthorized(controllers.DeleteTodoById)).Methods("DELETE")
	Router.Handle("/todo/item/{item_id}", auth.IsAuthorized(controllers.RemoveItemById)).Methods("DELETE")
	Router.Handle("/todo/item/{item_id}", auth.IsAuthorized(controllers.MarkItemById)).Methods("PATCH")

	Router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	Router.HandleFunc("/user/signup", controllers.RegisterUser).Methods("POST")
	// Router.HandleFunc("/user/delete", controllers.DeleteUser).Methods("DELETE")
	Router.Handle("/user/all", auth.IsAuthorized(controllers.GetAllUsers)).Methods("GET")
	Router.Handle("/user/{id}", auth.IsAuthorized(controllers.GetUserById)).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", Router))
}

func main() {
	requestHandler()
}
