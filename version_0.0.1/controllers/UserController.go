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

type AuthRequestBody struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)

	var auth_user AuthRequestBody
	if err := json.Unmarshal(req_body, &auth_user); err != nil {
		panic(err)
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	unique_id := timeStamp + "_" + auth_user.Username

	sql := "insert into todo_user values('" + auth_user.Username + "','" + auth_user.Password + "','" + unique_id + "')"

	db := database.GetConnection()
	defer db.Close()

	_, err := db.Exec(sql)

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "some error occured while executing insert"}
		json.NewEncoder(w).Encode(error_resp)
		panic(err.Error())
	}
	succ_resp := responses.LoginSuccess{Status: 200, Message: "User registered successfully", Username: auth_user.Username, Userid: unique_id}
	json.NewEncoder(w).Encode(succ_resp)

}

func LoginUser(w http.ResponseWriter, req *http.Request) {

	req_body, _ := ioutil.ReadAll(req.Body)

	var auth_user AuthRequestBody
	if err := json.Unmarshal(req_body, &auth_user); err != nil {
		panic(err)
	}

	sql := "select tu_userid from todo_user where tu_username='" + auth_user.Username + "' and tu_password='" + auth_user.Password + "';"
	db := database.GetConnection()
	defer db.Close()

	res, err := db.Query(sql)

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var userid string
	for res.Next() {
		err := res.Scan(&userid)

		if err != nil {
			log.Fatal(err)
		}
	}

	if userid == "" {
		error_resp := responses.LoginError{Status: 404, Message: "User not found, register new user"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
	} else {
		succ_resp := responses.LoginSuccess{Status: 200, Message: "User found", Username: auth_user.Username, Userid: userid}
		json.NewEncoder(w).Encode(succ_resp)
	}
}

func GetUserById(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["id"]

	db := database.GetConnection()
	defer db.Close()

	res, err := db.Query("select * from todo_user where tu_userid='" + id + "'")

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "users not found"}
		json.NewEncoder(w).Encode(error_resp)
		log.Fatal(err)
	}

	var userList []models.TodoUser

	for res.Next() {
		var user models.TodoUser
		err := res.Scan(&user.UserName, &user.UserPassword, &user.UserId)

		if err != nil {
			log.Fatal(err)
		}

		userList = append(userList, user)
	}

	if userList != nil {
		succ_resp := responses.FoundUser{Status: 200, Message: "User found", User: userList[0]}
		json.NewEncoder(w).Encode(succ_resp)
	} else {
		error_resp := responses.LoginError{Status: 404, Message: "User not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
	}

}

func GetAllUsers(w http.ResponseWriter, req *http.Request) {

	db := database.GetConnection()

	res, err := db.Query("select * from todo_user")

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "users not found"}
		json.NewEncoder(w).Encode(error_resp)
		log.Fatal(err)
	}

	var userList []models.TodoUser

	for res.Next() {
		var user models.TodoUser
		err := res.Scan(&user.UserName, &user.UserPassword, &user.UserId)

		if err != nil {
			log.Fatal(err)
		}

		userList = append(userList, user)
	}

	if userList != nil {
		succ_resp := responses.FoundUsers{Status: 200, Message: "User found", Users: userList}
		json.NewEncoder(w).Encode(succ_resp)
	} else {
		error_resp := responses.LoginError{Status: 404, Message: "No users found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
	}
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {

}
