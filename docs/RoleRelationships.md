# RoleRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**RoleRelationshipsOrganization**](RoleRelationshipsOrganization.md) |  | [optional] 
**Space** | Pointer to [**RoleRelationshipsSpace**](RoleRelationshipsSpace.md) |  | [optional] 
**User** | [**RoleRelationshipsUser**](RoleRelationshipsUser.md) |  | 

## Methods

### NewRoleRelationships

`func NewRoleRelationships(user RoleRelationshipsUser, ) *RoleRelationships`

NewRoleRelationships instantiates a new RoleRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRelationshipsWithDefaults

`func NewRoleRelationshipsWithDefaults() *RoleRelationships`

NewRoleRelationshipsWithDefaults instantiates a new RoleRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *RoleRelationships) GetOrganization() RoleRelationshipsOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RoleRelationships) GetOrganizationOk() (*RoleRelationshipsOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RoleRelationships) SetOrganization(v RoleRelationshipsOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RoleRelationships) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSpace

`func (o *RoleRelationships) GetSpace() RoleRelationshipsSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *RoleRelationships) GetSpaceOk() (*RoleRelationshipsSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *RoleRelationships) SetSpace(v RoleRelationshipsSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *RoleRelationships) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetUser

`func (o *RoleRelationships) GetUser() RoleRelationshipsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleRelationships) GetUserOk() (*RoleRelationshipsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleRelationships) SetUser(v RoleRelationshipsUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


