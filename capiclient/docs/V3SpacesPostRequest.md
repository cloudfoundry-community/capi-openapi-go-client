# V3SpacesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V3DropletsGuidPatchRequestMetadata**](V3DropletsGuidPatchRequestMetadata.md) |  | [optional] 
**Name** | **string** |  | 
**Relationships** | [**V3SpacesPostRequestRelationships**](V3SpacesPostRequestRelationships.md) |  | 

## Methods

### NewV3SpacesPostRequest

`func NewV3SpacesPostRequest(name string, relationships V3SpacesPostRequestRelationships, ) *V3SpacesPostRequest`

NewV3SpacesPostRequest instantiates a new V3SpacesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SpacesPostRequestWithDefaults

`func NewV3SpacesPostRequestWithDefaults() *V3SpacesPostRequest`

NewV3SpacesPostRequestWithDefaults instantiates a new V3SpacesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V3SpacesPostRequest) GetMetadata() V3DropletsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3SpacesPostRequest) GetMetadataOk() (*V3DropletsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3SpacesPostRequest) SetMetadata(v V3DropletsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3SpacesPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3SpacesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3SpacesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3SpacesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *V3SpacesPostRequest) GetRelationships() V3SpacesPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3SpacesPostRequest) GetRelationshipsOk() (*V3SpacesPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3SpacesPostRequest) SetRelationships(v V3SpacesPostRequestRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


