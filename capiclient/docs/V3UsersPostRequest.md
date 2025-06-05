# V3UsersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Unique identifier for the user. For UAA users this will match the user ID of an existing UAA user’s GUID; in the case of UAA clients, this will match the UAA client ID | [optional] 
**Metadata** | Pointer to [**V3UsersPostRequestMetadata**](V3UsersPostRequestMetadata.md) |  | [optional] 

## Methods

### NewV3UsersPostRequest

`func NewV3UsersPostRequest() *V3UsersPostRequest`

NewV3UsersPostRequest instantiates a new V3UsersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3UsersPostRequestWithDefaults

`func NewV3UsersPostRequestWithDefaults() *V3UsersPostRequest`

NewV3UsersPostRequestWithDefaults instantiates a new V3UsersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *V3UsersPostRequest) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *V3UsersPostRequest) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *V3UsersPostRequest) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *V3UsersPostRequest) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMetadata

`func (o *V3UsersPostRequest) GetMetadata() V3UsersPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3UsersPostRequest) GetMetadataOk() (*V3UsersPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3UsersPostRequest) SetMetadata(v V3UsersPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3UsersPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


