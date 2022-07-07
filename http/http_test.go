package http

import "testing"

func TestStartServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenGcacheOrLoad(DefaultGroupName)
			StartServer()
		})
	}
}

func TestStart(t *testing.T) {
	GenGcacheOrLoad(DefaultGroupName)
	StartServer()
}
