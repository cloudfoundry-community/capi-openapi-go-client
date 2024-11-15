# ProcessStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **NullableString** |  | [optional] 
**DiskQuota** | Pointer to **NullableInt32** |  | [optional] 
**FdsQuota** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**InstanceInternalIp** | Pointer to **string** |  | [optional] 
**InstancePorts** | Pointer to [**[]ProcessStatsInstancePortsInner**](ProcessStatsInstancePortsInner.md) |  | [optional] 
**IsolationSegment** | Pointer to **NullableString** |  | [optional] 
**LogRateLimit** | Pointer to **NullableInt32** |  | [optional] 
**MemQuota** | Pointer to **NullableInt32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to [**ProcessStatsUsage**](ProcessStatsUsage.md) |  | [optional] 

## Methods

### NewProcessStats

`func NewProcessStats() *ProcessStats`

NewProcessStats instantiates a new ProcessStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessStatsWithDefaults

`func NewProcessStatsWithDefaults() *ProcessStats`

NewProcessStatsWithDefaults instantiates a new ProcessStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ProcessStats) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProcessStats) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProcessStats) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProcessStats) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ProcessStats) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ProcessStats) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetDiskQuota

`func (o *ProcessStats) GetDiskQuota() int32`

GetDiskQuota returns the DiskQuota field if non-nil, zero value otherwise.

### GetDiskQuotaOk

`func (o *ProcessStats) GetDiskQuotaOk() (*int32, bool)`

GetDiskQuotaOk returns a tuple with the DiskQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskQuota

`func (o *ProcessStats) SetDiskQuota(v int32)`

SetDiskQuota sets DiskQuota field to given value.

### HasDiskQuota

`func (o *ProcessStats) HasDiskQuota() bool`

HasDiskQuota returns a boolean if a field has been set.

### SetDiskQuotaNil

`func (o *ProcessStats) SetDiskQuotaNil(b bool)`

 SetDiskQuotaNil sets the value for DiskQuota to be an explicit nil

### UnsetDiskQuota
`func (o *ProcessStats) UnsetDiskQuota()`

UnsetDiskQuota ensures that no value is present for DiskQuota, not even an explicit nil
### GetFdsQuota

`func (o *ProcessStats) GetFdsQuota() int32`

GetFdsQuota returns the FdsQuota field if non-nil, zero value otherwise.

### GetFdsQuotaOk

`func (o *ProcessStats) GetFdsQuotaOk() (*int32, bool)`

GetFdsQuotaOk returns a tuple with the FdsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdsQuota

`func (o *ProcessStats) SetFdsQuota(v int32)`

SetFdsQuota sets FdsQuota field to given value.

### HasFdsQuota

`func (o *ProcessStats) HasFdsQuota() bool`

HasFdsQuota returns a boolean if a field has been set.

### GetHost

`func (o *ProcessStats) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProcessStats) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProcessStats) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProcessStats) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIndex

`func (o *ProcessStats) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ProcessStats) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ProcessStats) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ProcessStats) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInstanceInternalIp

`func (o *ProcessStats) GetInstanceInternalIp() string`

GetInstanceInternalIp returns the InstanceInternalIp field if non-nil, zero value otherwise.

### GetInstanceInternalIpOk

`func (o *ProcessStats) GetInstanceInternalIpOk() (*string, bool)`

GetInstanceInternalIpOk returns a tuple with the InstanceInternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceInternalIp

`func (o *ProcessStats) SetInstanceInternalIp(v string)`

SetInstanceInternalIp sets InstanceInternalIp field to given value.

### HasInstanceInternalIp

`func (o *ProcessStats) HasInstanceInternalIp() bool`

HasInstanceInternalIp returns a boolean if a field has been set.

### GetInstancePorts

`func (o *ProcessStats) GetInstancePorts() []ProcessStatsInstancePortsInner`

GetInstancePorts returns the InstancePorts field if non-nil, zero value otherwise.

### GetInstancePortsOk

