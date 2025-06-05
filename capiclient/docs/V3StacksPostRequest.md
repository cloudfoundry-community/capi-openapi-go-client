# V3StacksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the stack; must be no longer than 250 characters | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | **string** | Name of the stack; must be unique and no longer than 250 characters | 

## Methods

### NewV3StacksPostRequest

`func NewV3StacksPostRequest(name string, ) *V3StacksPostRequest`

NewV3StacksPostRequest instantiates a new V3StacksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3StacksPostRequestWithDefaults

`func NewV3StacksPostRequestWithDefaults() *V3StacksPostRequest`

NewV3StacksPostRequestWithDefaults instantiates a new V3StacksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V3StacksPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V3StacksPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V3StacksPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V3StacksPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *V3StacksPostRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3StacksPostRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3StacksPostRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3StacksPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3StacksPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3StacksPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3StacksPostRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


