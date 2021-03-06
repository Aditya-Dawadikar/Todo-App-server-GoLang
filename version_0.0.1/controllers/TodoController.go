package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo_service/version_0.0.1/database"
	"todo_service/version_0.0.1/models"
	"todo_service/version_0.0.1/responses"

	"github.com/gorilla/mux"
)

func GetAllTodo(w http.ResponseWriter, req *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	var todoList []models.TodoList

	sql := "select * from todo_list"
	res, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var todo models.TodoList
		err := res.Scan(&todo.OwnerId, &todo.Name, &todo.Id, &todo.Desc)

		if err != nil {
			log.Fatal(err)
		}

		sql1 := "select * from todo_list_item where tli_list_id='" + todo.Id + "'"

		res1, err1 := db.Query(sql1)

		if err1 != nil {
			log.Fatal(err1)
		}

		var itemList []models.ListItem
		for res1.Next() {
			var item models.ListItem
			err := res1.Scan(&item.ItemListId, &item.ItemId, &item.ItemName, &item.ItemStatus, &item.ItemPriority)
			if err != nil {
				log.Fatal(err)
			}

			itemList = append(itemList, item)
		}

		todo.List = itemList

		todoList = append(todoList, todo)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while fetching data"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	} else {
		succ_resp := responses.FoundTodos{Status: 200, Message: "found todos", Todos: todoList}
		json.NewEncoder(w).Encode(succ_resp)
	}
}

func CreateNewTodo(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)

	type TodoList struct {
		OwnerId string `json:"owner_id"`
		Name    string `json:"name"`
		Desc    string `json:"description"`
	}

	var todo TodoList

	if err := json.Unmarshal(req_body, &todo); err != nil {
		panic(err)
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	unique_id := timeStamp + "_todo_list"

	db := database.GetConnection()

	defer db.Close()

	sql := "insert into todo_list values ('" + todo.OwnerId + "','" + todo.Name + "','" + unique_id + "','" + todo.Desc + "')"

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 500, Message: "some error occured while executing insert"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	} else {
		succ_resp := responses.GeneralSuccess{Status: 200, Message: "todo added successfully"}
		json.NewEncoder(w).Encode(succ_resp)
	}
}

func GetAllTodoByUserId(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["user_id"]

	db := database.GetConnection()

	defer db.Close()

	var todoList []models.TodoList

	sql := "select * from todo_list where tl_owner_id='" + id + "'"

	// fmt.Println(sql)

	res, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var todo models.TodoList
		err := res.Scan(&todo.OwnerId, &todo.Name, &todo.Id, &todo.Desc)

		if err != nil {
			log.Fatal(err)
		}

		sql1 := "select * from todo_list_item where tli_list_id='" + todo.Id + "'"

		res1, err1 := db.Query(sql1)

		if err1 != nil {
			log.Fatal(err1)
		}

		var itemList []models.ListItem
		for res1.Next() {
			var item models.ListItem
			err := res1.Scan(&item.ItemListId, &item.ItemId, &item.ItemName, &item.ItemStatus, &item.ItemPriority)
			if err != nil {
				log.Fatal(err)
			}

			itemList = append(itemList, item)
		}

		todo.List = itemList

		todoList = append(todoList, todo)
	}
	// fmt.Println(todoList)

	if err != nil {
		error_resp := responses.UnknownError{Status: 500, Message: "some error occured while fetching data"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	} else {
		succ_resp := responses.FoundTodos{Status: 200, Message: "found todos", Todos: todoList}
		json.NewEncoder(w).Encode(succ_resp)
	}
}

func DeleteTodoById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	db := database.GetConnection()

	defer db.Close()

	sql := "delete from todo_list where tl_id='" + id + "'"

	// fmt.Println(sql)

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 500, Message: "some error occured while executing delete"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	}
	succ_resp := responses.GeneralSuccess{Status: 200, Message: "todo removed successfully"}
	json.NewEncoder(w).Encode(succ_resp)
}

func GetTodoById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	db := database.GetConnection()

	defer db.Close()

	var todoList []models.TodoList

	sql := "select * from todo_list where tl_id='" + id + "'"

	res, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var todo models.TodoList
		err := res.Scan(&todo.OwnerId, &todo.Name, &todo.Id, &todo.Desc)

		if err != nil {
			log.Fatal(err)
		}

		sql1 := "select * from todo_list_item where tli_list_id='" + todo.Id + "'"

		res1, err1 := db.Query(sql1)

		if err1 != nil {
			log.Fatal(err1)
		}

		var itemList []models.ListItem
		for res1.Next() {
			var item models.ListItem
			err := res1.Scan(&item.ItemListId, &item.ItemId, &item.ItemName, &item.ItemStatus, &item.ItemPriority)
			if err != nil {
				log.Fatal(err)
			}

			itemList = append(itemList, item)
		}

		todo.List = itemList

		todoList = append(todoList, todo)
	}

	requiredTodo := todoList[0]

	// fmt.Println(todoList)
	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while fetching data"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	} else {
		succ_resp := responses.FoundTodo{Status: 200, Message: "found todos", Todo: requiredTodo}
		json.NewEncoder(w).Encode(succ_resp)
	}
}
