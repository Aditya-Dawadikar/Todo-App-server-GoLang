package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo_service/version_0.0.1/database"
	"todo_service/version_0.0.1/responses"

	"github.com/gorilla/mux"
)

func AddItemById(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)
	params := mux.Vars(req)
	id := params["id"]

	type ItemList struct {
		ItemName     string `json:"item_name"`
		ItemStatus   string `json:"item_status"`
		ItemPriority int64  `json:"item_priority"`
	}

	var item ItemList

	if err := json.Unmarshal(req_body, &item); err != nil {
		panic(err)
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	unique_id := timeStamp + "_todo_item"

	db := database.GetConnection()

	defer db.Close()

	// fmt.Println(item, unique_id)

	sql := "insert into todo_list_item values ('" + id + "','" + unique_id + "','" + item.ItemName + "','" + item.ItemStatus + "','" + strconv.FormatInt(item.ItemPriority, 10) + "')"

	fmt.Println(sql)

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while executing insert"}
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	}
	succ_resp := responses.GeneralSuccess{Status: 200, Message: "todo item added successfully"}
	json.NewEncoder(w).Encode(succ_resp)
}

func RemoveItemById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["item_id"]

	db := database.GetConnection()

	defer db.Close()

	sql := "delete from todo_list_item where tli_id='" + id + "'"

	fmt.Println(sql)

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while executing delete"}
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	}
	succ_resp := responses.GeneralSuccess{Status: 200, Message: "todo item removed successfully"}
	json.NewEncoder(w).Encode(succ_resp)
}

func MarkItemById(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)
	params := mux.Vars(req)
	id := params["item_id"]

	type ItemStatus struct {
		Status string `json:"item_status"`
	}

	var itemStatus ItemStatus

	if err := json.Unmarshal(req_body, &itemStatus); err != nil {
		panic(err)
	}

	db := database.GetConnection()

	defer db.Close()

	fmt.Println(itemStatus)

	sql := "update todo_list_item set tli_status='" + itemStatus.Status + "' where tli_id='" + id + "'"

	fmt.Println(sql)

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while executing update"}
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	}
	succ_resp := responses.GeneralSuccess{Status: 200, Message: "todo item updated successfully"}
	json.NewEncoder(w).Encode(succ_resp)
}
