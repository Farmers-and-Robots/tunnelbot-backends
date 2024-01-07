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

func TestFarmAssociation_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         FarmAssociation
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

func TestFarmAssociation_GetFarmId(t *testing.T) {
	tests := []struct {
		name string
		m    FarmAssociation
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

func TestFarmAssociation_SetFarmId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *FarmAssociation
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

func TestFarmAssociation_GetFarmName(t *testing.T) {
	tests := []struct {
		name string
		m    FarmAssociation
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetFarmName())
		})
	}
}

func TestFarmAssociation_SetFarmName(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *FarmAssociation
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetFarmName(tt.args.val)
		})
	}
}

func TestFarmAssociation_GetRole(t *testing.T) {
	tests := []struct {
		name string
		m    FarmAssociation
		want Role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetRole())
		})
	}
}

func TestFarmAssociation_SetRole(t *testing.T) {
	type args struct {
		val Role
	}
	tests := []struct {
		name string
		m    *FarmAssociation
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetRole(tt.args.val)
		})
	}
}
