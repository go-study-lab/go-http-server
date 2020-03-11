package dao

import (
	"example.com/http_demo/model/dao/table"
	"github.com/kelseyhightower/confd/log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user *table.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{user: &table.User{UserName: "Kevin", Secret: "123456"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
	}{
		{
			name: "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, user := range gotUsers {
				log.Info("user: %v", user)
			}

		})
	}
}

func TestUpdateUserNameById(t *testing.T) {
	type args struct {
		userName string
		userId   int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				userName: "Klein",
				userId:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserNameById(tt.args.userName, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserNameById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByNameAndPassword(t *testing.T) {
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			name: "test",
			args: args{
				name:     "Klein",
				password: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := GetUserByNameAndPassword(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByNameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Info("user %v", gotUser)
		})
	}
}