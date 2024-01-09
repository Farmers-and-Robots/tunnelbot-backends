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

func TestFarm_Validate(t *testing.T) {
	tests := []struct {
		name      string
		m         Farm
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "valid farm",
			m:         ValidFarm,
			assertion: assert.NoError,
		},
		{
			name: "empty farm",
			m: Farm{},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.m.Validate())
		})
	}
}

func TestFarm_GetAddress(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want *Address
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: &ValidAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetAddress())
		})
	}
}

func TestFarm_SetAddress(t *testing.T) {
	type args struct {
		val *Address
	}
	tests := []struct {
		name string
		m    *Farm
		args args
	}{
		{
			name: "valid farm",
			m:    &ValidFarm,
			args: args{
				val: &NewValidAddress,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetAddress(tt.args.val)
			assert.Equal(t, tt.args.val, tt.m.GetAddress())
		})
	}
}

func TestFarm_GetCreateDate(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want interface{}
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: ValidCreateDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetCreateDate())
		})
	}
}

func TestFarm_SetCreateDate(t *testing.T) {
	type args struct {
		val time.Time
	}
	tests := []struct {
		name string
		m    *Farm
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetCreateDate(tt.args.val)
		})
	}
}

func TestFarm_GetId(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want int64
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetId())
		})
	}
}

func TestFarm_SetId(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		m    *Farm
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

func TestFarm_GetKind(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want string
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: "frm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetKind())
		})
	}
}

func TestFarm_SetKind(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Farm
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

func TestFarm_GetName(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want string
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: "TestFarm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetName())
		})
	}
}

func TestFarm_SetName(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		m    *Farm
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetName(tt.args.val)
		})
	}
}

func TestFarm_GetOwners(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want []PersonAssociation
	}{
		{
			name: "valid farm",
			m:    ValidFarm,
			want: []PersonAssociation{
				ValidPersonAssociation,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetOwners())
		})
	}
}

func TestFarm_SetOwners(t *testing.T) {
	type args struct {
		val []PersonAssociation
	}
	tests := []struct {
		name string
		m    *Farm
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetOwners(tt.args.val)
		})
	}
}

func TestFarm_GetPeople(t *testing.T) {
	tests := []struct {
		name string
		m    Farm
		want []PersonAssociation
	}{

		{
			name: "valid farm",
			m:    ValidFarm,
			want: []PersonAssociation{
				ValidPersonAssociation,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.GetPeople())
		})
	}
}

func TestFarm_SetPeople(t *testing.T) {
	type args struct {
		val []PersonAssociation
	}
	tests := []struct {
		name string
		m    *Farm
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetPeople(tt.args.val)
		})
	}
}
