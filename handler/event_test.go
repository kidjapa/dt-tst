package handler

import (
	"dt-tst/api/balance"
	"dt-tst/api/event"
	"dt-tst/router"
	"dt-tst/utils/mock_data"
	"dt-tst/utils/types"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	postEventDepositRequest = &types.EventRequest{
		Type:        types.DepositRequestTypeDeposit,
		Destination: "100",
		Amount:      150,
	}
	postEventDepositResponse = &types.DepositResponse{
		Destination: types.EventCommonResponse{
			Id:      "100",
			Balance: 150,
		},
	}
)

func TestPostEvent_Deposit(t *testing.T) {
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
		depositRequest *types.EventRequest
	}
	tests := []struct {
		name               string
		mockData           *mock_data.MockDataHandler
		h                  *Handler
		args               args
		wantRes            *types.DepositResponse
		wantResError       string
		wantResForcedError bool
		wantErr            bool
	}{
		{
			name:     "Test post event: deposit into account_id 100",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				depositRequest: postEventDepositRequest,
			},
			wantRes:            postEventDepositResponse,
			wantErr:            false,
			wantResForcedError: false,
		},
		{
			name:     "Test post event: deposit into non existing account_id",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeDeposit,
					Destination: "321321",
					Amount:      150,
				},
			},
			wantResError:       "0",
			wantResForcedError: true,
			wantErr:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.h
			c, rec := getContextJsonReq(http.MethodPost, "/event", strings.NewReader(getJsonStringfy(tt.args.depositRequest)), e)
			gotRes := h.postEvent(c)
			if (gotRes != nil) != tt.wantErr {
				t.Errorf("PostWithdraw() error = %v, wantErr %v", gotRes, tt.wantErr)
				return
			}
			if tt.wantResForcedError {
				if rec.Code != http.StatusNotFound {
					t.Errorf("PostWithdraw() http.Code = %v, want: 404", rec.Code)
					return
				}
				if rec.Body.String() != tt.wantResError {
					t.Errorf("PostWithdraw() rec.Body.String() = %s, want: %s", rec.Body.String(), tt.wantResError)
					return
				}
			} else {
				w := &types.DepositResponse{}
				err = json.Unmarshal(rec.Body.Bytes(), w)
				if (err != nil) != tt.wantErr {
					t.Errorf("PostWithdraw() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(w, tt.wantRes) {
					t.Errorf("postEvent() gotRes = %v, want %v", gotRes, tt.wantRes)
				}
			}
		})
	}
}

func TestPostEvent_Withdraw(t *testing.T) {
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
		withdrawRequest *types.EventRequest
	}
	tests := []struct {
		name               string
		mockData           *mock_data.MockDataHandler
		h                  *Handler
		args               args
		wantRes            *types.WithdrawResponse
		wantResError       string
		wantResForcedError bool
		wantErr            bool
	}{
		{
			name:     "Test post event: withdraw from account_id 100",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				withdrawRequest: &types.EventRequest{
					Type:   types.DepositRequestTypeWithDraw,
					Origin: "100",
					Amount: 100,
				},
			},
			wantRes: &types.WithdrawResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -100,
				},
			},
			wantErr:            false,
			wantResForcedError: false,
		},
		{
			name:     "Test post event: withdraw from account_id 3808",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				withdrawRequest: &types.EventRequest{
					Type:   types.DepositRequestTypeWithDraw,
					Origin: "3808",
					Amount: 100,
				},
			},
			wantRes: &types.WithdrawResponse{
				Origin: types.EventCommonResponse{
					Id:      "3808",
					Balance: 14463.41,
				},
			},
			wantErr:            false,
			wantResForcedError: false,
		},
		{
			name:     "Test post event: withdraw from non existing account_id",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				withdrawRequest: &types.EventRequest{
					Type:   types.DepositRequestTypeWithDraw,
					Origin: "321321",
					Amount: 100,
				},
			},
			wantRes:            &types.WithdrawResponse{},
			wantResError:       "0",
			wantResForcedError: true,
			wantErr:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.h
			c, rec := getContextJsonReq(http.MethodPost, "/event", strings.NewReader(getJsonStringfy(tt.args.withdrawRequest)), e)
			gotRes := h.postEvent(c)
			if (gotRes != nil) != tt.wantErr {
				t.Errorf("PostWithdraw() error = %v, wantErr %v", gotRes, tt.wantErr)
				return
			}
			if tt.wantResForcedError {
				if rec.Code != http.StatusNotFound {
					t.Errorf("PostWithdraw() http.Code = %v, want: 404", rec.Code)
					return
				}
				if rec.Body.String() != tt.wantResError {
					t.Errorf("PostWithdraw() rec.Body.String() = %s, want: %s", rec.Body.String(), tt.wantResError)
					return
				}
			} else {
				w := &types.WithdrawResponse{}
				err = json.Unmarshal(rec.Body.Bytes(), w)
				if err != nil {
					t.Errorf("PostWithdraw() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(w, tt.wantRes) {
					t.Errorf("PostWithdraw() gotRes = %v, want %v", w, tt.wantRes)
				}
			}
		})
	}
}

