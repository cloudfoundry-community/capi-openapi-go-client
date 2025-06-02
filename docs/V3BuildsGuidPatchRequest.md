# V3BuildsGuidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**V3BuildsGuidPatchRequestCreatedBy**](V3BuildsGuidPatchRequestCreatedBy.md) |  | [optional] 
**Error** | Pointer to **string** | Error message if build failed | [optional] 
**Lifecycle** | Pointer to [**V3BuildsGuidPatchRequestLifecycle**](V3BuildsGuidPatchRequestLifecycle.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**State** | Pointer to **string** | New state for the build | [optional] 

## Methods

### NewV3BuildsGuidPatchRequest

`func NewV3BuildsGuidPatchRequest() *V3BuildsGuidPatchRequest`

NewV3BuildsGuidPatchRequest instantiates a new V3BuildsGuidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildsGuidPatchRequestWithDefaults

`func NewV3BuildsGuidPatchRequestWithDefaults() *V3BuildsGuidPatchRequest`

NewV3BuildsGuidPatchRequestWithDefaults instantiates a new V3BuildsGuidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *V3BuildsGuidPatchRequest) GetCreatedBy() V3BuildsGuidPatchRequestCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *V3BuildsGuidPatchRequest) GetCreatedByOk() (*V3BuildsGuidPatchRequestCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *V3BuildsGuidPatchRequest) SetCreatedBy(v V3BuildsGuidPatchRequestCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *V3BuildsGuidPatchRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetError

`func (o *V3BuildsGuidPatchRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V3BuildsGuidPatchRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V3BuildsGuidPatchRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V3BuildsGuidPatchRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLifecycle

`func (o *V3BuildsGuidPatchRequest) GetLifecycle() V3BuildsGuidPatchRequestLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3BuildsGuidPatchRequest) GetLifecycleOk() (*V3BuildsGuidPatchRequestLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3BuildsGuidPatchRequest) SetLifecycle(v V3BuildsGuidPatchRequestLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3BuildsGuidPatchRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMetadata

`func (o *V3BuildsGuidPatchRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3BuildsGuidPatchRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3BuildsGuidPatchRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3BuildsGuidPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetState

`func (o *V3BuildsGuidPatchRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V3BuildsGuidPatchRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V3BuildsGuidPatchRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V3BuildsGuidPatchRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


