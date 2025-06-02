# Process

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **NullableString** | Start command for the process | 
**CreatedAt** | **time.Time** |  | 
**DiskInMb** | **int32** | Disk allocated per instance in MB | 
**Guid** | **string** | Unique identifier for the process | 
**HealthCheck** | [**ProcessHealthCheck**](ProcessHealthCheck.md) |  | 
**Instances** | **int32** | Number of instances | 
**Links** | [**ProcessLinks**](ProcessLinks.md) |  | 
**LogRateLimitInBytesPerSecond** | **int32** | Log rate limit per instance (-1 for unlimited) | 
**MemoryInMb** | **int32** | Memory allocated per instance in MB | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**ReadinessHealthCheck** | [**NullableProcessReadinessHealthCheck**](ProcessReadinessHealthCheck.md) |  | 
**Relationships** | [**ProcessRelationships**](ProcessRelationships.md) |  | 
**Type** | **string** | Process type (e.g., web, worker) | 
**UpdatedAt** | **time.Time** |  | 
**Version** | Pointer to **string** | Process version identifier | [optional] 

## Methods

### NewProcess

`func NewProcess(command NullableString, createdAt time.Time, diskInMb int32, guid string, healthCheck ProcessHealthCheck, instances int32, links ProcessLinks, logRateLimitInBytesPerSecond int32, memoryInMb int32, metadata V3AppsGuidTasksPostRequestMetadata, readinessHealthCheck NullableProcessReadinessHealthCheck, relationships ProcessRelationships, type_ string, updatedAt time.Time, ) *Process`

NewProcess instantiates a new Process object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessWithDefaults

`func NewProcessWithDefaults() *Process`

NewProcessWithDefaults instantiates a new Process object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Process) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Process) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Process) SetCommand(v string)`

SetCommand sets Command field to given value.


### SetCommandNil

`func (o *Process) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *Process) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetCreatedAt

`func (o *Process) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Process) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Process) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDiskInMb

`func (o *Process) GetDiskInMb() int32`

GetDiskInMb returns the DiskInMb field if non-nil, zero value otherwise.

### GetDiskInMbOk

`func (o *Process) GetDiskInMbOk() (*int32, bool)`

GetDiskInMbOk returns a tuple with the DiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInMb

`func (o *Process) SetDiskInMb(v int32)`

SetDiskInMb sets DiskInMb field to given value.


### GetGuid

`func (o *Process) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Process) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Process) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetHealthCheck

`func (o *Process) GetHealthCheck() ProcessHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *Process) GetHealthCheckOk() (*ProcessHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *Process) SetHealthCheck(v ProcessHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.


### GetInstances

`func (o *Process) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Process) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Process) SetInstances(v int32)`

SetInstances sets Instances field to given value.


### GetLinks

`func (o *Process) GetLinks() ProcessLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Process) GetLinksOk() (*ProcessLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Process) SetLinks(v ProcessLinks)`

SetLinks sets Links field to given value.


### GetLogRateLimitInBytesPerSecond

`func (o *Process) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *Process) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *Process) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.


### GetMemoryInMb

`func (o *Process) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *Process) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *Process) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.


### GetMetadata

`func (o *Process) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Process) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Process) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.


### GetReadinessHealthCheck

`func (o *Process) GetReadinessHealthCheck() ProcessReadinessHealthCheck`

GetReadinessHealthCheck returns the ReadinessHealthCheck field if non-nil, zero value otherwise.

### GetReadinessHealthCheckOk

`func (o *Process) GetReadinessHealthCheckOk() (*ProcessReadinessHealthCheck, bool)`

GetReadinessHealthCheckOk returns a tuple with the ReadinessHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessHealthCheck

`func (o *Process) SetReadinessHealthCheck(v ProcessReadinessHealthCheck)`

SetReadinessHealthCheck sets ReadinessHealthCheck field to given value.


### SetReadinessHealthCheckNil

`func (o *Process) SetReadinessHealthCheckNil(b bool)`

 SetReadinessHealthCheckNil sets the value for ReadinessHealthCheck to be an explicit nil

### UnsetReadinessHealthCheck
`func (o *Process) UnsetReadinessHealthCheck()`

UnsetReadinessHealthCheck ensures that no value is present for ReadinessHealthCheck, not even an explicit nil
### GetRelationships

`func (o *Process) GetRelationships() ProcessRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Process) GetRelationshipsOk() (*ProcessRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Process) SetRelationships(v ProcessRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *Process) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Process) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Process) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Process) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Process) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Process) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Process) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Process) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Process) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Process) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


