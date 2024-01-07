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

func TestAddPersonQueryParameters_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         AddPersonQueryParameters
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

func TestAddPersonQueryParameters_GetFarmId(t *testing.T) {
	tests := []struct {
		name string
		m    AddPersonQueryParameters
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetFarmId())
		})
	}
}

func TestAddPersonQueryParameters_SetFarmId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *AddPersonQueryParameters
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetFarmId(tt.args.val)
		})
	}
}