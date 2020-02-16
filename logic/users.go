package logic

import (
	"example.com/http_demo/model/dao"
	"example.com/http_demo/model/dao/table"
)

func GetAllUsers() (users []*table.User, err error) {
	users, err = dao.GetAllUsers()

	return
}
