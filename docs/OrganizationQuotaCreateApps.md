# OrganizationQuotaCreateApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogRateLimitInBytesPerSecond** | Pointer to **NullableInt32** | Log rate limit per second (null &#x3D; unlimited) | [optional] 
**PerAppTasks** | Pointer to **NullableInt32** | Maximum tasks per app (null &#x3D; unlimited) | [optional] 
**PerProcessMemoryInMb** | Pointer to **NullableInt32** | Memory limit per app process (null &#x3D; unlimited) | [optional] 
**TotalInstances** | Pointer to **NullableInt32** | Total app instances allowed (null &#x3D; unlimited) | [optional] 
**TotalMemoryInMb** | Pointer to **NullableInt32** | Total memory limit for all apps (null &#x3D; unlimited) | [optional] 

## Methods

### NewOrganizationQuotaCreateApps

`func NewOrganizationQuotaCreateApps() *OrganizationQuotaCreateApps`

NewOrganizationQuotaCreateApps instantiates a new OrganizationQuotaCreateApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaCreateAppsWithDefaults

`func NewOrganizationQuotaCreateAppsWithDefaults() *OrganizationQuotaCreateApps`

NewOrganizationQuotaCreateAppsWithDefaults instantiates a new OrganizationQuotaCreateApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogRateLimitInBytesPerSecond

`func (o *OrganizationQuotaCreateApps) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *OrganizationQuotaCreateApps) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *OrganizationQuotaCreateApps) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.

### HasLogRateLimitInBytesPerSecond

`func (o *OrganizationQuotaCreateApps) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

### SetLogRateLimitInBytesPerSecondNil

`func (o *OrganizationQuotaCreateApps) SetLogRateLimitInBytesPerSecondNil(b bool)`

 SetLogRateLimitInBytesPerSecondNil sets the value for LogRateLimitInBytesPerSecond to be an explicit nil

### UnsetLogRateLimitInBytesPerSecond
`func (o *OrganizationQuotaCreateApps) UnsetLogRateLimitInBytesPerSecond()`

UnsetLogRateLimitInBytesPerSecond ensures that no value is present for LogRateLimitInBytesPerSecond, not even an explicit nil
### GetPerAppTasks

`func (o *OrganizationQuotaCreateApps) GetPerAppTasks() int32`

GetPerAppTasks returns the PerAppTasks field if non-nil, zero value otherwise.

### GetPerAppTasksOk

`func (o *OrganizationQuotaCreateApps) GetPerAppTasksOk() (*int32, bool)`

GetPerAppTasksOk returns a tuple with the PerAppTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAppTasks

`func (o *OrganizationQuotaCreateApps) SetPerAppTasks(v int32)`

SetPerAppTasks sets PerAppTasks field to given value.

### HasPerAppTasks

`func (o *OrganizationQuotaCreateApps) HasPerAppTasks() bool`

HasPerAppTasks returns a boolean if a field has been set.

### SetPerAppTasksNil

`func (o *OrganizationQuotaCreateApps) SetPerAppTasksNil(b bool)`

 SetPerAppTasksNil sets the value for PerAppTasks to be an explicit nil

### UnsetPerAppTasks
`func (o *OrganizationQuotaCreateApps) UnsetPerAppTasks()`

UnsetPerAppTasks ensures that no value is present for PerAppTasks, not even an explicit nil
### GetPerProcessMemoryInMb

`func (o *OrganizationQuotaCreateApps) GetPerProcessMemoryInMb() int32`

GetPerProcessMemoryInMb returns the PerProcessMemoryInMb field if non-nil, zero value otherwise.

### GetPerProcessMemoryInMbOk

`func (o *OrganizationQuotaCreateApps) GetPerProcessMemoryInMbOk() (*int32, bool)`

GetPerProcessMemoryInMbOk returns a tuple with the PerProcessMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcessMemoryInMb

`func (o *OrganizationQuotaCreateApps) SetPerProcessMemoryInMb(v int32)`

SetPerProcessMemoryInMb sets PerProcessMemoryInMb field to given value.

### HasPerProcessMemoryInMb

`func (o *OrganizationQuotaCreateApps) HasPerProcessMemoryInMb() bool`

HasPerProcessMemoryInMb returns a boolean if a field has been set.

### SetPerProcessMemoryInMbNil

`func (o *OrganizationQuotaCreateApps) SetPerProcessMemoryInMbNil(b bool)`

 SetPerProcessMemoryInMbNil sets the value for PerProcessMemoryInMb to be an explicit nil

### UnsetPerProcessMemoryInMb
`func (o *OrganizationQuotaCreateApps) UnsetPerProcessMemoryInMb()`

UnsetPerProcessMemoryInMb ensures that no value is present for PerProcessMemoryInMb, not even an explicit nil
### GetTotalInstances

`func (o *OrganizationQuotaCreateApps) GetTotalInstances() int32`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *OrganizationQuotaCreateApps) GetTotalInstancesOk() (*int32, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *OrganizationQuotaCreateApps) SetTotalInstances(v int32)`

SetTotalInstances sets TotalInstances field to given value.

### HasTotalInstances

`func (o *OrganizationQuotaCreateApps) HasTotalInstances() bool`

HasTotalInstances returns a boolean if a field has been set.

### SetTotalInstancesNil

`func (o *OrganizationQuotaCreateApps) SetTotalInstancesNil(b bool)`

 SetTotalInstancesNil sets the value for TotalInstances to be an explicit nil

### UnsetTotalInstances
`func (o *OrganizationQuotaCreateApps) UnsetTotalInstances()`

UnsetTotalInstances ensures that no value is present for TotalInstances, not even an explicit nil
### GetTotalMemoryInMb

`func (o *OrganizationQuotaCreateApps) GetTotalMemoryInMb() int32`

GetTotalMemoryInMb returns the TotalMemoryInMb field if non-nil, zero value otherwise.

### GetTotalMemoryInMbOk

`func (o *OrganizationQuotaCreateApps) GetTotalMemoryInMbOk() (*int32, bool)`

GetTotalMemoryInMbOk returns a tuple with the TotalMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryInMb

`func (o *OrganizationQuotaCreateApps) SetTotalMemoryInMb(v int32)`

SetTotalMemoryInMb sets TotalMemoryInMb field to given value.

### HasTotalMemoryInMb

`func (o *OrganizationQuotaCreateApps) HasTotalMemoryInMb() bool`

HasTotalMemoryInMb returns a boolean if a field has been set.

### SetTotalMemoryInMbNil

`func (o *OrganizationQuotaCreateApps) SetTotalMemoryInMbNil(b bool)`

 SetTotalMemoryInMbNil sets the value for TotalMemoryInMb to be an explicit nil

### UnsetTotalMemoryInMb
`func (o *OrganizationQuotaCreateApps) UnsetTotalMemoryInMb()`

UnsetTotalMemoryInMb ensures that no value is present for TotalMemoryInMb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


