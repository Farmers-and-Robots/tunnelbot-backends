package openapi

import (
	"context"
	_nethttp "net/http"
	"reflect"
	"testing"
)

func TestApiGetPeopleRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *PeopleApiService
	}
	tests := []struct {
		name   string
		fields fields
		want   []Person
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetPeopleRequest{
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

func TestApiGetPersonByIdRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *PeopleApiService
		personId   int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Person
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetPersonByIdRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				personId:   tt.fields.personId,
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

func TestPeopleApiService_GetPeople(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    PeopleApiService
		args args
		want ApiGetPeopleRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetPeople(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleApiService_GetPeopleExecute(t *testing.T) {
	type args struct {
		r ApiGetPeopleRequest
	}
	tests := []struct {
		name  string
		a     PeopleApiService
		args  args
		want  []Person
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetPeopleExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeopleExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPeopleExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetPeopleExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPeopleApiService_GetPersonById(t *testing.T) {
	type args struct {
		ctx      context.Context
		personId int64
	}
	tests := []struct {
		name string
		a    PeopleApiService
		args args
		want ApiGetPersonByIdRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetPersonById(tt.args.ctx, tt.args.personId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPersonById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleApiService_GetPersonByIdExecute(t *testing.T) {
	type args struct {
		r ApiGetPersonByIdRequest
	}
	tests := []struct {
		name  string
		a     PeopleApiService
		args  args
		want  Person
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetPersonByIdExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPersonByIdExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPersonByIdExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetPersonByIdExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
