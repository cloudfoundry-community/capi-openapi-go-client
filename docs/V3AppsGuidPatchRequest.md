# V3AppsGuidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Lifecycle** | Pointer to [**V3AppsGuidPatchRequestLifecycle**](V3AppsGuidPatchRequestLifecycle.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidPatchRequestMetadata**](V3AppsGuidPatchRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | New name for the app | [optional] 
**State** | Pointer to **string** | Desired state of the app | [optional] 

## Methods

### NewV3AppsGuidPatchRequest

`func NewV3AppsGuidPatchRequest() *V3AppsGuidPatchRequest`

NewV3AppsGuidPatchRequest instantiates a new V3AppsGuidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidPatchRequestWithDefaults

`func NewV3AppsGuidPatchRequestWithDefaults() *V3AppsGuidPatchRequest`

NewV3AppsGuidPatchRequestWithDefaults instantiates a new V3AppsGuidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *V3AppsGuidPatchRequest) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *V3AppsGuidPatchRequest) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *V3AppsGuidPatchRequest) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *V3AppsGuidPatchRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLifecycle

`func (o *V3AppsGuidPatchRequest) GetLifecycle() V3AppsGuidPatchRequestLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3AppsGuidPatchRequest) GetLifecycleOk() (*V3AppsGuidPatchRequestLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3AppsGuidPatchRequest) SetLifecycle(v V3AppsGuidPatchRequestLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3AppsGuidPatchRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMetadata

`func (o *V3AppsGuidPatchRequest) GetMetadata() V3AppsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3AppsGuidPatchRequest) GetMetadataOk() (*V3AppsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3AppsGuidPatchRequest) SetMetadata(v V3AppsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3AppsGuidPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3AppsGuidPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3AppsGuidPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3AppsGuidPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3AppsGuidPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *V3AppsGuidPatchRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V3AppsGuidPatchRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V3AppsGuidPatchRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V3AppsGuidPatchRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


