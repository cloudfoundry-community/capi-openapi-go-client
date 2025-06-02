# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp when the role was created | 
**Guid** | **string** | Unique identifier for the role | 
**Links** | [**RoleLinks**](RoleLinks.md) |  | 
**Metadata** | Pointer to [**RoleMetadata**](RoleMetadata.md) |  | [optional] 
**Relationships** | [**RoleRelationships**](RoleRelationships.md) |  | 
**Type** | **string** | The type of role | 
**UpdatedAt** | **time.Time** | Timestamp when the role was last updated | 

## Methods

### NewRole

`func NewRole(createdAt time.Time, guid string, links RoleLinks, relationships RoleRelationships, type_ string, updatedAt time.Time, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Role) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Role) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Role) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGuid

`func (o *Role) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Role) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Role) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *Role) GetLinks() RoleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Role) GetLinksOk() (*RoleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Role) SetLinks(v RoleLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *Role) GetMetadata() RoleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Role) GetMetadataOk() (*RoleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Role) SetMetadata(v RoleMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Role) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationships

`func (o *Role) GetRelationships() RoleRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Role) GetRelationshipsOk() (*RoleRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Role) SetRelationships(v RoleRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *Role) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Role) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Role) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Role) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


