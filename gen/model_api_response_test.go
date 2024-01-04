package openapi

import (
	"reflect"
	"testing"
)

func TestApiResponse_GetCode(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.GetCode(); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_GetCodeOk(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *int32
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			got, got1 := o.GetCodeOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCodeOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCodeOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestApiResponse_GetMessage(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_GetMessageOk(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			got, got1 := o.GetMessageOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMessageOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetMessageOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestApiResponse_GetType(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_GetTypeOk(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			got, got1 := o.GetTypeOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTypeOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestApiResponse_HasCode(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.HasCode(); got != tt.want {
				t.Errorf("HasCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_HasMessage(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.HasMessage(); got != tt.want {
				t.Errorf("HasMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_HasType(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			if got := o.HasType(); got != tt.want {
				t.Errorf("HasType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
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

func TestApiResponse_SetCode(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
	}
	type args struct {
		v int32
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			o.SetCode(tt.args.v)
		})
	}
}

func TestApiResponse_SetMessage(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			o.SetMessage(tt.args.v)
		})
	}
}

func TestApiResponse_SetType(t *testing.T) {
	type fields struct {
		Code    *int32
		Type    *string
		Message *string
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
			o := &ApiResponse{
				Code:    tt.fields.Code,
				Type:    tt.fields.Type,
				Message: tt.fields.Message,
			}
			o.SetType(tt.args.v)
		})
	}
}

func TestNewApiResponse(t *testing.T) {
	tests := []struct {
		name string
		want *ApiResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApiResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApiResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewApiResponseWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *ApiResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApiResponseWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApiResponseWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableApiResponse(t *testing.T) {
	type args struct {
		val *ApiResponse
	}
	tests := []struct {
		name string
		args args
		want *NullableApiResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableApiResponse(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableApiResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableApiResponse_Get(t *testing.T) {
	type fields struct {
		value *ApiResponse
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *ApiResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableApiResponse{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableApiResponse_IsSet(t *testing.T) {
	type fields struct {
		value *ApiResponse
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
			v := NullableApiResponse{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableApiResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		value *ApiResponse
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
			v := NullableApiResponse{
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

func TestNullableApiResponse_Set(t *testing.T) {
	type fields struct {
		value *ApiResponse
		isSet bool
	}
	type args struct {
		val *ApiResponse
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
			v := &NullableApiResponse{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableApiResponse_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *ApiResponse
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
			v := &NullableApiResponse{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableApiResponse_Unset(t *testing.T) {
	type fields struct {
		value *ApiResponse
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
			v := &NullableApiResponse{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
