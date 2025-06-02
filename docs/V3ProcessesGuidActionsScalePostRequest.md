# V3ProcessesGuidActionsScalePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskInMb** | Pointer to **int32** | Disk limit per instance in MB | [optional] 
**Instances** | Pointer to **int32** | Number of instances | [optional] 
**LogRateLimitInBytesPerSecond** | Pointer to **int32** | Log rate limit per instance (-1 for unlimited) | [optional] 
**MemoryInMb** | Pointer to **int32** | Memory limit per instance in MB | [optional] 

## Methods

### NewV3ProcessesGuidActionsScalePostRequest

`func NewV3ProcessesGuidActionsScalePostRequest() *V3ProcessesGuidActionsScalePostRequest`

NewV3ProcessesGuidActionsScalePostRequest instantiates a new V3ProcessesGuidActionsScalePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ProcessesGuidActionsScalePostRequestWithDefaults

`func NewV3ProcessesGuidActionsScalePostRequestWithDefaults() *V3ProcessesGuidActionsScalePostRequest`

NewV3ProcessesGuidActionsScalePostRequestWithDefaults instantiates a new V3ProcessesGuidActionsScalePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) GetDiskInMb() int32`

GetDiskInMb returns the DiskInMb field if non-nil, zero value otherwise.

### GetDiskInMbOk

`func (o *V3ProcessesGuidActionsScalePostRequest) GetDiskInMbOk() (*int32, bool)`

GetDiskInMbOk returns a tuple with the DiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) SetDiskInMb(v int32)`

SetDiskInMb sets DiskInMb field to given value.

### HasDiskInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) HasDiskInMb() bool`

HasDiskInMb returns a boolean if a field has been set.

### GetInstances

`func (o *V3ProcessesGuidActionsScalePostRequest) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V3ProcessesGuidActionsScalePostRequest) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V3ProcessesGuidActionsScalePostRequest) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V3ProcessesGuidActionsScalePostRequest) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLogRateLimitInBytesPerSecond

`func (o *V3ProcessesGuidActionsScalePostRequest) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *V3ProcessesGuidActionsScalePostRequest) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *V3ProcessesGuidActionsScalePostRequest) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.

### HasLogRateLimitInBytesPerSecond

`func (o *V3ProcessesGuidActionsScalePostRequest) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

### GetMemoryInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *V3ProcessesGuidActionsScalePostRequest) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *V3ProcessesGuidActionsScalePostRequest) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


