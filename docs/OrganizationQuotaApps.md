# OrganizationQuotaApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogRateLimitInBytesPerSecond** | **NullableInt32** | Log rate limit per second (null &#x3D; unlimited) | 
**PerAppTasks** | **NullableInt32** | Maximum tasks per app (null &#x3D; unlimited) | 
**PerProcessMemoryInMb** | **NullableInt32** | Memory limit per app process (null &#x3D; unlimited) | 
**TotalInstances** | **NullableInt32** | Total app instances allowed (null &#x3D; unlimited) | 
**TotalMemoryInMb** | **NullableInt32** | Total memory limit for all apps (null &#x3D; unlimited) | 

## Methods

### NewOrganizationQuotaApps

`func NewOrganizationQuotaApps(logRateLimitInBytesPerSecond NullableInt32, perAppTasks NullableInt32, perProcessMemoryInMb NullableInt32, totalInstances NullableInt32, totalMemoryInMb NullableInt32, ) *OrganizationQuotaApps`

NewOrganizationQuotaApps instantiates a new OrganizationQuotaApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaAppsWithDefaults

`func NewOrganizationQuotaAppsWithDefaults() *OrganizationQuotaApps`

NewOrganizationQuotaAppsWithDefaults instantiates a new OrganizationQuotaApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogRateLimitInBytesPerSecond

`func (o *OrganizationQuotaApps) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *OrganizationQuotaApps) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *OrganizationQuotaApps) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.


### SetLogRateLimitInBytesPerSecondNil

`func (o *OrganizationQuotaApps) SetLogRateLimitInBytesPerSecondNil(b bool)`

 SetLogRateLimitInBytesPerSecondNil sets the value for LogRateLimitInBytesPerSecond to be an explicit nil

### UnsetLogRateLimitInBytesPerSecond
`func (o *OrganizationQuotaApps) UnsetLogRateLimitInBytesPerSecond()`

UnsetLogRateLimitInBytesPerSecond ensures that no value is present for LogRateLimitInBytesPerSecond, not even an explicit nil
### GetPerAppTasks

`func (o *OrganizationQuotaApps) GetPerAppTasks() int32`

GetPerAppTasks returns the PerAppTasks field if non-nil, zero value otherwise.

### GetPerAppTasksOk

`func (o *OrganizationQuotaApps) GetPerAppTasksOk() (*int32, bool)`

GetPerAppTasksOk returns a tuple with the PerAppTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAppTasks

`func (o *OrganizationQuotaApps) SetPerAppTasks(v int32)`

SetPerAppTasks sets PerAppTasks field to given value.


### SetPerAppTasksNil

`func (o *OrganizationQuotaApps) SetPerAppTasksNil(b bool)`

 SetPerAppTasksNil sets the value for PerAppTasks to be an explicit nil

### UnsetPerAppTasks
`func (o *OrganizationQuotaApps) UnsetPerAppTasks()`

UnsetPerAppTasks ensures that no value is present for PerAppTasks, not even an explicit nil
### GetPerProcessMemoryInMb

`func (o *OrganizationQuotaApps) GetPerProcessMemoryInMb() int32`

GetPerProcessMemoryInMb returns the PerProcessMemoryInMb field if non-nil, zero value otherwise.

### GetPerProcessMemoryInMbOk

`func (o *OrganizationQuotaApps) GetPerProcessMemoryInMbOk() (*int32, bool)`

GetPerProcessMemoryInMbOk returns a tuple with the PerProcessMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcessMemoryInMb

`func (o *OrganizationQuotaApps) SetPerProcessMemoryInMb(v int32)`

SetPerProcessMemoryInMb sets PerProcessMemoryInMb field to given value.


### SetPerProcessMemoryInMbNil

`func (o *OrganizationQuotaApps) SetPerProcessMemoryInMbNil(b bool)`

 SetPerProcessMemoryInMbNil sets the value for PerProcessMemoryInMb to be an explicit nil

### UnsetPerProcessMemoryInMb
`func (o *OrganizationQuotaApps) UnsetPerProcessMemoryInMb()`

UnsetPerProcessMemoryInMb ensures that no value is present for PerProcessMemoryInMb, not even an explicit nil
### GetTotalInstances

`func (o *OrganizationQuotaApps) GetTotalInstances() int32`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *OrganizationQuotaApps) GetTotalInstancesOk() (*int32, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *OrganizationQuotaApps) SetTotalInstances(v int32)`

SetTotalInstances sets TotalInstances field to given value.


### SetTotalInstancesNil

`func (o *OrganizationQuotaApps) SetTotalInstancesNil(b bool)`

 SetTotalInstancesNil sets the value for TotalInstances to be an explicit nil

### UnsetTotalInstances
`func (o *OrganizationQuotaApps) UnsetTotalInstances()`

UnsetTotalInstances ensures that no value is present for TotalInstances, not even an explicit nil
### GetTotalMemoryInMb

`func (o *OrganizationQuotaApps) GetTotalMemoryInMb() int32`

GetTotalMemoryInMb returns the TotalMemoryInMb field if non-nil, zero value otherwise.

### GetTotalMemoryInMbOk

`func (o *OrganizationQuotaApps) GetTotalMemoryInMbOk() (*int32, bool)`

GetTotalMemoryInMbOk returns a tuple with the TotalMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryInMb

`func (o *OrganizationQuotaApps) SetTotalMemoryInMb(v int32)`

SetTotalMemoryInMb sets TotalMemoryInMb field to given value.


### SetTotalMemoryInMbNil

`func (o *OrganizationQuotaApps) SetTotalMemoryInMbNil(b bool)`

 SetTotalMemoryInMbNil sets the value for TotalMemoryInMb to be an explicit nil

### UnsetTotalMemoryInMb
`func (o *OrganizationQuotaApps) UnsetTotalMemoryInMb()`

UnsetTotalMemoryInMb ensures that no value is present for TotalMemoryInMb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


