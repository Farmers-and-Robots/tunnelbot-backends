# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**FarmAssociations** | Pointer to [**[]FarmAssociation**](FarmAssociation.md) |  | [optional] 

## Methods

### NewPerson

`func NewPerson() *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Person) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Person) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Person) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Person) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Person) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Person) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Person) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Person) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDisplayName

`func (o *Person) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Person) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Person) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Person) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFirstName

`func (o *Person) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Person) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Person) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Person) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Person) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Person) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Person) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Person) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Person) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Person) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Person) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Person) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *Person) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Person) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Person) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Person) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFarmAssociations

`func (o *Person) GetFarmAssociations() []FarmAssociation`

GetFarmAssociations returns the FarmAssociations field if non-nil, zero value otherwise.

### GetFarmAssociationsOk

`func (o *Person) GetFarmAssociationsOk() (*[]FarmAssociation, bool)`

GetFarmAssociationsOk returns a tuple with the FarmAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmAssociations

`func (o *Person) SetFarmAssociations(v []FarmAssociation)`

SetFarmAssociations sets FarmAssociations field to given value.

### HasFarmAssociations

`func (o *Person) HasFarmAssociations() bool`

HasFarmAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


