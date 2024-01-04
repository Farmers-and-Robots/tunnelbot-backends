package openapi

import (
	"context"
	_nethttp "net/http"
	"reflect"
	"testing"
)

func TestApiAddTunnelRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnel     *Tunnel
	}
	tests := []struct {
		name   string
		fields fields
		want   Tunnel
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnel:     tt.fields.tunnel,
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

func TestApiAddTunnelRequest_Tunnel(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnel     *Tunnel
	}
	type args struct {
		tunnel Tunnel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiAddTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnel:     tt.fields.tunnel,
			}
			if got := r.Tunnel(tt.args.tunnel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiDeleteTunnelRequest_ApiKey(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnelId   int64
		apiKey     *string
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiDeleteTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiDeleteTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnelId:   tt.fields.tunnelId,
				apiKey:     tt.fields.apiKey,
			}
			if got := r.ApiKey(tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiDeleteTunnelRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnelId   int64
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
			r := ApiDeleteTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnelId:   tt.fields.tunnelId,
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

func TestApiGetTunnelByIdRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnelId   int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Tunnel
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetTunnelByIdRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnelId:   tt.fields.tunnelId,
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

func TestApiGetTunnelsRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
	}
	tests := []struct {
		name   string
		fields fields
		want   []Tunnel
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetTunnelsRequest{
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

func TestApiUpdateTunnelRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnel     *Tunnel
	}
	tests := []struct {
		name   string
		fields fields
		want   Tunnel
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnel:     tt.fields.tunnel,
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

func TestApiUpdateTunnelRequest_Tunnel(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *TunnelApiService
		tunnel     *Tunnel
	}
	type args struct {
		tunnel Tunnel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiUpdateTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateTunnelRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				tunnel:     tt.fields.tunnel,
			}
			if got := r.Tunnel(tt.args.tunnel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_AddTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    TunnelApiService
		args args
		want ApiAddTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddTunnel(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_AddTunnelExecute(t *testing.T) {
	type args struct {
		r ApiAddTunnelRequest
	}
	tests := []struct {
		name  string
		a     TunnelApiService
		args  args
		want  Tunnel
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.AddTunnelExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTunnelExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddTunnelExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AddTunnelExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTunnelApiService_DeleteTunnel(t *testing.T) {
	type args struct {
		ctx      context.Context
		tunnelId int64
	}
	tests := []struct {
		name string
		a    TunnelApiService
		args args
		want ApiDeleteTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.DeleteTunnel(tt.args.ctx, tt.args.tunnelId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_DeleteTunnelExecute(t *testing.T) {
	type args struct {
		r ApiDeleteTunnelRequest
	}
	tests := []struct {
		name  string
		a     TunnelApiService
		args  args
		want  *_nethttp.Response
		want1 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.DeleteTunnelExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTunnelExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DeleteTunnelExecute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTunnelApiService_GetTunnelById(t *testing.T) {
	type args struct {
		ctx      context.Context
		tunnelId int64
	}
	tests := []struct {
		name string
		a    TunnelApiService
		args args
		want ApiGetTunnelByIdRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetTunnelById(tt.args.ctx, tt.args.tunnelId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTunnelById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_GetTunnelByIdExecute(t *testing.T) {
	type args struct {
		r ApiGetTunnelByIdRequest
	}
	tests := []struct {
		name  string
		a     TunnelApiService
		args  args
		want  Tunnel
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetTunnelByIdExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTunnelByIdExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetTunnelByIdExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetTunnelByIdExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTunnelApiService_GetTunnels(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    TunnelApiService
		args args
		want ApiGetTunnelsRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetTunnels(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTunnels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_GetTunnelsExecute(t *testing.T) {
	type args struct {
		r ApiGetTunnelsRequest
	}
	tests := []struct {
		name  string
		a     TunnelApiService
		args  args
		want  []Tunnel
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetTunnelsExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTunnelsExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetTunnelsExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetTunnelsExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTunnelApiService_UpdateTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    TunnelApiService
		args args
		want ApiUpdateTunnelRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UpdateTunnel(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTunnelApiService_UpdateTunnelExecute(t *testing.T) {
	type args struct {
		r ApiUpdateTunnelRequest
	}
	tests := []struct {
		name  string
		a     TunnelApiService
		args  args
		want  Tunnel
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.UpdateTunnelExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTunnelExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UpdateTunnelExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("UpdateTunnelExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
