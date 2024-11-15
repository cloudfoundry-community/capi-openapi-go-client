# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DiskInMb** | Pointer to **int32** |  | [optional] 
**DropletGuid** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**TaskLinks**](TaskLinks.md) |  | [optional] 
**LogRateLimitInBytesPerSecond** | Pointer to **int32** |  | [optional] 
**MemoryInMb** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseRelationships**](V3AppsGuidDropletsCurrentGet200ResponseRelationships.md) |  | [optional] 
**Result** | Pointer to [**TaskResult**](TaskResult.md) |  | [optional] 
**SequenceId** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Task) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Task) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Task) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Task) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDiskInMb

`func (o *Task) GetDiskInMb() int32`

GetDiskInMb returns the DiskInMb field if non-nil, zero value otherwise.

### GetDiskInMbOk

`func (o *Task) GetDiskInMbOk() (*int32, bool)`

GetDiskInMbOk returns a tuple with the DiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInMb

`func (o *Task) SetDiskInMb(v int32)`

SetDiskInMb sets DiskInMb field to given value.

### HasDiskInMb

`func (o *Task) HasDiskInMb() bool`

HasDiskInMb returns a boolean if a field has been set.

### GetDropletGuid

`func (o *Task) GetDropletGuid() string`

GetDropletGuid returns the DropletGuid field if non-nil, zero value otherwise.

### GetDropletGuidOk

`func (o *Task) GetDropletGuidOk() (*string, bool)`

GetDropletGuidOk returns a tuple with the DropletGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletGuid

`func (o *Task) SetDropletGuid(v string)`

SetDropletGuid sets DropletGuid field to given value.

### HasDropletGuid

`func (o *Task) HasDropletGuid() bool`

HasDropletGuid returns a boolean if a field has been set.

### GetGuid

`func (o *Task) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Task) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Task) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Task) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *Task) GetLinks() TaskLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Task) GetLinksOk() (*TaskLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Task) SetLinks(v TaskLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Task) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLogRateLimitInBytesPerSecond

`func (o *Task) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *Task) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *Task) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.

### HasLogRateLimitInBytesPerSecond

`func (o *Task) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

### GetMemoryInMb

`func (o *Task) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *Task) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *Task) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *Task) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetMetadata

`func (o *Task) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Task) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Task) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Task) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *Task) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Task) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Task) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Task) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetResult

`func (o *Task) GetResult() TaskResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Task) GetResultOk() (*TaskResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Task) SetResult(v TaskResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *Task) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSequenceId

`func (o *Task) GetSequenceId() int32`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *Task) GetSequenceIdOk() (*int32, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *Task) SetSequenceId(v int32)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *Task) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetState

`func (o *Task) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Task) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Task) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Task) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


