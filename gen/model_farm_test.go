package openapi

import (
	"reflect"
	"testing"
	"time"
)

func TestFarm_GetAddress(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetAddressOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   *Address
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			got, got1 := o.GetAddressOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddressOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAddressOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarm_GetCreateDate(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetCreateDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCreateDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetCreateDateOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   *time.Time
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			got, got1 := o.GetCreateDateOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCreateDateOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCreateDateOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarm_GetId(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetIdOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
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

func TestFarm_GetKind(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetKind(); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetKindOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
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

func TestFarm_GetName(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetNameOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
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

func TestFarm_GetOwners(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   []PersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetOwners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetOwnersOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]PersonAssociation
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			got, got1 := o.GetOwnersOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnersOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetOwnersOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarm_GetPeople(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   []PersonAssociation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.GetPeople(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_GetPeopleOk(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]PersonAssociation
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			got, got1 := o.GetPeopleOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeopleOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPeopleOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarm_HasAddress(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasAddress(); got != tt.want {
				t.Errorf("HasAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasCreateDate(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasCreateDate(); got != tt.want {
				t.Errorf("HasCreateDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasId(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasKind(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasKind(); got != tt.want {
				t.Errorf("HasKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasName(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasName(); got != tt.want {
				t.Errorf("HasName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasOwners(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasOwners(); got != tt.want {
				t.Errorf("HasOwners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_HasPeople(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			if got := o.HasPeople(); got != tt.want {
				t.Errorf("HasPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarm_MarshalJSON(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
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

func TestFarm_SetAddress(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	type args struct {
		v Address
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetAddress(tt.args.v)
		})
	}
}

func TestFarm_SetCreateDate(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	type args struct {
		v time.Time
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetCreateDate(tt.args.v)
		})
	}
}

func TestFarm_SetId(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetId(tt.args.v)
		})
	}
}

func TestFarm_SetKind(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetKind(tt.args.v)
		})
	}
}

func TestFarm_SetName(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetName(tt.args.v)
		})
	}
}

func TestFarm_SetOwners(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	type args struct {
		v []PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetOwners(tt.args.v)
		})
	}
}

func TestFarm_SetPeople(t *testing.T) {
	type fields struct {
		Id         *int64
		Kind       *string
		Name       *string
		CreateDate *time.Time
		Address    *Address
		Owners     *[]PersonAssociation
		People     *[]PersonAssociation
	}
	type args struct {
		v []PersonAssociation
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
			o := &Farm{
				Id:         tt.fields.Id,
				Kind:       tt.fields.Kind,
				Name:       tt.fields.Name,
				CreateDate: tt.fields.CreateDate,
				Address:    tt.fields.Address,
				Owners:     tt.fields.Owners,
				People:     tt.fields.People,
			}
			o.SetPeople(tt.args.v)
		})
	}
}

func TestNewFarm(t *testing.T) {
	tests := []struct {
		name string
		want *Farm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFarm(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFarmWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Farm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFarmWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFarmWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableFarm(t *testing.T) {
	type args struct {
		val *Farm
	}
	tests := []struct {
		name string
		args args
		want *NullableFarm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableFarm(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarm_Get(t *testing.T) {
	type fields struct {
		value *Farm
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Farm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableFarm{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarm_IsSet(t *testing.T) {
	type fields struct {
		value *Farm
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
			v := NullableFarm{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFarm_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Farm
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
			v := NullableFarm{
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

func TestNullableFarm_Set(t *testing.T) {
	type fields struct {
		value *Farm
		isSet bool
	}
	type args struct {
		val *Farm
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
			v := &NullableFarm{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableFarm_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Farm
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
			v := &NullableFarm{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableFarm_Unset(t *testing.T) {
	type fields struct {
		value *Farm
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
			v := &NullableFarm{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
