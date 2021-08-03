package mock_data

import (
	"reflect"
	"testing"
)

func TestMockData_GetAgentById(t *testing.T) {
	mockDataHandler, err := NewMockData()
	if err != nil {
		t.Errorf("GetAgentById() error = %v", err)
		return
	}
	type args struct {
		accountId string
	}
	tests := []struct {
		name      string
		fields    *MockDataHandler
		args      args
		wantAgent *Agent
		wantErr   bool
	}{
		{
			name:   "test agent id 3808",
			fields: mockDataHandler,
			args:   args{accountId: "3808"},
			wantAgent: &Agent{
				AccountId: "3808",
				Balance:   14563.41,
			},
			wantErr: false,
		},
		{
			name:      "test not found agent",
			fields:    mockDataHandler,
			args:      args{accountId: "123123"},
			wantAgent: &Agent{},
			wantErr:   true,
		},
		{
			name:      "test non numeric string account id",
			fields:    mockDataHandler,
			args:      args{accountId: "asdfasdf"},
			wantAgent: &Agent{},
			wantErr:   true,
		},
		{
			name:      "test empty account id",
			fields:    mockDataHandler,
			args:      args{accountId: ""},
			wantAgent: &Agent{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDataHandler{
				Data: &MockData{Agents: tt.fields.Data.Agents},
			}
			gotAgent, err := m.GetAgentById(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAgentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAgent, tt.wantAgent) {
				t.Errorf("GetAgentById() gotAgent = %v, want %v", gotAgent, tt.wantAgent)
			}
		})
	}
}

func TestMockData_TransferInAgent(t *testing.T) {
	mockDataHandler, err := NewMockData()
	if err != nil {
		t.Errorf("TransferInAgent() error = %v", err)
		return
	}
	type args struct {
		agent         *Agent
		destAccountId string
		amount        float64
	}
	tests := []struct {
		name         string
		fields       *MockDataHandler
		args         args
		wantResAgent *Agent
		wantErr      bool
	}{
		{
			name:   "Test transfer 3808 to 2695",
			fields: mockDataHandler,
			args: args{
				agent: &Agent{
					AccountId: "3808",
					Balance:   14563.41,
				},
				destAccountId: "2695",
				amount:        150,
			},
			wantResAgent: &Agent{
				AccountId: "2695",
				Balance:   5902.64,
			},
			wantErr:      false,
		},
		{
			name:   "Test transfer 3808 to 1648",
			fields: mockDataHandler,
			args: args{
				agent: &Agent{
					AccountId: "3808",
					Balance:   14413.41,
				},
				destAccountId: "2695",
				amount:        14413.41,
			},
			wantResAgent: &Agent{
				AccountId: "2695",
				Balance:   20316.05,
			},
			wantErr:      false,
		},
		{
			name:   "Test transfer 3808 to non account",
			fields: mockDataHandler,
			args: args{
				agent: &Agent{
					AccountId: "3808",
					Balance:   14413.41,
				},
				destAccountId: "132123",
				amount:        14413.41,
			},
			wantResAgent: &Agent{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDataHandler{
				Data: &MockData{Agents: tt.fields.Data.Agents},
			}
			gotResAgent, err := m.TransferInAgent(tt.args.agent, tt.args.destAccountId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferInAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResAgent, tt.wantResAgent) {
				t.Errorf("TransferInAgent() gotResAgent = %v, want %v", gotResAgent, tt.wantResAgent)
			}
		})
	}
}

func TestMockData_DepositInAgent(t *testing.T) {
	mockDataHandler, err := NewMockData()
	if err != nil {
		t.Errorf("DepositInAgent() error = %v", err)
		return
	}
	type args struct {
		accountId string
		amount    float64
	}
	tests := []struct {
		name         string
		fields       *MockDataHandler
		args         args
		wantResAgent *Agent
		wantErr      bool
	}{
		{
			name:         "Test deposit into 3808",
			fields:       mockDataHandler,
			args:         args{
				accountId: "3808",
				amount:    150,
			},
			wantResAgent: &Agent{
				AccountId: "3808",
				Balance:   14713.41,
			},
			wantErr:      false,
		},
		{
			name:         "Test deposit into 1648",
			fields:       mockDataHandler,
			args:         args{
				accountId: "1648",
				amount:    150,
			},
			wantResAgent: &Agent{
				AccountId: "1648",
				Balance:   10270.07,
			},
			wantErr:      false,
		},
		{
			name:         "Test deposit an non account",
			fields:       mockDataHandler,
			args:         args{
				accountId: "",
				amount:    150,
			},
			wantResAgent: &Agent{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDataHandler{
				Data: &MockData{Agents: tt.fields.Data.Agents},
			}
			gotResAgent, err := m.DepositInAgent(tt.args.accountId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositInAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResAgent, tt.wantResAgent) {
				t.Errorf("DepositInAgent() gotResAgent = %v, want %v", gotResAgent, tt.wantResAgent)
			}
		})
	}
}