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
	"todo_service/version_0.0.1/models"
	"todo_service/version_0.0.1/responses"

	"github.com/gorilla/mux"
)

type AuthRequestBody struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func checkUserExists(username string, password string) bool {

	sql := "select tu_userid from todo_user where tu_username='" + username + "' and tu_password='" + password + "'"

	db := database.GetConnection()
	defer db.Close()

	row := db.QueryRow(sql)

	var user string

	err := row.Scan(&user)

	if err != nil {
		return false
	} else {
		fmt.Println(user)
		return true
	}

}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)

	var auth_user AuthRequestBody
	if err := json.Unmarshal(req_body, &auth_user); err != nil {
		panic(err)
	}

	if auth_user.Username == "" || auth_user.Password == "" {
		error_resp := responses.UnknownError{Status: 403, Message: "user credentials are missing"}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(error_resp)
		return
	}

	if checkUserExists(auth_user.Username, auth_user.Password) {
		error_resp := responses.UnknownError{Status: 400, Message: "This username and password already exists"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_resp)
		return
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	unique_id := timeStamp + "_" + auth_user.Username

	sql := "insert into todo_user values('" + auth_user.Username + "','" + auth_user.Password + "','" + unique_id + "')"

	db := database.GetConnection()
	defer db.Close()

	_, err := db.Exec(sql)

	if err != nil {
		error_resp := responses.UnknownError{Status: 500, Message: "some error occured while executing insert"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_resp)
		return
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

	row := db.QueryRow(sql)

	var userid string
	err := row.Scan(&userid)

	if err != nil || userid == "" {
		error_resp := responses.LoginError{Status: 404, Message: "User not found, please register new user"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
		return
	}

	succ_resp := responses.LoginSuccess{Status: 200, Message: "User found", Username: auth_user.Username, Userid: userid}
	json.NewEncoder(w).Encode(succ_resp)

}

func GetUserById(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["id"]

	db := database.GetConnection()
	defer db.Close()

	res := db.QueryRow("select * from todo_user where tu_userid='" + id + "'")

	var user models.TodoUser
	err := res.Scan(&user.UserName, &user.UserPassword, &user.UserId)

	if err != nil {
		error_resp := responses.UnknownError{Status: 404, Message: "users not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
		log.Fatal(err)
		return
	}

	succ_resp := responses.FoundUser{Status: 200, Message: "User found", User: user}
	json.NewEncoder(w).Encode(succ_resp)

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

	if len(userList) == 0 {
		error_resp := responses.LoginError{Status: 404, Message: "No users found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_resp)
		return
	}

	succ_resp := responses.FoundUsers{Status: 200, Message: "User found", Users: userList}
	json.NewEncoder(w).Encode(succ_resp)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {

}
