package openapi

import (
	"reflect"
	"testing"
)

func TestDevice_GetCategory(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
	}
	tests := []struct {
		name   string
		fields fields
		want   DeviceCategory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.GetCategory(); got != tt.want {
				t.Errorf("GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetCategoryOk(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeviceCategory
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			got, got1 := o.GetCategoryOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategoryOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCategoryOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDevice_GetId(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetIdOk(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
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

func TestDevice_GetKind(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.GetKind(); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetKindOk(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
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

func TestDevice_GetName(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetNameOk(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
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

func TestDevice_HasCategory(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.HasCategory(); got != tt.want {
				t.Errorf("HasCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_HasId(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_HasKind(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			if got := o.HasKind(); got != tt.want {
				t.Errorf("HasKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_MarshalJSON(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
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

func TestDevice_SetCategory(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
	}
	type args struct {
		v DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			o.SetCategory(tt.args.v)
		})
	}
}

func TestDevice_SetId(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			o.SetId(tt.args.v)
		})
	}
}

func TestDevice_SetKind(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			o.SetKind(tt.args.v)
		})
	}
}

func TestDevice_SetName(t *testing.T) {
	type fields struct {
		Id       *int64
		Kind     *string
		Name     string
		Category *DeviceCategory
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
			o := &Device{
				Id:       tt.fields.Id,
				Kind:     tt.fields.Kind,
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
			}
			o.SetName(tt.args.v)
		})
	}
}

func TestNewDevice(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Device
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDevice(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeviceWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Device
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeviceWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeviceWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableDevice(t *testing.T) {
	type args struct {
		val *Device
	}
	tests := []struct {
		name string
		args args
		want *NullableDevice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableDevice(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableDevice_Get(t *testing.T) {
	type fields struct {
		value *Device
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Device
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableDevice{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableDevice_IsSet(t *testing.T) {
	type fields struct {
		value *Device
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
			v := NullableDevice{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableDevice_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Device
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
			v := NullableDevice{
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

func TestNullableDevice_Set(t *testing.T) {
	type fields struct {
		value *Device
		isSet bool
	}
	type args struct {
		val *Device
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
			v := &NullableDevice{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableDevice_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Device
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
			v := &NullableDevice{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableDevice_Unset(t *testing.T) {
	type fields struct {
		value *Device
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
			v := &NullableDevice{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
