package openapi

import (
	"context"
	_nethttp "net/http"
	"reflect"
	"testing"
)

func TestApiGetEventByIdRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *EventsApiService
		eventId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Event
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetEventByIdRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				eventId:    tt.fields.eventId,
			}
			got, got1, got2 := r.Execute()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Execute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestApiGetEventsRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *EventsApiService
	}
	tests := []struct {
		name   string
		fields fields
		want   []Event
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetEventsRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
			}
			got, got1, got2 := r.Execute()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Execute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestEventsApiService_GetEventById(t *testing.T) {
	type args struct {
		ctx     context.Context
		eventId int64
	}
	tests := []struct {
		name string
		a    EventsApiService
		args args
		want ApiGetEventByIdRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEventById(tt.args.ctx, tt.args.eventId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventsApiService_GetEventByIdExecute(t *testing.T) {
	type args struct {
		r ApiGetEventByIdRequest
	}
	tests := []struct {
		name  string
		a     EventsApiService
		args  args
		want  Event
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetEventByIdExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventByIdExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetEventByIdExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetEventByIdExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestEventsApiService_GetEvents(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    EventsApiService
		args args
		want ApiGetEventsRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetEvents(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventsApiService_GetEventsExecute(t *testing.T) {
	type args struct {
		r ApiGetEventsRequest
	}
	tests := []struct {
		name  string
		a     EventsApiService
		args  args
		want  []Event
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetEventsExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventsExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetEventsExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetEventsExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
