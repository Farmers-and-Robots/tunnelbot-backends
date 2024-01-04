package openapi

import (
	"context"
	_nethttp "net/http"
	"reflect"
	"testing"
)

func TestApiAddDeviceRequest_Device(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		device     *Device
	}
	type args struct {
		device Device
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiAddDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				device:     tt.fields.device,
			}
			if got := r.Device(tt.args.device); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Device() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiAddDeviceRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		device     *Device
	}
	tests := []struct {
		name   string
		fields fields
		want   Device
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiAddDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				device:     tt.fields.device,
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

func TestApiDeleteDeviceRequest_ApiKey(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		deviceId   int64
		apiKey     *string
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiDeleteDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiDeleteDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				deviceId:   tt.fields.deviceId,
				apiKey:     tt.fields.apiKey,
			}
			if got := r.ApiKey(tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiDeleteDeviceRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		deviceId   int64
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
			r := ApiDeleteDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				deviceId:   tt.fields.deviceId,
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

func TestApiGetDeviceByIdRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		deviceId   int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Device
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetDeviceByIdRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				deviceId:   tt.fields.deviceId,
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

func TestApiGetDevicesRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
	}
	tests := []struct {
		name   string
		fields fields
		want   []Device
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetDevicesRequest{
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

func TestApiUpdateDeviceRequest_Device(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		device     *Device
	}
	type args struct {
		device Device
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiUpdateDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				device:     tt.fields.device,
			}
			if got := r.Device(tt.args.device); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Device() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiUpdateDeviceRequest_Execute(t *testing.T) {
	type fields struct {
		ctx        context.Context
		ApiService *DeviceApiService
		device     *Device
	}
	tests := []struct {
		name   string
		fields fields
		want   Device
		want1  *_nethttp.Response
		want2  GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiUpdateDeviceRequest{
				ctx:        tt.fields.ctx,
				ApiService: tt.fields.ApiService,
				device:     tt.fields.device,
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

func TestDeviceApiService_AddDevice(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    DeviceApiService
		args args
		want ApiAddDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddDevice(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceApiService_AddDeviceExecute(t *testing.T) {
	type args struct {
		r ApiAddDeviceRequest
	}
	tests := []struct {
		name  string
		a     DeviceApiService
		args  args
		want  Device
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.AddDeviceExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDeviceExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddDeviceExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AddDeviceExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestDeviceApiService_DeleteDevice(t *testing.T) {
	type args struct {
		ctx      context.Context
		deviceId int64
	}
	tests := []struct {
		name string
		a    DeviceApiService
		args args
		want ApiDeleteDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.DeleteDevice(tt.args.ctx, tt.args.deviceId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceApiService_DeleteDeviceExecute(t *testing.T) {
	type args struct {
		r ApiDeleteDeviceRequest
	}
	tests := []struct {
		name  string
		a     DeviceApiService
		args  args
		want  *_nethttp.Response
		want1 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.DeleteDeviceExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDeviceExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DeleteDeviceExecute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeviceApiService_GetDeviceById(t *testing.T) {
	type args struct {
		ctx      context.Context
		deviceId int64
	}
	tests := []struct {
		name string
		a    DeviceApiService
		args args
		want ApiGetDeviceByIdRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetDeviceById(tt.args.ctx, tt.args.deviceId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeviceById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceApiService_GetDeviceByIdExecute(t *testing.T) {
	type args struct {
		r ApiGetDeviceByIdRequest
	}
	tests := []struct {
		name  string
		a     DeviceApiService
		args  args
		want  Device
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetDeviceByIdExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeviceByIdExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetDeviceByIdExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetDeviceByIdExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestDeviceApiService_GetDevices(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    DeviceApiService
		args args
		want ApiGetDevicesRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetDevices(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceApiService_GetDevicesExecute(t *testing.T) {
	type args struct {
		r ApiGetDevicesRequest
	}
	tests := []struct {
		name  string
		a     DeviceApiService
		args  args
		want  []Device
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.GetDevicesExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevicesExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetDevicesExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetDevicesExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestDeviceApiService_UpdateDevice(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		a    DeviceApiService
		args args
		want ApiUpdateDeviceRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UpdateDevice(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceApiService_UpdateDeviceExecute(t *testing.T) {
	type args struct {
		r ApiUpdateDeviceRequest
	}
	tests := []struct {
		name  string
		a     DeviceApiService
		args  args
		want  Device
		want1 *_nethttp.Response
		want2 GenericOpenAPIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.a.UpdateDeviceExecute(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateDeviceExecute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UpdateDeviceExecute() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("UpdateDeviceExecute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
