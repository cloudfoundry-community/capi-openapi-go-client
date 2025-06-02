# RoleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**RoleCreateMetadata**](RoleCreateMetadata.md) |  | [optional] 
**Relationships** | [**RoleCreateRelationships**](RoleCreateRelationships.md) |  | 
**Type** | **string** | The type of role to create | 

## Methods

### NewRoleCreate

`func NewRoleCreate(relationships RoleCreateRelationships, type_ string, ) *RoleCreate`

NewRoleCreate instantiates a new RoleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateWithDefaults

`func NewRoleCreateWithDefaults() *RoleCreate`

NewRoleCreateWithDefaults instantiates a new RoleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RoleCreate) GetMetadata() RoleCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleCreate) GetMetadataOk() (*RoleCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleCreate) SetMetadata(v RoleCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RoleCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationships

`func (o *RoleCreate) GetRelationships() RoleCreateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RoleCreate) GetRelationshipsOk() (*RoleCreateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RoleCreate) SetRelationships(v RoleCreateRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *RoleCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleCreate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


