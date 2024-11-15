# V3ServiceRouteBindingsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | [**ServiceRouteBindingRelationships**](ServiceRouteBindingRelationships.md) |  | 

## Methods

### NewV3ServiceRouteBindingsPostRequest

`func NewV3ServiceRouteBindingsPostRequest(relationships ServiceRouteBindingRelationships, ) *V3ServiceRouteBindingsPostRequest`

NewV3ServiceRouteBindingsPostRequest instantiates a new V3ServiceRouteBindingsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServiceRouteBindingsPostRequestWithDefaults

`func NewV3ServiceRouteBindingsPostRequestWithDefaults() *V3ServiceRouteBindingsPostRequest`

NewV3ServiceRouteBindingsPostRequestWithDefaults instantiates a new V3ServiceRouteBindingsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V3ServiceRouteBindingsPostRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3ServiceRouteBindingsPostRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3ServiceRouteBindingsPostRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3ServiceRouteBindingsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParameters

`func (o *V3ServiceRouteBindingsPostRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V3ServiceRouteBindingsPostRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V3ServiceRouteBindingsPostRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V3ServiceRouteBindingsPostRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRelationships

`func (o *V3ServiceRouteBindingsPostRequest) GetRelationships() ServiceRouteBindingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3ServiceRouteBindingsPostRequest) GetRelationshipsOk() (*ServiceRouteBindingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3ServiceRouteBindingsPostRequest) SetRelationships(v ServiceRouteBindingRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


