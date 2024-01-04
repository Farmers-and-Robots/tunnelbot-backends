package openapi

import (
	"reflect"
	"testing"
)

func TestNewNullableRole(t *testing.T) {
	type args struct {
		val *Role
	}
	tests := []struct {
		name string
		args args
		want *NullableRole
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableRole(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableRole_Get(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableRole_IsSet(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableRole_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableRole_Set(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	type args struct {
		val *Role
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableRole_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableRole_Unset(t *testing.T) {
	type fields struct {
		value *Role
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NullableRole{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestRole_Ptr(t *testing.T) {
	tests := []struct {
		name string
		v    Role
		want *Role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Ptr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRole_UnmarshalJSON(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		v       Role
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
