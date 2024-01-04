# \PeopleApi

All URIs are relative to *https://tunnelbot.swagger.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeople**](PeopleApi.md#GetPeople) | **Get** /people | Return the people associated with a farm
[**GetPersonById**](PeopleApi.md#GetPersonById) | **Get** /people/{personId} | Find person by id



## GetPeople

> []Person GetPeople(ctx).Execute()

Return the people associated with a farm



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
    resp, r, err := api_client.PeopleApi.GetPeople(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GetPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeople`: []Person
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GetPeople`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeopleRequest struct via the builder pattern


### Return type

[**[]Person**](Person.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonById

> Person GetPersonById(ctx, personId).Execute()

Find person by id



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
    personId := int64(789) // int64 | id of person to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GetPersonById(context.Background(), personId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GetPersonById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonById`: Person
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GetPersonById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **int64** | id of person to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Person**](Person.md)

### Authorization

[firebase](../README.md#firebase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