func TestPostEvent_Transfer(t *testing.T) {
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
		transferRequest *types.EventRequest
	}
	tests := []struct {
		name               string
		mockData           *mock_data.MockDataHandler
		h                  *Handler
		args               args
		wantRes            *types.TransferResponse
		wantResError       string
		wantResForcedError bool
		wantErr            bool
	}{
		{
			name:     "Test post event: transfer from account_id 100",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				transferRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "100",
					Destination: "3808",
					Amount:      100,
				},
			},
			wantRes: &types.TransferResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -100,
				},
				Destination: types.EventCommonResponse{
					Id:      "3808",
					Balance: 14663.41,
				},
			},
			wantErr:            false,
			wantResForcedError: false,
		},
		{
			name:     "Test post event: transfer from account_id 3808",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				transferRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "100",
					Destination: "2695",
					Amount:      100,
				},
			},
			wantRes: &types.TransferResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -200,
				},
				Destination: types.EventCommonResponse{
					Id:      "2695",
					Balance: 5852.64,
				},
			},
			wantErr:            false,
			wantResForcedError: false,
		},
		{
			name:     "Test post event: transfer from non existing account_id",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				transferRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "321321",
					Destination: "100",
					Amount:      1500,
				},
			},
			wantRes:            &types.TransferResponse{},
			wantResError:       "0",
			wantResForcedError: true,
			wantErr:            false,
		},
		{
			name:     "Test post event: verify balance transfer for account_id 100 after transfer from a non existing account",
			mockData: mockDataHandler,
			h:        h,
			args: args{
				transferRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "100",
					Destination: "2695",
					Amount:      100,
				},
			},
			wantRes: &types.TransferResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -300,
				},
				Destination: types.EventCommonResponse{
					Id:      "2695",
					Balance: 5952.64,
				},
			},
			wantErr:            false,
			wantResForcedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.h
			c, rec := getContextJsonReq(http.MethodPost, "/event", strings.NewReader(getJsonStringfy(tt.args.transferRequest)), e)
			gotRes := h.postEvent(c)
			if (gotRes != nil) != tt.wantErr {
				t.Errorf("PostWithdraw() error = %v, wantErr %v", gotRes, tt.wantErr)
				return
			}
			if tt.wantResForcedError {
				if rec.Code != http.StatusNotFound {
					t.Errorf("PostWithdraw() http.Code = %v, want: 404", rec.Code)
					return
				}
				if rec.Body.String() != tt.wantResError {
					t.Errorf("PostWithdraw() rec.Body.String() = %s, want: %s", rec.Body.String(), tt.wantResError)
					return
				}
			} else {
				w := &types.TransferResponse{}
				err = json.Unmarshal(rec.Body.Bytes(), w)
				if err != nil {
					t.Errorf("PostWithdraw() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(w, tt.wantRes) {
					t.Errorf("PostWithdraw() gotRes = %v, want %v", w, tt.wantRes)
				}
			}
		})
	}
}

/**
Helper functions
*/

func getJsonStringfy(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}

func getContextJsonReq(method, target string, body io.Reader, e *echo.Echo) (c echo.Context, rec *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	return
}
