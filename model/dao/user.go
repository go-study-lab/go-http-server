package dao

import "example.com/http_demo/model/dao/table"

func CreateUser(user *table.User) (err error) {
	err = DB().Create(user).Error

	return
}

func GetUserById(userId int64) (user *table.User, err error) {
	user = new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error

	return
}

func GetAllUsers() (users []*table.User, err error) {
	err = DB().Find(&users).Error
	return
}

func GetUserByNameAndPassword (name, password string) (user *table.User, err error) {
	user = new(table.User)
	err = DB().Where("username = ? AND secret = ?", name, password).
		  First(&user).Error

	return
}

func UpdateUserNameById(userName string, userId int64) (err error) {
	user := new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error
	if err != nil {
		return
	}

	user.UserName = userName
	err = DB().Save(user).Error

	return
}

func DeleteUserById(userId int64) (err error) {
	user := new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error
	if err != nil {
		return
	}
	err = DB().Delete(user).Error

	return
}
