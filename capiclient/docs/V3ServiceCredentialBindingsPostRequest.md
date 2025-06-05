# V3ServiceCredentialBindingsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V3DropletsGuidPatchRequestMetadata**](V3DropletsGuidPatchRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | [**V3ServiceCredentialBindingsPostRequestRelationships**](V3ServiceCredentialBindingsPostRequestRelationships.md) |  | 
**Type** | **string** |  | 

## Methods

### NewV3ServiceCredentialBindingsPostRequest

`func NewV3ServiceCredentialBindingsPostRequest(relationships V3ServiceCredentialBindingsPostRequestRelationships, type_ string, ) *V3ServiceCredentialBindingsPostRequest`

NewV3ServiceCredentialBindingsPostRequest instantiates a new V3ServiceCredentialBindingsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServiceCredentialBindingsPostRequestWithDefaults

`func NewV3ServiceCredentialBindingsPostRequestWithDefaults() *V3ServiceCredentialBindingsPostRequest`

NewV3ServiceCredentialBindingsPostRequestWithDefaults instantiates a new V3ServiceCredentialBindingsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V3ServiceCredentialBindingsPostRequest) GetMetadata() V3DropletsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3ServiceCredentialBindingsPostRequest) GetMetadataOk() (*V3DropletsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3ServiceCredentialBindingsPostRequest) SetMetadata(v V3DropletsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3ServiceCredentialBindingsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3ServiceCredentialBindingsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3ServiceCredentialBindingsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3ServiceCredentialBindingsPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3ServiceCredentialBindingsPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *V3ServiceCredentialBindingsPostRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V3ServiceCredentialBindingsPostRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V3ServiceCredentialBindingsPostRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V3ServiceCredentialBindingsPostRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRelationships

`func (o *V3ServiceCredentialBindingsPostRequest) GetRelationships() V3ServiceCredentialBindingsPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3ServiceCredentialBindingsPostRequest) GetRelationshipsOk() (*V3ServiceCredentialBindingsPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3ServiceCredentialBindingsPostRequest) SetRelationships(v V3ServiceCredentialBindingsPostRequestRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *V3ServiceCredentialBindingsPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3ServiceCredentialBindingsPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3ServiceCredentialBindingsPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


