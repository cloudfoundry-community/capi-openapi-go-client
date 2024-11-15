# V3PackagesGuidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V3PackagesPostRequestMetadata**](V3PackagesPostRequestMetadata.md) |  | [optional] 
**Password** | Pointer to **string** | The password for the image&#39;s registry. Only possible for Docker package. | [optional] 
**Username** | Pointer to **string** | The username for the image&#39;s registry. Only possible for Docker package. | [optional] 

## Methods

### NewV3PackagesGuidPatchRequest

`func NewV3PackagesGuidPatchRequest() *V3PackagesGuidPatchRequest`

NewV3PackagesGuidPatchRequest instantiates a new V3PackagesGuidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PackagesGuidPatchRequestWithDefaults

`func NewV3PackagesGuidPatchRequestWithDefaults() *V3PackagesGuidPatchRequest`

NewV3PackagesGuidPatchRequestWithDefaults instantiates a new V3PackagesGuidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V3PackagesGuidPatchRequest) GetMetadata() V3PackagesPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3PackagesGuidPatchRequest) GetMetadataOk() (*V3PackagesPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3PackagesGuidPatchRequest) SetMetadata(v V3PackagesPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3PackagesGuidPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *V3PackagesGuidPatchRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V3PackagesGuidPatchRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V3PackagesGuidPatchRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V3PackagesGuidPatchRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *V3PackagesGuidPatchRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V3PackagesGuidPatchRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V3PackagesGuidPatchRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V3PackagesGuidPatchRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


