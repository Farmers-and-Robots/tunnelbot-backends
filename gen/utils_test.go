package openapi

import (
	"reflect"
	"testing"
	"time"
)

func TestNewNullableBool(t *testing.T) {
	type args struct {
		val *bool
	}
	tests := []struct {
		name string
		args args
		want *NullableBool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableBool(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableFloat32(t *testing.T) {
	type args struct {
		val *float32
	}
	tests := []struct {
		name string
		args args
		want *NullableFloat32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableFloat32(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableFloat64(t *testing.T) {
	type args struct {
		val *float64
	}
	tests := []struct {
		name string
		args args
		want *NullableFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableFloat64(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableInt(t *testing.T) {
	type args struct {
		val *int
	}
	tests := []struct {
		name string
		args args
		want *NullableInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableInt(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableInt32(t *testing.T) {
	type args struct {
		val *int32
	}
	tests := []struct {
		name string
		args args
		want *NullableInt32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableInt32(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableInt64(t *testing.T) {
	type args struct {
		val *int64
	}
	tests := []struct {
		name string
		args args
		want *NullableInt64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableInt64(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableString(t *testing.T) {
	type args struct {
		val *string
	}
	tests := []struct {
		name string
		args args
		want *NullableString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableString(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableTime(t *testing.T) {
	type args struct {
		val *time.Time
	}
	tests := []struct {
		name string
		args args
		want *NullableTime
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableTime(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableBool_Get(t *testing.T) {
	type fields struct {
		value *bool
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableBool{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableBool_IsSet(t *testing.T) {
	type fields struct {
		value *bool
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
			v := NullableBool{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableBool_MarshalJSON(t *testing.T) {
	type fields struct {
		value *bool
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
			v := NullableBool{
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

func TestNullableBool_Set(t *testing.T) {
	type fields struct {
		value *bool
		isSet bool
	}
	type args struct {
		val *bool
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
			v := &NullableBool{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableBool_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *bool
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
			v := &NullableBool{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableBool_Unset(t *testing.T) {
	type fields struct {
		value *bool
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
			v := &NullableBool{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableFloat32_Get(t *testing.T) {
	type fields struct {
		value *float32
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableFloat32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat32_IsSet(t *testing.T) {
	type fields struct {
		value *float32
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
			v := NullableFloat32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat32_MarshalJSON(t *testing.T) {
	type fields struct {
		value *float32
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
			v := NullableFloat32{
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

func TestNullableFloat32_Set(t *testing.T) {
	type fields struct {
		value *float32
		isSet bool
	}
	type args struct {
		val *float32
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
			v := &NullableFloat32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableFloat32_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *float32
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
			v := &NullableFloat32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableFloat32_Unset(t *testing.T) {
	type fields struct {
		value *float32
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
			v := &NullableFloat32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableFloat64_Get(t *testing.T) {
	type fields struct {
		value *float64
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableFloat64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat64_IsSet(t *testing.T) {
	type fields struct {
		value *float64
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
			v := NullableFloat64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat64_MarshalJSON(t *testing.T) {
	type fields struct {
		value *float64
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
			v := NullableFloat64{
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

func TestNullableFloat64_Set(t *testing.T) {
	type fields struct {
		value *float64
		isSet bool
	}
	type args struct {
		val *float64
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
			v := &NullableFloat64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableFloat64_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *float64
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
			v := &NullableFloat64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableFloat64_Unset(t *testing.T) {
	type fields struct {
		value *float64
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
			v := &NullableFloat64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableInt32_Get(t *testing.T) {
	type fields struct {
		value *int32
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableInt32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt32_IsSet(t *testing.T) {
	type fields struct {
		value *int32
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
			v := NullableInt32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt32_MarshalJSON(t *testing.T) {
	type fields struct {
		value *int32
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
			v := NullableInt32{
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

func TestNullableInt32_Set(t *testing.T) {
	type fields struct {
		value *int32
		isSet bool
	}
	type args struct {
		val *int32
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
			v := &NullableInt32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableInt32_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *int32
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
			v := &NullableInt32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableInt32_Unset(t *testing.T) {
	type fields struct {
		value *int32
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
			v := &NullableInt32{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableInt64_Get(t *testing.T) {
	type fields struct {
		value *int64
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableInt64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt64_IsSet(t *testing.T) {
	type fields struct {
		value *int64
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
			v := NullableInt64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt64_MarshalJSON(t *testing.T) {
	type fields struct {
		value *int64
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
			v := NullableInt64{
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

func TestNullableInt64_Set(t *testing.T) {
	type fields struct {
		value *int64
		isSet bool
	}
	type args struct {
		val *int64
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
			v := &NullableInt64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableInt64_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *int64
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
			v := &NullableInt64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableInt64_Unset(t *testing.T) {
	type fields struct {
		value *int64
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
			v := &NullableInt64{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableInt_Get(t *testing.T) {
	type fields struct {
		value *int
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableInt{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt_IsSet(t *testing.T) {
	type fields struct {
		value *int
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
			v := NullableInt{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt_MarshalJSON(t *testing.T) {
	type fields struct {
		value *int
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
			v := NullableInt{
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

func TestNullableInt_Set(t *testing.T) {
	type fields struct {
		value *int
		isSet bool
	}
	type args struct {
		val *int
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
			v := &NullableInt{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableInt_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *int
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
			v := &NullableInt{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableInt_Unset(t *testing.T) {
	type fields struct {
		value *int
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
			v := &NullableInt{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableString_Get(t *testing.T) {
	type fields struct {
		value *string
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableString{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableString_IsSet(t *testing.T) {
	type fields struct {
		value *string
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
			v := NullableString{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableString_MarshalJSON(t *testing.T) {
	type fields struct {
		value *string
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
			v := NullableString{
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

func TestNullableString_Set(t *testing.T) {
	type fields struct {
		value *string
		isSet bool
	}
	type args struct {
		val *string
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
			v := &NullableString{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableString_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *string
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
			v := &NullableString{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableString_Unset(t *testing.T) {
	type fields struct {
		value *string
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
			v := &NullableString{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestNullableTime_Get(t *testing.T) {
	type fields struct {
		value *time.Time
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableTime{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTime_IsSet(t *testing.T) {
	type fields struct {
		value *time.Time
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
			v := NullableTime{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTime_MarshalJSON(t *testing.T) {
	type fields struct {
		value *time.Time
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
			v := NullableTime{
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

func TestNullableTime_Set(t *testing.T) {
	type fields struct {
		value *time.Time
		isSet bool
	}
	type args struct {
		val *time.Time
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
			v := &NullableTime{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *time.Time
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
			v := &NullableTime{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableTime_Unset(t *testing.T) {
	type fields struct {
		value *time.Time
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
			v := &NullableTime{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestPtrBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrBool(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrFloat32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrFloat64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrInt(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrInt32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrInt64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrString(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrTime(t *testing.T) {
	type args struct {
		v time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrTime(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
