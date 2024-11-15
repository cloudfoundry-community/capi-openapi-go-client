# Process

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DiskInMb** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**Instances** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**ProcessLinks**](ProcessLinks.md) |  | [optional] 
**LogRateLimitInBytesPerSecond** | Pointer to **int32** |  | [optional] 
**MemoryInMb** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**ReadinessHealthCheck** | Pointer to [**ReadinessHealthCheck**](ReadinessHealthCheck.md) |  | [optional] 
**Relationships** | Pointer to [**ProcessRelationships**](ProcessRelationships.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewProcess

`func NewProcess() *Process`

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

### HasCommand

`func (o *Process) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

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

### HasCreatedAt

`func (o *Process) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasDiskInMb

`func (o *Process) HasDiskInMb() bool`

HasDiskInMb returns a boolean if a field has been set.

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

### HasGuid

`func (o *Process) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHealthCheck

`func (o *Process) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *Process) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *Process) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *Process) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

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

### HasInstances

`func (o *Process) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

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

### HasLinks

`func (o *Process) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasLogRateLimitInBytesPerSecond

`func (o *Process) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

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

### HasMemoryInMb

`func (o *Process) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetMetadata

`func (o *Process) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Process) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Process) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Process) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReadinessHealthCheck

`func (o *Process) GetReadinessHealthCheck() ReadinessHealthCheck`

GetReadinessHealthCheck returns the ReadinessHealthCheck field if non-nil, zero value otherwise.

### GetReadinessHealthCheckOk

`func (o *Process) GetReadinessHealthCheckOk() (*ReadinessHealthCheck, bool)`

GetReadinessHealthCheckOk returns a tuple with the ReadinessHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessHealthCheck

`func (o *Process) SetReadinessHealthCheck(v ReadinessHealthCheck)`

SetReadinessHealthCheck sets ReadinessHealthCheck field to given value.

### HasReadinessHealthCheck

`func (o *Process) HasReadinessHealthCheck() bool`

HasReadinessHealthCheck returns a boolean if a field has been set.

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

### HasRelationships

`func (o *Process) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

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

### HasType

`func (o *Process) HasType() bool`

HasType returns a boolean if a field has been set.

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

### HasUpdatedAt

`func (o *Process) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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


