# V3ServiceBrokersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | [**Authentication**](Authentication.md) |  | 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Name** | **string** |  | 
**Relationships** | Pointer to [**V3ServiceBrokersPostRequestRelationships**](V3ServiceBrokersPostRequestRelationships.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewV3ServiceBrokersPostRequest

`func NewV3ServiceBrokersPostRequest(authentication Authentication, name string, url string, ) *V3ServiceBrokersPostRequest`

NewV3ServiceBrokersPostRequest instantiates a new V3ServiceBrokersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServiceBrokersPostRequestWithDefaults

`func NewV3ServiceBrokersPostRequestWithDefaults() *V3ServiceBrokersPostRequest`

NewV3ServiceBrokersPostRequestWithDefaults instantiates a new V3ServiceBrokersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *V3ServiceBrokersPostRequest) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *V3ServiceBrokersPostRequest) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *V3ServiceBrokersPostRequest) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.


### GetMetadata

`func (o *V3ServiceBrokersPostRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3ServiceBrokersPostRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3ServiceBrokersPostRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3ServiceBrokersPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3ServiceBrokersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3ServiceBrokersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3ServiceBrokersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *V3ServiceBrokersPostRequest) GetRelationships() V3ServiceBrokersPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3ServiceBrokersPostRequest) GetRelationshipsOk() (*V3ServiceBrokersPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3ServiceBrokersPostRequest) SetRelationships(v V3ServiceBrokersPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3ServiceBrokersPostRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetUrl

`func (o *V3ServiceBrokersPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V3ServiceBrokersPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V3ServiceBrokersPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


