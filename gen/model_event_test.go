package openapi

import (
	"reflect"
	"testing"
	"time"
)

func TestEvent_GetDate(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.GetDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_GetDateOk(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			got, got1 := o.GetDateOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDateOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetDateOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvent_GetDescription(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.GetDescription(); got != tt.want {
				t.Errorf("GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_GetDescriptionOk(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			got, got1 := o.GetDescriptionOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDescriptionOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetDescriptionOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvent_GetId(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_GetIdOk(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
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

func TestEvent_GetKind(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.GetKind(); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_GetKindOk(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
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

func TestEvent_GetOperator(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.GetOperator(); got != tt.want {
				t.Errorf("GetOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_GetOperatorOk(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			got, got1 := o.GetOperatorOk()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOperatorOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetOperatorOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvent_HasDate(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.HasDate(); got != tt.want {
				t.Errorf("HasDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_HasDescription(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.HasDescription(); got != tt.want {
				t.Errorf("HasDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_HasId(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_HasKind(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.HasKind(); got != tt.want {
				t.Errorf("HasKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_HasOperator(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			if got := o.HasOperator(); got != tt.want {
				t.Errorf("HasOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_MarshalJSON(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
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

func TestEvent_SetDate(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			o.SetDate(tt.args.v)
		})
	}
}

func TestEvent_SetDescription(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			o.SetDescription(tt.args.v)
		})
	}
}

func TestEvent_SetId(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			o.SetId(tt.args.v)
		})
	}
}

func TestEvent_SetKind(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			o.SetKind(tt.args.v)
		})
	}
}

func TestEvent_SetOperator(t *testing.T) {
	type fields struct {
		Id          *int64
		Kind        *string
		Description *string
		Date        *time.Time
		Operator    *string
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
			o := &Event{
				Id:          tt.fields.Id,
				Kind:        tt.fields.Kind,
				Description: tt.fields.Description,
				Date:        tt.fields.Date,
				Operator:    tt.fields.Operator,
			}
			o.SetOperator(tt.args.v)
		})
	}
}

func TestNewEvent(t *testing.T) {
	tests := []struct {
		name string
		want *Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEventWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventWithDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventWithDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullableEvent(t *testing.T) {
	type args struct {
		val *Event
	}
	tests := []struct {
		name string
		args args
		want *NullableEvent
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullableEvent(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullableEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableEvent_Get(t *testing.T) {
	type fields struct {
		value *Event
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullableEvent{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableEvent_IsSet(t *testing.T) {
	type fields struct {
		value *Event
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
			v := NullableEvent{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if got := v.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableEvent_MarshalJSON(t *testing.T) {
	type fields struct {
		value *Event
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
			v := NullableEvent{
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

func TestNullableEvent_Set(t *testing.T) {
	type fields struct {
		value *Event
		isSet bool
	}
	type args struct {
		val *Event
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
			v := &NullableEvent{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Set(tt.args.val)
		})
	}
}

func TestNullableEvent_UnmarshalJSON(t *testing.T) {
	type fields struct {
		value *Event
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
			v := &NullableEvent{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullableEvent_Unset(t *testing.T) {
	type fields struct {
		value *Event
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
			v := &NullableEvent{
				value: tt.fields.value,
				isSet: tt.fields.isSet,
			}
			v.Unset()
		})
	}
}
