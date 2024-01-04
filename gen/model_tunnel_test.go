package openapi

import (
	"reflect"
	"testing"
)

func TestNewNullableTunnel(t *testing.T) {
	type args struct {
		val *Tunnel
	}
	tests := []struct {
		name string
		args args
		want *NullableTunnel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableTunnel(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTunnel(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Tunnel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTunnel(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTunnelWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Tunnel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTunnelWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTunnelWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTunnel_Get(t *testing.T) {
	type fields struct {
		value *Tunnel
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tunnel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableTunnel{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTunnel_IsSet(t *testing.T) {
	type fields struct {
		value *Tunnel
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
			v := NullableTunnel{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTunnel_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Tunnel
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
			v := NullableTunnel{
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

func TestNullableTunnel_Set(t *testing.T) {
	type fields struct {
		value *Tunnel
		isSet bool
	}
	type args struct {
		val *Tunnel
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
			v := &NullableTunnel{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableTunnel_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Tunnel
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
			v := &NullableTunnel{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableTunnel_Unset(t *testing.T) {
	type fields struct {
		value *Tunnel
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
			v := &NullableTunnel{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestTunnel_GetId(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			if got := o.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnel_GetIdOk(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			got, got1 := o.GetIdOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIdOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetIdOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTunnel_GetKind(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			if got := o.GetKind(); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnel_GetKindOk(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			got, got1 := o.GetKindOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKindOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetKindOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTunnel_GetName(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			if got := o.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnel_GetNameOk(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			got, got1 := o.GetNameOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNameOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNameOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTunnel_HasId(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			if got := o.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnel_HasKind(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			if got := o.HasKind(); got != tt.want {
				t.Errorf("HasKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnel_MarshalJSON(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
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

func TestTunnel_SetId(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
	}
	type args struct {
		v int64
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			o.SetId(tt.args.v)
		})
	}
}

func TestTunnel_SetKind(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			o.SetKind(tt.args.v)
		})
	}
}

func TestTunnel_SetName(t *testing.T) {
	type fields struct {
		Id   *int64
		Kind *string
		Name string
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
			o := &Tunnel{
				Id:   tt.fields.Id,
				Kind: tt.fields.Kind,
				Name: tt.fields.Name,
			}
			o.SetName(tt.args.v)
		})
	}
}
