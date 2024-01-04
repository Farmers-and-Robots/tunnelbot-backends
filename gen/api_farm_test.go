package openapi

import (
	"context"
	_nethttp "net/http"
	"reflect"
	"testing"
)

func TestApiAddFarmRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farm       *Farm
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farm:       tt.fields.farm,
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

func TestApiAddFarmRequest_Farm(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farm       *Farm
	}
	type args struct {
		farm Farm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiAddFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farm:       tt.fields.farm,
			}
			if got := r.Farm(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Farm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiAddPersonRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		farm       *Farm
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddPersonRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				farm:       tt.fields.farm,
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

func TestApiAddPersonRequest_Farm(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		farm       *Farm
	}
	type args struct {
		farm Farm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiAddPersonRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddPersonRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				farm:       tt.fields.farm,
			}
			if got := r.Farm(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Farm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiDeleteFarmRequest_ApiKey(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		apiKey     *string
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiDeleteFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiDeleteFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				apiKey:     tt.fields.apiKey,
			}
			if got := r.ApiKey(tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiDeleteFarmRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		apiKey     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *_nethttp.Response
		want1  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiDeleteFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				apiKey:     tt.fields.apiKey,
			}
			got, got1 := r.Execute()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestApiFindFarmByNameRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		name       *string
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiFindFarmByNameRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				name:       tt.fields.name,
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

func TestApiFindFarmByNameRequest_Name(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		name       *string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiFindFarmByNameRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiFindFarmByNameRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				name:       tt.fields.name,
			}
			if got := r.Name(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetFarmByIdRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetFarmByIdRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
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

func TestApiRemovePersonRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		farm       *Farm
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiRemovePersonRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				farm:       tt.fields.farm,
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

func TestApiRemovePersonRequest_Farm(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farmId     int64
		farm       *Farm
	}
	type args struct {
		farm Farm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiRemovePersonRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiRemovePersonRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farmId:     tt.fields.farmId,
				farm:       tt.fields.farm,
			}
			if got := r.Farm(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Farm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiUpdateFarmRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farm       *Farm
	}
	tests := []struct {
		name   string
		fields fields
		want   Farm
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farm:       tt.fields.farm,
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

func TestApiUpdateFarmRequest_Farm(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *FarmApiService
		farm       *Farm
	}
	type args struct {
		farm Farm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiUpdateFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateFarmRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				farm:       tt.fields.farm,
			}
			if got := r.Farm(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Farm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_AddFarm(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiAddFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddFarm(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_AddFarmExecute(t *testing.T) {
	type args struct {
		r ApiAddFarmRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.AddFarmExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFarmExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddFarmExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AddFarmExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFarmApiService_AddPerson(t *testing.T) {
	type args struct {
		ctx    context.Context
		farmId int64
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiAddPersonRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddPerson(tt.args.ctx, tt.args.farmId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_AddPersonExecute(t *testing.T) {
	type args struct {
		r ApiAddPersonRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.AddPersonExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPersonExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddPersonExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AddPersonExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFarmApiService_DeleteFarm(t *testing.T) {
	type args struct {
		ctx    context.Context
		farmId int64
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiDeleteFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.DeleteFarm(tt.args.ctx, tt.args.farmId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_DeleteFarmExecute(t *testing.T) {
	type args struct {
		r ApiDeleteFarmRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  *_nethttp.Response
		want1 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.DeleteFarmExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFarmExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DeleteFarmExecute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFarmApiService_FindFarmByName(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiFindFarmByNameRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.FindFarmByName(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFarmByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_FindFarmByNameExecute(t *testing.T) {
	type args struct {
		r ApiFindFarmByNameRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.FindFarmByNameExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFarmByNameExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindFarmByNameExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("FindFarmByNameExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFarmApiService_GetFarmById(t *testing.T) {
	type args struct {
		ctx    context.Context
		farmId int64
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiGetFarmByIdRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetFarmById(tt.args.ctx, tt.args.farmId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_GetFarmByIdExecute(t *testing.T) {
	type args struct {
		r ApiGetFarmByIdRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetFarmByIdExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFarmByIdExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetFarmByIdExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetFarmByIdExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFarmApiService_RemovePerson(t *testing.T) {
	type args struct {
		ctx    context.Context
		farmId int64
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiRemovePersonRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.RemovePerson(tt.args.ctx, tt.args.farmId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovePerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_RemovePersonExecute(t *testing.T) {
	type args struct {
		r ApiRemovePersonRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.RemovePersonExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovePersonExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RemovePersonExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("RemovePersonExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFarmApiService_UpdateFarm(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    FarmApiService
		args args
		want ApiUpdateFarmRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UpdateFarm(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmApiService_UpdateFarmExecute(t *testing.T) {
	type args struct {
		r ApiUpdateFarmRequest
	}
	tests := []struct {
		name  string
		a     FarmApiService
		args  args
		want  Farm
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.UpdateFarmExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateFarmExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UpdateFarmExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("UpdateFarmExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
