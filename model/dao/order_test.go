package dao

import (
	"example.com/http_demo/model/dao/table"
	"example.com/http_demo/utils/zlog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestGetAllOrders(t *testing.T) {
	tests := []struct {
		name       string
		wantOrders []*table.Order
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:    "test1",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrders, err := GetAllOrders()
			//if !tt.wantErr(t, err, fmt.Sprintf("GetAllOrders()")) {
			//	return
			//}
			zlog.Info("orders data log", zap.Any("data", gotOrders), zap.Any("err", err))
		})
	}
}
