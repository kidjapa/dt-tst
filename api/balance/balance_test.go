package balance

import (
	"dt-tst/utils/mock_data"
	"testing"
)

func TestHandler_GetBalance(t *testing.T) {
	mockDataHandler, err := mock_data.NewMockData(mock_data.WithFilePathName("../../utils/mock_data/mock_data.json"))
	if err != nil {
		t.Errorf("PostDeposit() error = %v", err)
		return
	}
	type args struct {
		accountId string
	}
	tests := []struct {
		name        string
		mockData    *mock_data.MockDataHandler
		args        args
		wantBalance float64
		wantErr     bool
	}{
		{
			name:        "Get balance from account_id 100",
			mockData:    mockDataHandler,
			args:        args{
				accountId: "100",
			},
			wantBalance: 0,
			wantErr:     false,
		},
		{
			name:        "Get balance from non existing account_id",
			mockData:    mockDataHandler,
			args:        args{
				accountId: "123123",
			},
			wantBalance: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				MockDataHandler: tt.mockData,
			}
			gotBalance, err := h.GetBalance(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBalance != tt.wantBalance {
				t.Errorf("GetBalance() gotBalance = %v, want %v", gotBalance, tt.wantBalance)
			}
		})
	}
}
