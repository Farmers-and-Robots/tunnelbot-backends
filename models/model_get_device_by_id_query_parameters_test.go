// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Tunnelbot - OpenAPI 3.0
//	Version: 1
package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeviceByIdQueryParameters_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         GetDeviceByIdQueryParameters
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.m.Validate())
		})
	}
}

func TestGetDeviceByIdQueryParameters_GetDeviceId(t *testing.T) {
	tests := []struct {
		name string
		m    GetDeviceByIdQueryParameters
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetDeviceId())
		})
	}
}

func TestGetDeviceByIdQueryParameters_SetDeviceId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *GetDeviceByIdQueryParameters
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetDeviceId(tt.args.val)
		})
	}
}