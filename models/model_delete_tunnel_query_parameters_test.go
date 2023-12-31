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

func TestDeleteTunnelQueryParameters_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         DeleteTunnelQueryParameters
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

func TestDeleteTunnelQueryParameters_GetApiKey(t *testing.T) {
	tests := []struct {
		name string
		m    DeleteTunnelQueryParameters
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetApiKey())
		})
	}
}

func TestDeleteTunnelQueryParameters_SetApiKey(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *DeleteTunnelQueryParameters
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetApiKey(tt.args.val)
		})
	}
}

func TestDeleteTunnelQueryParameters_GetTunnelId(t *testing.T) {
	tests := []struct {
		name string
		m    DeleteTunnelQueryParameters
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetTunnelId())
		})
	}
}

func TestDeleteTunnelQueryParameters_SetTunnelId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *DeleteTunnelQueryParameters
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetTunnelId(tt.args.val)
		})
	}
}
