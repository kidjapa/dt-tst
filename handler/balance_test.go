package handler

import (
	"dt-tst/api/balance"
	"dt-tst/api/event"
	"dt-tst/router"
	"dt-tst/utils/mock_data"
	"dt-tst/utils/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestGetReset(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	e := router.New()
	h := &Handler{
		BalanceStore: balance.NewHandler(mockDataHandler),
		EventStore:   event.NewHandler(mockDataHandler),
	}
	type args struct {
		BalanceRequest types.BalanceRequest
	}
	tests := []struct {
		name    string
		wantRes string
		args    args
	}{
		{
			name:    "Test get \"balance\" from account_id 100",
			wantRes: "0",
			args: args{
				BalanceRequest: types.BalanceRequest{AccountId: "100"},
			},
		},
		{
			name:    "Test get \"balance\" from non existing account_id",
			wantRes: "0",
			args: args{
				BalanceRequest: types.BalanceRequest{AccountId: "321321"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(url.Values)
			q.Add("account_id", tt.args.BalanceRequest.AccountId)
			c, rec := getContextGetReq(q, e)
			gotRes := h.getBalance(c)
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

func getContextGetReq(q url.Values, e *echo.Echo) (c echo.Context, rec *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/balance?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	return
}
