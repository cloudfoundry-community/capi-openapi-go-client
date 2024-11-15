# V3PackagesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**V3PackagesPostRequestData**](V3PackagesPostRequestData.md) |  | [optional] 
**Metadata** | Pointer to [**V3PackagesPostRequestMetadata**](V3PackagesPostRequestMetadata.md) |  | [optional] 
**Relationships** | Pointer to [**V3PackagesPostRequestRelationships**](V3PackagesPostRequestRelationships.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the package; valid values are bits, docker | [optional] 

## Methods

### NewV3PackagesPostRequest

`func NewV3PackagesPostRequest() *V3PackagesPostRequest`

NewV3PackagesPostRequest instantiates a new V3PackagesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PackagesPostRequestWithDefaults

`func NewV3PackagesPostRequestWithDefaults() *V3PackagesPostRequest`

NewV3PackagesPostRequestWithDefaults instantiates a new V3PackagesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V3PackagesPostRequest) GetData() V3PackagesPostRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V3PackagesPostRequest) GetDataOk() (*V3PackagesPostRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V3PackagesPostRequest) SetData(v V3PackagesPostRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *V3PackagesPostRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *V3PackagesPostRequest) GetMetadata() V3PackagesPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3PackagesPostRequest) GetMetadataOk() (*V3PackagesPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3PackagesPostRequest) SetMetadata(v V3PackagesPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3PackagesPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationships

`func (o *V3PackagesPostRequest) GetRelationships() V3PackagesPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3PackagesPostRequest) GetRelationshipsOk() (*V3PackagesPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3PackagesPostRequest) SetRelationships(v V3PackagesPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3PackagesPostRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *V3PackagesPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3PackagesPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3PackagesPostRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V3PackagesPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


