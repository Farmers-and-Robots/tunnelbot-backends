package openapi

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewAPIResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name string
		args args
		want *APIResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIResponse(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIResponseWithError(t *testing.T) {
	type args struct {
		errorMessage string
	}
	tests := []struct {
		name string
		args args
		want *APIResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIResponseWithError(tt.args.errorMessage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponseWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}
