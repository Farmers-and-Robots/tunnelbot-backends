package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRouter())
		})
	}
}

func Test_optionsHandlerFunc(t *testing.T) {
	type args struct {
		allowedMethods []string
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, optionsHandlerFunc(tt.args.allowedMethods...))
		})
	}
}
