package openapi

import (
	"reflect"
	"testing"
)

func TestNewNullablePersonAssociation(t *testing.T) {
	type args struct {
		val *PersonAssociation
	}
	tests := []struct {
		name string
		args args
		want *NullablePersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullablePersonAssociation(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullablePersonAssociation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonAssociation(t *testing.T) {
	tests := []struct {
		name string
		want *PersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersonAssociation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersonAssociation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonAssociationWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *PersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersonAssociationWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersonAssociationWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePersonAssociation_Get(t *testing.T) {
	type fields struct {
		value *PersonAssociation
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *PersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullablePersonAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePersonAssociation_IsSet(t *testing.T) {
	type fields struct {
		value *PersonAssociation
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
			v := NullablePersonAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePersonAssociation_MarshalJSON(t *testing.T) {
	type fields struct {
		value *PersonAssociation
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
			v := NullablePersonAssociation{
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

func TestNullablePersonAssociation_Set(t *testing.T) {
	type fields struct {
		value *PersonAssociation
		isSet bool
	}
	type args struct {
		val *PersonAssociation
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
			v := &NullablePersonAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullablePersonAssociation_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *PersonAssociation
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
			v := &NullablePersonAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullablePersonAssociation_Unset(t *testing.T) {
	type fields struct {
		value *PersonAssociation
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
			v := &NullablePersonAssociation{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestPersonAssociation_GetDisplayName(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.GetDisplayName(); got != tt.want {
				t.Errorf("GetDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_GetDisplayNameOk(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			got, got1 := o.GetDisplayNameOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDisplayNameOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetDisplayNameOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPersonAssociation_GetEmail(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.GetEmail(); got != tt.want {
				t.Errorf("GetEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_GetEmailOk(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			got, got1 := o.GetEmailOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmailOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetEmailOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPersonAssociation_GetPersonId(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.GetPersonId(); got != tt.want {
				t.Errorf("GetPersonId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_GetPersonIdOk(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			got, got1 := o.GetPersonIdOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPersonIdOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPersonIdOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPersonAssociation_GetRole(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_GetRoleOk(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
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

func TestPersonAssociation_HasDisplayName(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.HasDisplayName(); got != tt.want {
				t.Errorf("HasDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_HasEmail(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.HasEmail(); got != tt.want {
				t.Errorf("HasEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_HasPersonId(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.HasPersonId(); got != tt.want {
				t.Errorf("HasPersonId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_HasRole(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			if got := o.HasRole(); got != tt.want {
				t.Errorf("HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonAssociation_MarshalJSON(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
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

func TestPersonAssociation_SetDisplayName(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			o.SetDisplayName(tt.args.v)
		})
	}
}

func TestPersonAssociation_SetEmail(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			o.SetEmail(tt.args.v)
		})
	}
}

func TestPersonAssociation_SetPersonId(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			o.SetPersonId(tt.args.v)
		})
	}
}

func TestPersonAssociation_SetRole(t *testing.T) {
	type fields struct {
		PersonId    *int64
		DisplayName *string
		Email       *string
		Role        *Role
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
			o := &PersonAssociation{
				PersonId:    tt.fields.PersonId,
				DisplayName: tt.fields.DisplayName,
				Email:       tt.fields.Email,
				Role:        tt.fields.Role,
			}
			o.SetRole(tt.args.v)
		})
	}
}
