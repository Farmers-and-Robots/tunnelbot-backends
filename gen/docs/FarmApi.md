# \FarmApi

All URIs are relative to *https://tunnelbot.swagger.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFarm**](FarmApi.md#AddFarm) | **Post** /farm | Add a new farm to the db
[**AddPerson**](FarmApi.md#AddPerson) | **Put** /farm/addFarmAccess/{farmId} | Add farm access for an existing person
[**DeleteFarm**](FarmApi.md#DeleteFarm) | **Delete** /farm/{farmId} | Deletes a farm
[**FindFarmByName**](FarmApi.md#FindFarmByName) | **Get** /farm/findByName | Finds a farm by name
[**GetFarmById**](FarmApi.md#GetFarmById) | **Get** /farm/{farmId} | Find farm by id
[**RemovePerson**](FarmApi.md#RemovePerson) | **Put** /farm/removeFarmAccess/{farmId} | Remove farm access for an existing person
[**UpdateFarm**](FarmApi.md#UpdateFarm) | **Put** /farm | Update an existing farm



## AddFarm

> Farm AddFarm(ctx).Farm(farm).Execute()

Add a new farm to the db



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
    farm := *openapiclient.NewFarm() // Farm | Create a new farm in the db

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.AddFarm(context.Background()).Farm(farm).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.AddFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFarm`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.AddFarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **farm** | [**Farm**](Farm.md) | Create a new farm in the db | 

### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPerson

> Farm AddPerson(ctx, farmId).Farm(farm).Execute()

Add farm access for an existing person



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
    farmId := int64(789) // int64 | id of farm to return
    farm := *openapiclient.NewFarm() // Farm | Add farm access for an existent person

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.AddPerson(context.Background(), farmId).Farm(farm).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.AddPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPerson`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.AddPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **int64** | id of farm to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **farm** | [**Farm**](Farm.md) | Add farm access for an existent person | 

### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFarm

> DeleteFarm(ctx, farmId).ApiKey(apiKey).Execute()

Deletes a farm



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
    farmId := int64(789) // int64 | Farm id to delete
    apiKey := "apiKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.DeleteFarm(context.Background(), farmId).ApiKey(apiKey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.DeleteFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **int64** | Farm id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFarmRequest struct via the builder pattern


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


## FindFarmByName

> Farm FindFarmByName(ctx).Name(name).Execute()

Finds a farm by name



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
    name := "name_example" // string | String for farm name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.FindFarmByName(context.Background()).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.FindFarmByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindFarmByName`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.FindFarmByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindFarmByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | String for farm name | 

### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarmById

> Farm GetFarmById(ctx, farmId).Execute()

Find farm by id



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
    farmId := int64(789) // int64 | id of farm to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.GetFarmById(context.Background(), farmId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.GetFarmById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFarmById`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.GetFarmById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **int64** | id of farm to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePerson

> Farm RemovePerson(ctx, farmId).Farm(farm).Execute()

Remove farm access for an existing person



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
    farmId := int64(789) // int64 | id of farm to return
    farm := *openapiclient.NewFarm() // Farm | remove farm access for an existent person

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.RemovePerson(context.Background(), farmId).Farm(farm).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.RemovePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePerson`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.RemovePerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **int64** | id of farm to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **farm** | [**Farm**](Farm.md) | remove farm access for an existent person | 

### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFarm

> Farm UpdateFarm(ctx).Farm(farm).Execute()

Update an existing farm



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
    farm := *openapiclient.NewFarm() // Farm | Update an existent farm

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FarmApi.UpdateFarm(context.Background()).Farm(farm).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FarmApi.UpdateFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFarm`: Farm
    fmt.Fprintf(os.Stdout, "Response from `FarmApi.UpdateFarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **farm** | [**Farm**](Farm.md) | Update an existent farm | 

### Return type

[**Farm**](Farm.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

