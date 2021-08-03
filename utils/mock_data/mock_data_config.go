package mock_data

type MockDataConfig struct {
	FilePath string // store the path string for the json file
}

func (m *MockDataHandler) newChargeValuePaymentSettings(opts []MockDataOption) *MockDataConfig {
	var ao MockDataConfig
	for _, opt := range opts {
		opt.Apply(&ao)
	}
	return &ao
}

// MockDataOption interface for GetChargeValuesPaymentConfig
type MockDataOption interface {
	Apply(*MockDataConfig)
}

type filePath string
func (fp filePath) Apply(g *MockDataConfig) {
	g.FilePath = string(fp)
}

func WithFilePathName(path string) MockDataOption {
	return filePath(path)
}