package config

import (
	"os"
	"testing"
)

func TestGetPortEnv(t *testing.T) {
	_ = os.Setenv("ENV_PORT", "8080")

	tests := []struct {
		name string
		want int64
	}{
		{name: "teste 8080", want: 8080},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPortEnv(); got != tt.want {
				t.Errorf("GetPortEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
