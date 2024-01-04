package openapi

import (
	"reflect"
	"testing"
)

func TestAddress_GetCity(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.GetCity(); got != tt.want {
				t.Errorf("GetCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_GetCityOk(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			got, got1 := o.GetCityOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCityOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCityOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddress_GetState(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_GetStateOk(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			got, got1 := o.GetStateOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStateOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetStateOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddress_GetStreet(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.GetStreet(); got != tt.want {
				t.Errorf("GetStreet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_GetStreetOk(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			got, got1 := o.GetStreetOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStreetOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetStreetOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddress_GetZip(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.GetZip(); got != tt.want {
				t.Errorf("GetZip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_GetZipOk(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			got, got1 := o.GetZipOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetZipOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetZipOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddress_HasCity(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.HasCity(); got != tt.want {
				t.Errorf("HasCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_HasState(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.HasState(); got != tt.want {
				t.Errorf("HasState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_HasStreet(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.HasStreet(); got != tt.want {
				t.Errorf("HasStreet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_HasZip(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			if got := o.HasZip(); got != tt.want {
				t.Errorf("HasZip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_MarshalJSON(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
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
			o := Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			got, err := o.MarshalJSON()
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

func TestAddress_SetCity(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	type args struct {
		v string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			o.SetCity(tt.args.v)
		})
	}
}

func TestAddress_SetState(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	type args struct {
		v string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			o.SetState(tt.args.v)
		})
	}
}

func TestAddress_SetStreet(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	type args struct {
		v string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			o.SetStreet(tt.args.v)
		})
	}
}

func TestAddress_SetZip(t *testing.T) {
	type fields struct {
		Street *string
		City   *string
		State  *string
		Zip    *string
	}
	type args struct {
		v string
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
			o := &Address{
				Street: tt.fields.Street,
				City:   tt.fields.City,
				State:  tt.fields.State,
				Zip:    tt.fields.Zip,
			}
			o.SetZip(tt.args.v)
		})
	}
}

func TestNewAddress(t *testing.T) {
	tests := []struct {
		name string
		want *Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddressWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddressWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddressWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableAddress(t *testing.T) {
	type args struct {
		val *Address
	}
	tests := []struct {
		name string
		args args
		want *NullableAddress
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableAddress(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableAddress_Get(t *testing.T) {
	type fields struct {
		value *Address
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableAddress{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableAddress_IsSet(t *testing.T) {
	type fields struct {
		value *Address
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
			v := NullableAddress{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableAddress_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Address
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
			v := NullableAddress{
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

func TestNullableAddress_Set(t *testing.T) {
	type fields struct {
		value *Address
		isSet bool
	}
	type args struct {
		val *Address
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
			v := &NullableAddress{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableAddress_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Address
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
			v := &NullableAddress{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableAddress_Unset(t *testing.T) {
	type fields struct {
		value *Address
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
			v := &NullableAddress{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
