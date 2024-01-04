# \TunnelApi

All URIs are relative to *https://tunnelbot.swagger.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTunnel**](TunnelApi.md#AddTunnel) | **Post** /tunnels | Add a new tunnel to the db
[**DeleteTunnel**](TunnelApi.md#DeleteTunnel) | **Delete** /tunnels/{tunnelId} | Deletes a tunnel
[**GetTunnelById**](TunnelApi.md#GetTunnelById) | **Get** /tunnels/{tunnelId} | Find tunnel by id
[**GetTunnels**](TunnelApi.md#GetTunnels) | **Get** /tunnels | Return the tunnels on a farm
[**UpdateTunnel**](TunnelApi.md#UpdateTunnel) | **Put** /tunnels | Update an existing tunnel



## AddTunnel

> Tunnel AddTunnel(ctx).Tunnel(tunnel).Execute()

Add a new tunnel to the db



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tunnel := *openapiclient.NewTunnel("tunnel01") // Tunnel | Create a new tunnel in the db

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TunnelApi.AddTunnel(context.Background()).Tunnel(tunnel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelApi.AddTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTunnel`: Tunnel
    fmt.Fprintf(os.Stdout, "Response from `TunnelApi.AddTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tunnel** | [**Tunnel**](Tunnel.md) | Create a new tunnel in the db | 

### Return type

[**Tunnel**](Tunnel.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTunnel

> DeleteTunnel(ctx, tunnelId).ApiKey(apiKey).Execute()

Deletes a tunnel



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tunnelId := int64(789) // int64 | Tunnel id to delete
    apiKey := "apiKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TunnelApi.DeleteTunnel(context.Background(), tunnelId).ApiKey(apiKey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelApi.DeleteTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tunnelId** | **int64** | Tunnel id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunnelById

> Tunnel GetTunnelById(ctx, tunnelId).Execute()

Find tunnel by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tunnelId := int64(789) // int64 | id of tunnel to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TunnelApi.GetTunnelById(context.Background(), tunnelId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelApi.GetTunnelById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTunnelById`: Tunnel
    fmt.Fprintf(os.Stdout, "Response from `TunnelApi.GetTunnelById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tunnelId** | **int64** | id of tunnel to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tunnel**](Tunnel.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunnels

> []Tunnel GetTunnels(ctx).Execute()

Return the tunnels on a farm



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TunnelApi.GetTunnels(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelApi.GetTunnels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTunnels`: []Tunnel
    fmt.Fprintf(os.Stdout, "Response from `TunnelApi.GetTunnels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelsRequest struct via the builder pattern


### Return type

[**[]Tunnel**](Tunnel.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTunnel

> Tunnel UpdateTunnel(ctx).Tunnel(tunnel).Execute()

Update an existing tunnel



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tunnel := *openapiclient.NewTunnel("tunnel01") // Tunnel | Update an existent tunnel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TunnelApi.UpdateTunnel(context.Background()).Tunnel(tunnel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelApi.UpdateTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTunnel`: Tunnel
    fmt.Fprintf(os.Stdout, "Response from `TunnelApi.UpdateTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tunnel** | [**Tunnel**](Tunnel.md) | Update an existent tunnel | 

### Return type

[**Tunnel**](Tunnel.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

