# V3RolesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to [**V3RolesPostRequestRelationships**](V3RolesPostRequestRelationships.md) |  | [optional] 
**Type** | Pointer to **string** | Role to create, see valid role types | [optional] 

## Methods

### NewV3RolesPostRequest

`func NewV3RolesPostRequest() *V3RolesPostRequest`

NewV3RolesPostRequest instantiates a new V3RolesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3RolesPostRequestWithDefaults

`func NewV3RolesPostRequestWithDefaults() *V3RolesPostRequest`

NewV3RolesPostRequestWithDefaults instantiates a new V3RolesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *V3RolesPostRequest) GetRelationships() V3RolesPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3RolesPostRequest) GetRelationshipsOk() (*V3RolesPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3RolesPostRequest) SetRelationships(v V3RolesPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3RolesPostRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *V3RolesPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3RolesPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3RolesPostRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V3RolesPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


