# V3BuildsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to **map[string]string** | Environment variables for staging | [optional] 
**Lifecycle** | Pointer to [**V3BuildsPostRequestLifecycle**](V3BuildsPostRequestLifecycle.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Package** | [**V3BuildsPostRequestPackage**](V3BuildsPostRequestPackage.md) |  | 
**StagingDiskInMb** | Pointer to **int32** | Disk limit for staging container | [optional] 
**StagingLogRateLimitBytesPerSecond** | Pointer to **int32** | Log rate limit during staging | [optional] 
**StagingMemoryInMb** | Pointer to **int32** | Memory limit for staging container | [optional] 

## Methods

### NewV3BuildsPostRequest

`func NewV3BuildsPostRequest(package_ V3BuildsPostRequestPackage, ) *V3BuildsPostRequest`

NewV3BuildsPostRequest instantiates a new V3BuildsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildsPostRequestWithDefaults

`func NewV3BuildsPostRequestWithDefaults() *V3BuildsPostRequest`

NewV3BuildsPostRequestWithDefaults instantiates a new V3BuildsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *V3BuildsPostRequest) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *V3BuildsPostRequest) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *V3BuildsPostRequest) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *V3BuildsPostRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLifecycle

`func (o *V3BuildsPostRequest) GetLifecycle() V3BuildsPostRequestLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3BuildsPostRequest) GetLifecycleOk() (*V3BuildsPostRequestLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3BuildsPostRequest) SetLifecycle(v V3BuildsPostRequestLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3BuildsPostRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMetadata

`func (o *V3BuildsPostRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3BuildsPostRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3BuildsPostRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3BuildsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPackage

`func (o *V3BuildsPostRequest) GetPackage() V3BuildsPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *V3BuildsPostRequest) GetPackageOk() (*V3BuildsPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *V3BuildsPostRequest) SetPackage(v V3BuildsPostRequestPackage)`

SetPackage sets Package field to given value.


### GetStagingDiskInMb

`func (o *V3BuildsPostRequest) GetStagingDiskInMb() int32`

GetStagingDiskInMb returns the StagingDiskInMb field if non-nil, zero value otherwise.

### GetStagingDiskInMbOk

`func (o *V3BuildsPostRequest) GetStagingDiskInMbOk() (*int32, bool)`

GetStagingDiskInMbOk returns a tuple with the StagingDiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDiskInMb

`func (o *V3BuildsPostRequest) SetStagingDiskInMb(v int32)`

SetStagingDiskInMb sets StagingDiskInMb field to given value.

### HasStagingDiskInMb

`func (o *V3BuildsPostRequest) HasStagingDiskInMb() bool`

HasStagingDiskInMb returns a boolean if a field has been set.

### GetStagingLogRateLimitBytesPerSecond

`func (o *V3BuildsPostRequest) GetStagingLogRateLimitBytesPerSecond() int32`

GetStagingLogRateLimitBytesPerSecond returns the StagingLogRateLimitBytesPerSecond field if non-nil, zero value otherwise.

### GetStagingLogRateLimitBytesPerSecondOk

`func (o *V3BuildsPostRequest) GetStagingLogRateLimitBytesPerSecondOk() (*int32, bool)`

GetStagingLogRateLimitBytesPerSecondOk returns a tuple with the StagingLogRateLimitBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingLogRateLimitBytesPerSecond

`func (o *V3BuildsPostRequest) SetStagingLogRateLimitBytesPerSecond(v int32)`

SetStagingLogRateLimitBytesPerSecond sets StagingLogRateLimitBytesPerSecond field to given value.

### HasStagingLogRateLimitBytesPerSecond

`func (o *V3BuildsPostRequest) HasStagingLogRateLimitBytesPerSecond() bool`

HasStagingLogRateLimitBytesPerSecond returns a boolean if a field has been set.

### GetStagingMemoryInMb

`func (o *V3BuildsPostRequest) GetStagingMemoryInMb() int32`

GetStagingMemoryInMb returns the StagingMemoryInMb field if non-nil, zero value otherwise.

### GetStagingMemoryInMbOk

`func (o *V3BuildsPostRequest) GetStagingMemoryInMbOk() (*int32, bool)`

GetStagingMemoryInMbOk returns a tuple with the StagingMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMemoryInMb

`func (o *V3BuildsPostRequest) SetStagingMemoryInMb(v int32)`

SetStagingMemoryInMb sets StagingMemoryInMb field to given value.

### HasStagingMemoryInMb

`func (o *V3BuildsPostRequest) HasStagingMemoryInMb() bool`

HasStagingMemoryInMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


