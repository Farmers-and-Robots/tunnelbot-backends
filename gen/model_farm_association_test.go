package openapi

import (
	"reflect"
	"testing"
)

func TestFarmAssociation_GetFarmId(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.GetFarmId(); got != tt.want {
				t.Errorf("GetFarmId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_GetFarmIdOk(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			got, got1 := o.GetFarmIdOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmIdOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFarmIdOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarmAssociation_GetFarmName(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.GetFarmName(); got != tt.want {
				t.Errorf("GetFarmName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_GetFarmNameOk(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			got, got1 := o.GetFarmNameOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmNameOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFarmNameOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarmAssociation_GetRole(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
	}
	tests := []struct {
		name   string
		fields fields
		want   Role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_GetRoleOk(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
	}
	tests := []struct {
		name   string
		fields fields
		want   *Role
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			got, got1 := o.GetRoleOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetRoleOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarmAssociation_HasFarmId(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.HasFarmId(); got != tt.want {
				t.Errorf("HasFarmId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_HasFarmName(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.HasFarmName(); got != tt.want {
				t.Errorf("HasFarmName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_HasRole(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			if got := o.HasRole(); got != tt.want {
				t.Errorf("HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmAssociation_MarshalJSON(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
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

func TestFarmAssociation_SetFarmId(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			o.SetFarmId(tt.args.v)
		})
	}
}

func TestFarmAssociation_SetFarmName(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			o.SetFarmName(tt.args.v)
		})
	}
}

func TestFarmAssociation_SetRole(t *testing.T) {
	type fields struct {
		FarmId   *int64
		FarmName *string
		Role     *Role
	}
	type args struct {
		v Role
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
			o := &FarmAssociation{
				FarmId:   tt.fields.FarmId,
				FarmName: tt.fields.FarmName,
				Role:     tt.fields.Role,
			}
			o.SetRole(tt.args.v)
		})
	}
}

func TestNewFarmAssociation(t *testing.T) {
	tests := []struct {
		name string
		want *FarmAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFarmAssociation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFarmAssociation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFarmAssociationWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *FarmAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFarmAssociationWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFarmAssociationWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableFarmAssociation(t *testing.T) {
	type args struct {
		val *FarmAssociation
	}
	tests := []struct {
		name string
		args args
		want *NullableFarmAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableFarmAssociation(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableFarmAssociation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarmAssociation_Get(t *testing.T) {
	type fields struct {
		value *FarmAssociation
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *FarmAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableFarmAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarmAssociation_IsSet(t *testing.T) {
	type fields struct {
		value *FarmAssociation
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
			v := NullableFarmAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarmAssociation_MarshalJSON(t *testing.T) {
	type fields struct {
		value *FarmAssociation
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
			v := NullableFarmAssociation{
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

func TestNullableFarmAssociation_Set(t *testing.T) {
	type fields struct {
		value *FarmAssociation
		isSet bool
	}
	type args struct {
		val *FarmAssociation
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
			v := &NullableFarmAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableFarmAssociation_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *FarmAssociation
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
			v := &NullableFarmAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableFarmAssociation_Unset(t *testing.T) {
	type fields struct {
		value *FarmAssociation
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
			v := &NullableFarmAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
