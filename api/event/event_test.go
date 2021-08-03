package event

import (
	"dt-tst/utils/mock_data"
	"dt-tst/utils/types"
	"reflect"
	"testing"
)

func TestHandler_PostDeposit(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	type args struct {
		depositRequest *types.EventRequest
	}
	tests := []struct {
		name     string
		mockData *mock_data.MockDataHandler
		args     args
		wantRes  *types.DepositResponse
		wantErr  bool
	}{
		{
			name:     "Test deposit",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeDeposit,
					Destination: "100",
					Amount:      100,
				},
			},
			wantRes: &types.DepositResponse{
				Destination: types.EventCommonResponse{
					Id:      "100",
					Balance: 100,
				},
			},
			wantErr: false,
		},
		{
			name:     "Test deposit 0",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeDeposit,
					Destination: "100",
					Amount:      0,
				},
			},
			wantRes: &types.DepositResponse{
				Destination: types.EventCommonResponse{
					Id:      "100",
					Balance: 100,
				},
			},
			wantErr: false,
		},
		{
			name:     "Test deposit -100",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeDeposit,
					Destination: "100",
					Amount:      -100,
				},
			},
			wantRes: &types.DepositResponse{
				Destination: types.EventCommonResponse{
					Id:      "100",
					Balance: 0,
				},
			},
			wantErr: false,
		},
		{
			name:     "Test deposit non existing account id",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeDeposit,
					Destination: "321321",
					Amount:      100,
				},
			},
			wantRes: &types.DepositResponse{
				Destination: types.EventCommonResponse{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				MockDataHandler: tt.mockData,
			}
			gotRes, err := h.PostDeposit(tt.args.depositRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("PostDeposit() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestHandler_PostWithdraw(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	type args struct {
		depositRequest *types.EventRequest
	}
	tests := []struct {
		name     string
		mockData *mock_data.MockDataHandler
		args     args
		wantRes  *types.WithdrawResponse
		wantErr  bool
	}{
		{
			name:     "Test withdraw into account_id 100",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
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
			wantErr: false,
		},
		{
			name:     "Test withdraw into non existing account_id",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:   types.DepositRequestTypeWithDraw,
					Origin: "321321",
					Amount: 100,
				},
			},
			wantRes: &types.WithdrawResponse{
				Origin: types.EventCommonResponse{},
			},
			wantErr: true,
		},
		{
			name:     "Test persistence into pointers",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:   types.DepositRequestTypeWithDraw,
					Origin: "100",
					Amount: 100,
				},
			},
			wantRes: &types.WithdrawResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -200,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				MockDataHandler: tt.mockData,
			}
			gotRes, err := h.PostWithdraw(tt.args.depositRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("PostWithdraw() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestHandler_PostTransfer(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	type args struct {
		depositRequest *types.EventRequest
	}
	tests := []struct {
		name     string
		mockData *mock_data.MockDataHandler
		args     args
		wantRes  *types.TransferResponse
		wantErr  bool
	}{
		{
			name:     "Test withdraw into account_id 100",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
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
			wantErr: false,
		},
		{
			name:     "Test transfer from 100 to non existing account_id",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "100",
					Destination: "654654",
					Amount:      100,
				},
			},
			wantRes: &types.TransferResponse{
				Origin:      types.EventCommonResponse{},
				Destination: types.EventCommonResponse{},
			},
			wantErr: true,
		},
		{
			name:     "Test for determine if the account_id 100 has -80 after trying to transfer for non existing account",
			mockData: mockDataHandler,
			args: args{
				depositRequest: &types.EventRequest{
					Type:        types.DepositRequestTypeTransfer,
					Origin:      "100",
					Destination: "3808",
					Amount:      100,
				},
			},
			wantRes: &types.TransferResponse{
				Origin: types.EventCommonResponse{
					Id:      "100",
					Balance: -200,
				},
				Destination: types.EventCommonResponse{
					Id:      "3808",
					Balance: 14763.41,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				MockDataHandler: tt.mockData,
			}
			gotRes, err := h.PostTransfer(tt.args.depositRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("PostTransfer() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
