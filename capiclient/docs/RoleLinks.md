# RoleLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**NullableRoleLinksOrganization**](RoleLinksOrganization.md) |  | [optional] 
**Self** | [**RoleLinksSelf**](RoleLinksSelf.md) |  | 
**Space** | Pointer to [**NullableRoleLinksSpace**](RoleLinksSpace.md) |  | [optional] 
**User** | [**RoleLinksUser**](RoleLinksUser.md) |  | 

## Methods

### NewRoleLinks

`func NewRoleLinks(self RoleLinksSelf, user RoleLinksUser, ) *RoleLinks`

NewRoleLinks instantiates a new RoleLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleLinksWithDefaults

`func NewRoleLinksWithDefaults() *RoleLinks`

NewRoleLinksWithDefaults instantiates a new RoleLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *RoleLinks) GetOrganization() RoleLinksOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RoleLinks) GetOrganizationOk() (*RoleLinksOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RoleLinks) SetOrganization(v RoleLinksOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RoleLinks) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *RoleLinks) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *RoleLinks) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetSelf

`func (o *RoleLinks) GetSelf() RoleLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RoleLinks) GetSelfOk() (*RoleLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RoleLinks) SetSelf(v RoleLinksSelf)`

SetSelf sets Self field to given value.


### GetSpace

`func (o *RoleLinks) GetSpace() RoleLinksSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *RoleLinks) GetSpaceOk() (*RoleLinksSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *RoleLinks) SetSpace(v RoleLinksSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *RoleLinks) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### SetSpaceNil

`func (o *RoleLinks) SetSpaceNil(b bool)`

 SetSpaceNil sets the value for Space to be an explicit nil

### UnsetSpace
`func (o *RoleLinks) UnsetSpace()`

UnsetSpace ensures that no value is present for Space, not even an explicit nil
### GetUser

`func (o *RoleLinks) GetUser() RoleLinksUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleLinks) GetUserOk() (*RoleLinksUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleLinks) SetUser(v RoleLinksUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


