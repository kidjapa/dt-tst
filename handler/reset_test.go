package handler

import (
	"dt-tst/api/balance"
	"dt-tst/api/event"
	"dt-tst/api/reset"
	"dt-tst/router"
	"dt-tst/utils/mock_data"
	"net/url"
	"reflect"
	"testing"
)

func TestPostReset(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	e := router.New()
	h := &Handler{
		BalanceStore: balance.NewHandler(mockDataHandler),
		EventStore:   event.NewHandler(mockDataHandler),
		Reset: reset.NewHandler(mockDataHandler),
	}

	tests := []struct {
		name               string
		wantRes            string
	}{
		{
			name:    "Test reset",
			wantRes: "OK",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(url.Values)
			c, rec := getContextGetReq(q, e)
			gotRes := h.postReset(c)
			if gotRes != nil {
				t.Errorf("GetPortEnv() error: = %v", gotRes)
				return
			}
			if !reflect.DeepEqual(rec.Body.String(), tt.wantRes) {
				t.Errorf("PostWithdraw() gotRes = %s, want: %v", rec.Body.String(), tt.wantRes)
			}
		})
	}
}