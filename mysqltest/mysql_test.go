package mysqltest

import "testing"

func TestDbPing(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DbPing()
		})
	}
}

func TestMysqlDemoCode(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MysqlDemoCode()
		})
	}
}