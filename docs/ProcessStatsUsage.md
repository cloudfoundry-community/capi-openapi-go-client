# ProcessStatsUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float32** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**LogRate** | Pointer to **int32** |  | [optional] 
**Mem** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProcessStatsUsage

`func NewProcessStatsUsage() *ProcessStatsUsage`

NewProcessStatsUsage instantiates a new ProcessStatsUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessStatsUsageWithDefaults

`func NewProcessStatsUsageWithDefaults() *ProcessStatsUsage`

NewProcessStatsUsageWithDefaults instantiates a new ProcessStatsUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ProcessStatsUsage) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProcessStatsUsage) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProcessStatsUsage) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProcessStatsUsage) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *ProcessStatsUsage) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ProcessStatsUsage) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ProcessStatsUsage) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ProcessStatsUsage) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetLogRate

`func (o *ProcessStatsUsage) GetLogRate() int32`

GetLogRate returns the LogRate field if non-nil, zero value otherwise.

### GetLogRateOk

`func (o *ProcessStatsUsage) GetLogRateOk() (*int32, bool)`

GetLogRateOk returns a tuple with the LogRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRate

`func (o *ProcessStatsUsage) SetLogRate(v int32)`

SetLogRate sets LogRate field to given value.

### HasLogRate

`func (o *ProcessStatsUsage) HasLogRate() bool`

HasLogRate returns a boolean if a field has been set.

### GetMem

`func (o *ProcessStatsUsage) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *ProcessStatsUsage) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *ProcessStatsUsage) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *ProcessStatsUsage) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetTime

`func (o *ProcessStatsUsage) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ProcessStatsUsage) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ProcessStatsUsage) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ProcessStatsUsage) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


