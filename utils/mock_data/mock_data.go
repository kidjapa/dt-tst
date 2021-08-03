package mock_data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

/**
Simple mock data from mock json file with helper function to avoid unnecessarily
code into stores and handlers
*/

var (
	ErrorNotFounded = errors.New("record not found")
)

type MockData struct {
	Agents []Agent `json:"agents"`
}

type MockDataHandler struct {
	Data    *MockData
	Options *MockDataConfig
	Mu      sync.RWMutex
}

func NewMockData(opts ...MockDataOption) (mdh *MockDataHandler, err error) {
	mdh = &MockDataHandler{}
	mdh.Options = mdh.newChargeValuePaymentSettings(opts)

	// Apply default path if mdh.Options.FilePath are empty
	if mdh.Options.FilePath == "" {
		mdh.Options.FilePath = "mock_data.json"
	}

	mdh.updateData()
	return
}

func (m *MockDataHandler) Reset() (err error) {
	err = m.updateData()
	return
}

func (m *MockDataHandler) TransferInAgent(agent *Agent, destAccountId string, amount float64) (resAgent *Agent, err error) {
	m.Mu.Lock()
	resAgent = &Agent{}
	found := false
	indexToTransfer := -1
	indexToWithdraw := -1
	for i, a := range m.Data.Agents {
		if a.AccountId == agent.AccountId {
			indexToWithdraw = i
		}
		if a.AccountId == destAccountId {
			indexToTransfer = i
			found = true
			break
		}
	}
	if found && indexToTransfer != -1 && indexToWithdraw != -1 {
		m.Data.Agents[indexToTransfer].Balance += amount
		resAgent = &m.Data.Agents[indexToTransfer]
		m.Data.Agents[indexToWithdraw].Balance -= amount
		m.Mu.Unlock()
		return
	}
	m.Mu.Unlock()
	err = ErrorNotFounded
	return
}

func (m *MockDataHandler) WithdrawFromAgent(accountId string, amount float64) (resAgent *Agent, err error) {
	m.Mu.Lock()
	resAgent = &Agent{}
	for i, a := range m.Data.Agents {
		if a.AccountId == accountId {
			m.Data.Agents[i].Balance -= amount
			resAgent = &m.Data.Agents[i]
			m.Mu.Unlock()
			return
		}
	}
	m.Mu.Unlock()
	err = ErrorNotFounded
	return
}

func (m *MockDataHandler) DepositInAgent(accountId string, amount float64) (resAgent *Agent, err error) {
	m.Mu.Lock()
	resAgent = &Agent{}
	for i, a := range m.Data.Agents {
		if a.AccountId == accountId {
			m.Data.Agents[i].Balance += amount
			resAgent = &m.Data.Agents[i]
			m.Mu.Unlock()
			return
		}
	}
	m.Mu.Unlock()
	err = ErrorNotFounded
	return
}

// GetAgentById get a specified agent by their accountId
func (m *MockDataHandler) GetAgentById(accountId string) (agent *Agent, err error) {
	m.Mu.Lock()
	agent = &Agent{}
	for i, a := range m.Data.Agents {
		if a.AccountId == accountId {
			agent = &m.Data.Agents[i]
			m.Mu.Unlock()
			return
		}
	}
	m.Mu.Unlock()
	err = ErrorNotFounded
	return
}

func (m *MockDataHandler) updateData() (err error){
	file, err := os.Open(m.Options.FilePath)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &m.Data)
	return
}
