package models

type TodoUser struct {
	UserName     string `json:username`
	UserPassword string `json:password`
	UserId       string `json:userid`
}
