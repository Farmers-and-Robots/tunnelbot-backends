# Farm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Owners** | Pointer to [**[]PersonAssociation**](PersonAssociation.md) |  | [optional] 
**People** | Pointer to [**[]PersonAssociation**](PersonAssociation.md) |  | [optional] 

## Methods

### NewFarm

`func NewFarm() *Farm`

NewFarm instantiates a new Farm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmWithDefaults

`func NewFarmWithDefaults() *Farm`

NewFarmWithDefaults instantiates a new Farm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Farm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Farm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Farm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Farm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Farm) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Farm) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Farm) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Farm) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Farm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Farm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Farm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Farm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateDate

`func (o *Farm) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Farm) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Farm) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Farm) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetAddress

`func (o *Farm) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Farm) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Farm) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Farm) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetOwners

`func (o *Farm) GetOwners() []PersonAssociation`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Farm) GetOwnersOk() (*[]PersonAssociation, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Farm) SetOwners(v []PersonAssociation)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Farm) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetPeople

`func (o *Farm) GetPeople() []PersonAssociation`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *Farm) GetPeopleOk() (*[]PersonAssociation, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *Farm) SetPeople(v []PersonAssociation)`

SetPeople sets People field to given value.

### HasPeople

`func (o *Farm) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


