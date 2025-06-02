# RoleCreateRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**RoleCreateRelationshipsOrganization**](RoleCreateRelationshipsOrganization.md) |  | [optional] 
**Space** | Pointer to [**RoleCreateRelationshipsSpace**](RoleCreateRelationshipsSpace.md) |  | [optional] 
**User** | [**RoleCreateRelationshipsUser**](RoleCreateRelationshipsUser.md) |  | 

## Methods

### NewRoleCreateRelationships

`func NewRoleCreateRelationships(user RoleCreateRelationshipsUser, ) *RoleCreateRelationships`

NewRoleCreateRelationships instantiates a new RoleCreateRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateRelationshipsWithDefaults

`func NewRoleCreateRelationshipsWithDefaults() *RoleCreateRelationships`

NewRoleCreateRelationshipsWithDefaults instantiates a new RoleCreateRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *RoleCreateRelationships) GetOrganization() RoleCreateRelationshipsOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RoleCreateRelationships) GetOrganizationOk() (*RoleCreateRelationshipsOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RoleCreateRelationships) SetOrganization(v RoleCreateRelationshipsOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RoleCreateRelationships) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSpace

`func (o *RoleCreateRelationships) GetSpace() RoleCreateRelationshipsSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *RoleCreateRelationships) GetSpaceOk() (*RoleCreateRelationshipsSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *RoleCreateRelationships) SetSpace(v RoleCreateRelationshipsSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *RoleCreateRelationships) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetUser

`func (o *RoleCreateRelationships) GetUser() RoleCreateRelationshipsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleCreateRelationships) GetUserOk() (*RoleCreateRelationshipsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleCreateRelationships) SetUser(v RoleCreateRelationshipsUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