`func (o *ProcessStats) GetInstancePortsOk() (*[]ProcessStatsInstancePortsInner, bool)`

GetInstancePortsOk returns a tuple with the InstancePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePorts

`func (o *ProcessStats) SetInstancePorts(v []ProcessStatsInstancePortsInner)`

SetInstancePorts sets InstancePorts field to given value.

### HasInstancePorts

`func (o *ProcessStats) HasInstancePorts() bool`

HasInstancePorts returns a boolean if a field has been set.

### GetIsolationSegment

`func (o *ProcessStats) GetIsolationSegment() string`

GetIsolationSegment returns the IsolationSegment field if non-nil, zero value otherwise.

### GetIsolationSegmentOk

`func (o *ProcessStats) GetIsolationSegmentOk() (*string, bool)`

GetIsolationSegmentOk returns a tuple with the IsolationSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationSegment

`func (o *ProcessStats) SetIsolationSegment(v string)`

SetIsolationSegment sets IsolationSegment field to given value.

### HasIsolationSegment

`func (o *ProcessStats) HasIsolationSegment() bool`

HasIsolationSegment returns a boolean if a field has been set.

### SetIsolationSegmentNil

`func (o *ProcessStats) SetIsolationSegmentNil(b bool)`

 SetIsolationSegmentNil sets the value for IsolationSegment to be an explicit nil

### UnsetIsolationSegment
`func (o *ProcessStats) UnsetIsolationSegment()`

UnsetIsolationSegment ensures that no value is present for IsolationSegment, not even an explicit nil
### GetLogRateLimit

`func (o *ProcessStats) GetLogRateLimit() int32`

GetLogRateLimit returns the LogRateLimit field if non-nil, zero value otherwise.

### GetLogRateLimitOk

`func (o *ProcessStats) GetLogRateLimitOk() (*int32, bool)`

GetLogRateLimitOk returns a tuple with the LogRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimit

`func (o *ProcessStats) SetLogRateLimit(v int32)`

SetLogRateLimit sets LogRateLimit field to given value.

### HasLogRateLimit

`func (o *ProcessStats) HasLogRateLimit() bool`

HasLogRateLimit returns a boolean if a field has been set.

### SetLogRateLimitNil

`func (o *ProcessStats) SetLogRateLimitNil(b bool)`

 SetLogRateLimitNil sets the value for LogRateLimit to be an explicit nil

### UnsetLogRateLimit
`func (o *ProcessStats) UnsetLogRateLimit()`

UnsetLogRateLimit ensures that no value is present for LogRateLimit, not even an explicit nil
### GetMemQuota

`func (o *ProcessStats) GetMemQuota() int32`

GetMemQuota returns the MemQuota field if non-nil, zero value otherwise.

### GetMemQuotaOk

`func (o *ProcessStats) GetMemQuotaOk() (*int32, bool)`

GetMemQuotaOk returns a tuple with the MemQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemQuota

`func (o *ProcessStats) SetMemQuota(v int32)`

SetMemQuota sets MemQuota field to given value.

### HasMemQuota

`func (o *ProcessStats) HasMemQuota() bool`

HasMemQuota returns a boolean if a field has been set.

### SetMemQuotaNil

`func (o *ProcessStats) SetMemQuotaNil(b bool)`

 SetMemQuotaNil sets the value for MemQuota to be an explicit nil

### UnsetMemQuota
`func (o *ProcessStats) UnsetMemQuota()`

UnsetMemQuota ensures that no value is present for MemQuota, not even an explicit nil
### GetState

`func (o *ProcessStats) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProcessStats) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProcessStats) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProcessStats) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *ProcessStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProcessStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *ProcessStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ProcessStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ProcessStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ProcessStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsage

`func (o *ProcessStats) GetUsage() ProcessStatsUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ProcessStats) GetUsageOk() (*ProcessStatsUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ProcessStats) SetUsage(v ProcessStatsUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ProcessStats) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


