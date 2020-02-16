package logic

import (
	"github.com/kelseyhightower/confd/log"
	"testing"
)

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
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, user := range gotUsers{
				log.Info("user: %v", user)
			}
		})
	}
}