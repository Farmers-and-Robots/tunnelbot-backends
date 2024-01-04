# Tunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewTunnel

`func NewTunnel(name string, ) *Tunnel`

NewTunnel instantiates a new Tunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelWithDefaults

`func NewTunnelWithDefaults() *Tunnel`

NewTunnelWithDefaults instantiates a new Tunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tunnel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tunnel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tunnel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Tunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Tunnel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Tunnel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Tunnel) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Tunnel) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Tunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tunnel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


