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

func TestRole_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         Role
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