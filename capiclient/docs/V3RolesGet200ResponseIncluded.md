# V3RolesGet200ResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Spaces** | Pointer to [**[]Space**](Space.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewV3RolesGet200ResponseIncluded

`func NewV3RolesGet200ResponseIncluded() *V3RolesGet200ResponseIncluded`

NewV3RolesGet200ResponseIncluded instantiates a new V3RolesGet200ResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3RolesGet200ResponseIncludedWithDefaults

`func NewV3RolesGet200ResponseIncludedWithDefaults() *V3RolesGet200ResponseIncluded`

NewV3RolesGet200ResponseIncludedWithDefaults instantiates a new V3RolesGet200ResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *V3RolesGet200ResponseIncluded) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *V3RolesGet200ResponseIncluded) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *V3RolesGet200ResponseIncluded) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *V3RolesGet200ResponseIncluded) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetSpaces

`func (o *V3RolesGet200ResponseIncluded) GetSpaces() []Space`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *V3RolesGet200ResponseIncluded) GetSpacesOk() (*[]Space, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *V3RolesGet200ResponseIncluded) SetSpaces(v []Space)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *V3RolesGet200ResponseIncluded) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.

### GetUsers

`func (o *V3RolesGet200ResponseIncluded) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V3RolesGet200ResponseIncluded) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V3RolesGet200ResponseIncluded) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V3RolesGet200ResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


