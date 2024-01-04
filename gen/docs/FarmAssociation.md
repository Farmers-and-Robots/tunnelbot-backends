# FarmAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FarmId** | Pointer to **int64** |  | [optional] 
**FarmName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 

## Methods

### NewFarmAssociation

`func NewFarmAssociation() *FarmAssociation`

NewFarmAssociation instantiates a new FarmAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmAssociationWithDefaults

`func NewFarmAssociationWithDefaults() *FarmAssociation`

NewFarmAssociationWithDefaults instantiates a new FarmAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFarmId

`func (o *FarmAssociation) GetFarmId() int64`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *FarmAssociation) GetFarmIdOk() (*int64, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *FarmAssociation) SetFarmId(v int64)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *FarmAssociation) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetFarmName

`func (o *FarmAssociation) GetFarmName() string`

GetFarmName returns the FarmName field if non-nil, zero value otherwise.

### GetFarmNameOk

`func (o *FarmAssociation) GetFarmNameOk() (*string, bool)`

GetFarmNameOk returns a tuple with the FarmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmName

`func (o *FarmAssociation) SetFarmName(v string)`

SetFarmName sets FarmName field to given value.

### HasFarmName

`func (o *FarmAssociation) HasFarmName() bool`

HasFarmName returns a boolean if a field has been set.

### GetRole

`func (o *FarmAssociation) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FarmAssociation) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FarmAssociation) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *FarmAssociation) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


