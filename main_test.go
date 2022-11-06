package main

import "testing"

func TestExampleClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestExampleClient",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleClient()
		})
	}
}
