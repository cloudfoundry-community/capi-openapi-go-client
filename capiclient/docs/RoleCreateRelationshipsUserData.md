# RoleCreateRelationshipsUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** | User GUID | 
**Origin** | Pointer to **string** | Identity provider origin | [optional] [default to "uaa"]
**Username** | **string** | Username (email) | 

## Methods

### NewRoleCreateRelationshipsUserData

`func NewRoleCreateRelationshipsUserData(guid string, username string, ) *RoleCreateRelationshipsUserData`

NewRoleCreateRelationshipsUserData instantiates a new RoleCreateRelationshipsUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateRelationshipsUserDataWithDefaults

`func NewRoleCreateRelationshipsUserDataWithDefaults() *RoleCreateRelationshipsUserData`

NewRoleCreateRelationshipsUserDataWithDefaults instantiates a new RoleCreateRelationshipsUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *RoleCreateRelationshipsUserData) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *RoleCreateRelationshipsUserData) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *RoleCreateRelationshipsUserData) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetOrigin

`func (o *RoleCreateRelationshipsUserData) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *RoleCreateRelationshipsUserData) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *RoleCreateRelationshipsUserData) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *RoleCreateRelationshipsUserData) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetUsername

`func (o *RoleCreateRelationshipsUserData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RoleCreateRelationshipsUserData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RoleCreateRelationshipsUserData) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


