// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Tunnelbot - OpenAPI 3.0
//	Version: 1
package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvent_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         Event
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

func TestEvent_GetDate(t *testing.T) {
	tests := []struct {
		name string
		m    Event
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetDate())
		})
	}
}

func TestEvent_SetDate(t *testing.T) {
	type args struct {
		val time.Time
	}
	tests := []struct {
		name string
		m    *Event
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetDate(tt.args.val)
		})
	}
}

func TestEvent_GetDescription(t *testing.T) {
	tests := []struct {
		name string
		m    Event
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetDescription())
		})
	}
}

func TestEvent_SetDescription(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Event
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetDescription(tt.args.val)
		})
	}
}

func TestEvent_GetId(t *testing.T) {
	tests := []struct {
		name string
		m    Event
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetId())
		})
	}
}

func TestEvent_SetId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *Event
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetId(tt.args.val)
		})
	}
}

func TestEvent_GetKind(t *testing.T) {
	tests := []struct {
		name string
		m    Event
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetKind())
		})
	}
}

func TestEvent_SetKind(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Event
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetKind(tt.args.val)
		})
	}
}

func TestEvent_GetOperator(t *testing.T) {
	tests := []struct {
		name string
		m    Event
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetOperator())
		})
	}
}

func TestEvent_SetOperator(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Event
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetOperator(tt.args.val)
		})
	}
}
