# PersonAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonId** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 

## Methods

### NewPersonAssociation

`func NewPersonAssociation() *PersonAssociation`

NewPersonAssociation instantiates a new PersonAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonAssociationWithDefaults

`func NewPersonAssociationWithDefaults() *PersonAssociation`

NewPersonAssociationWithDefaults instantiates a new PersonAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonId

`func (o *PersonAssociation) GetPersonId() int64`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *PersonAssociation) GetPersonIdOk() (*int64, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *PersonAssociation) SetPersonId(v int64)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *PersonAssociation) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetDisplayName

`func (o *PersonAssociation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PersonAssociation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PersonAssociation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PersonAssociation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *PersonAssociation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PersonAssociation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PersonAssociation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PersonAssociation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *PersonAssociation) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PersonAssociation) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PersonAssociation) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *PersonAssociation) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


