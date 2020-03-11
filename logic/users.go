package logic

import (
	"example.com/http_demo/model/dao"
	"example.com/http_demo/model/dao/table"
)

func GetAllUsers() (users []*table.User, err error) {
	users, err = dao.GetAllUsers()

	return
}

func AuthenticateUser(name, password string) (user *table.User, err error) {
	user, err = dao.GetUserByNameAndPassword(name, password)
	return
}