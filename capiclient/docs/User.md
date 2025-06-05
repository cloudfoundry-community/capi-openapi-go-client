# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time with zone when the object was created | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the user | [optional] 
**Links** | Pointer to [**UserLinks**](UserLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3UsersPostRequestMetadata**](V3UsersPostRequestMetadata.md) |  | [optional] 
**Origin** | Pointer to **NullableString** | The identity provider for the UAA user; will be null for UAA clients | [optional] 
**PresentationName** | Pointer to **string** | The name displayed for the user; for UAA users, this is the same as the username. For UAA clients, this is the UAA client ID | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time with zone when the object was last updated | [optional] 
**Username** | Pointer to **NullableString** | The name registered in UAA; will be null for UAA clients and non-UAA users | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *User) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *User) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *User) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *User) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *User) GetLinks() UserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*UserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v UserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *User) GetMetadata() V3UsersPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *User) GetMetadataOk() (*V3UsersPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *User) SetMetadata(v V3UsersPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *User) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrigin

`func (o *User) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *User) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *User) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *User) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *User) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *User) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPresentationName

`func (o *User) GetPresentationName() string`

GetPresentationName returns the PresentationName field if non-nil, zero value otherwise.

### GetPresentationNameOk

`func (o *User) GetPresentationNameOk() (*string, bool)`

GetPresentationNameOk returns a tuple with the PresentationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentationName

`func (o *User) SetPresentationName(v string)`

SetPresentationName sets PresentationName field to given value.

### HasPresentationName

`func (o *User) HasPresentationName() bool`

HasPresentationName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *User) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *User) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


