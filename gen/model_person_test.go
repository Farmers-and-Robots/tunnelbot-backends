package openapi

import (
	"reflect"
	"testing"
)

func TestNewNullablePerson(t *testing.T) {
	type args struct {
		val *Person
	}
	tests := []struct {
		name string
		args args
		want *NullablePerson
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullablePerson(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullablePerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPerson(t *testing.T) {
	tests := []struct {
		name string
		want *Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPerson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersonWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersonWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePerson_Get(t *testing.T) {
	type fields struct {
		value *Person
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullablePerson{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePerson_IsSet(t *testing.T) {
	type fields struct {
		value *Person
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
			v := NullablePerson{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullablePerson_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Person
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
			v := NullablePerson{
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

func TestNullablePerson_Set(t *testing.T) {
	type fields struct {
		value *Person
		isSet bool
	}
	type args struct {
		val *Person
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
			v := &NullablePerson{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullablePerson_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Person
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
			v := &NullablePerson{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullablePerson_Unset(t *testing.T) {
	type fields struct {
		value *Person
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
			v := &NullablePerson{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}

func TestPerson_GetDisplayName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetDisplayName(); got != tt.want {
				t.Errorf("GetDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetDisplayNameOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
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

func TestPerson_GetEmail(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetEmail(); got != tt.want {
				t.Errorf("GetEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetEmailOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
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

func TestPerson_GetFarmAssociations(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   []FarmAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetFarmAssociations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmAssociations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetFarmAssociationsOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]FarmAssociation
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			got, got1 := o.GetFarmAssociationsOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmAssociationsOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFarmAssociationsOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPerson_GetFirstName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetFirstName(); got != tt.want {
				t.Errorf("GetFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetFirstNameOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			got, got1 := o.GetFirstNameOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstNameOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFirstNameOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPerson_GetId(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetIdOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
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

func TestPerson_GetKind(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetKind(); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetKindOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
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

func TestPerson_GetLastName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetLastName(); got != tt.want {
				t.Errorf("GetLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetLastNameOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			got, got1 := o.GetLastNameOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastNameOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLastNameOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPerson_GetPhone(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.GetPhone(); got != tt.want {
				t.Errorf("GetPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetPhoneOk(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			got, got1 := o.GetPhoneOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhoneOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPhoneOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPerson_HasDisplayName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasDisplayName(); got != tt.want {
				t.Errorf("HasDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasEmail(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasEmail(); got != tt.want {
				t.Errorf("HasEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasFarmAssociations(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasFarmAssociations(); got != tt.want {
				t.Errorf("HasFarmAssociations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasFirstName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasFirstName(); got != tt.want {
				t.Errorf("HasFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasId(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasKind(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasKind(); got != tt.want {
				t.Errorf("HasKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasLastName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasLastName(); got != tt.want {
				t.Errorf("HasLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_HasPhone(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			if got := o.HasPhone(); got != tt.want {
				t.Errorf("HasPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_MarshalJSON(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
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

func TestPerson_SetDisplayName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetDisplayName(tt.args.v)
		})
	}
}

func TestPerson_SetEmail(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetEmail(tt.args.v)
		})
	}
}

func TestPerson_SetFarmAssociations(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
	}
	type args struct {
		v []FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetFarmAssociations(tt.args.v)
		})
	}
}

func TestPerson_SetFirstName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetFirstName(tt.args.v)
		})
	}
}

func TestPerson_SetId(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetId(tt.args.v)
		})
	}
}

func TestPerson_SetKind(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetKind(tt.args.v)
		})
	}
}

func TestPerson_SetLastName(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetLastName(tt.args.v)
		})
	}
}

func TestPerson_SetPhone(t *testing.T) {
	type fields struct {
		Id               *int64
		Kind             *string
		DisplayName      *string
		FirstName        *string
		LastName         *string
		Email            *string
		Phone            *string
		FarmAssociations *[]FarmAssociation
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
			o := &Person{
				Id:               tt.fields.Id,
				Kind:             tt.fields.Kind,
				DisplayName:      tt.fields.DisplayName,
				FirstName:        tt.fields.FirstName,
				LastName:         tt.fields.LastName,
				Email:            tt.fields.Email,
				Phone:            tt.fields.Phone,
				FarmAssociations: tt.fields.FarmAssociations,
			}
			o.SetPhone(tt.args.v)
		})
	}
}
