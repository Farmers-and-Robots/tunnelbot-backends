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

func TestAddress_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         Address
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "valid address",
			m: ValidAddress,
			assertion: assert.NoError,
		},
		{
			name: "empty address",
			m: Address{
				City:   "",
				State:  "",
				Street: "",
				Zip:    "",
			},
			assertion: assert.NoError,
		},
		{
			name: "nil address",
			m: Address{},
			assertion: assert.NoError,
		},
		}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.m.Validate())
		})
	}
}

func TestAddress_GetCity(t *testing.T) {
	tests := []struct {
		name string
		m    Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetCity())
		})
	}
}

func TestAddress_SetCity(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Address
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetCity(tt.args.val)
		})
	}
}

func TestAddress_GetState(t *testing.T) {
	tests := []struct {
		name string
		m    Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetState())
		})
	}
}

func TestAddress_SetState(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Address
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetState(tt.args.val)
		})
	}
}

func TestAddress_GetStreet(t *testing.T) {
	tests := []struct {
		name string
		m    Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetStreet())
		})
	}
}

func TestAddress_SetStreet(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Address
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetStreet(tt.args.val)
		})
	}
}

func TestAddress_GetZip(t *testing.T) {
	tests := []struct {
		name string
		m    Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetZip())
		})
	}
}

func TestAddress_SetZip(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Address
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetZip(tt.args.val)
		})
	}
}
