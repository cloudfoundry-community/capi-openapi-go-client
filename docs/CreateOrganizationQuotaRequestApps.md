# CreateOrganizationQuotaRequestApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogRateLimitInBytesPerSecond** | Pointer to **NullableInt32** | Total log rate limit allowed for all the started processes and running tasks in an organization | [optional] 
**PerAppTasks** | Pointer to **NullableInt32** | Maximum number of running tasks in an organization | [optional] 
**PerProcessMemoryInMb** | Pointer to **NullableInt32** | Maximum memory for a single process or task | [optional] 
**TotalInstances** | Pointer to **NullableInt32** | Total instances of all the started processes allowed in an organization | [optional] 
**TotalMemoryInMb** | Pointer to **NullableInt32** | Total memory allowed for all the started processes and running tasks in an organization | [optional] 

## Methods

### NewCreateOrganizationQuotaRequestApps

`func NewCreateOrganizationQuotaRequestApps() *CreateOrganizationQuotaRequestApps`

NewCreateOrganizationQuotaRequestApps instantiates a new CreateOrganizationQuotaRequestApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationQuotaRequestAppsWithDefaults

`func NewCreateOrganizationQuotaRequestAppsWithDefaults() *CreateOrganizationQuotaRequestApps`

NewCreateOrganizationQuotaRequestAppsWithDefaults instantiates a new CreateOrganizationQuotaRequestApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogRateLimitInBytesPerSecond

`func (o *CreateOrganizationQuotaRequestApps) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *CreateOrganizationQuotaRequestApps) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *CreateOrganizationQuotaRequestApps) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.

### HasLogRateLimitInBytesPerSecond

`func (o *CreateOrganizationQuotaRequestApps) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

### SetLogRateLimitInBytesPerSecondNil

`func (o *CreateOrganizationQuotaRequestApps) SetLogRateLimitInBytesPerSecondNil(b bool)`

 SetLogRateLimitInBytesPerSecondNil sets the value for LogRateLimitInBytesPerSecond to be an explicit nil

### UnsetLogRateLimitInBytesPerSecond
`func (o *CreateOrganizationQuotaRequestApps) UnsetLogRateLimitInBytesPerSecond()`

UnsetLogRateLimitInBytesPerSecond ensures that no value is present for LogRateLimitInBytesPerSecond, not even an explicit nil
### GetPerAppTasks

`func (o *CreateOrganizationQuotaRequestApps) GetPerAppTasks() int32`

GetPerAppTasks returns the PerAppTasks field if non-nil, zero value otherwise.

### GetPerAppTasksOk

`func (o *CreateOrganizationQuotaRequestApps) GetPerAppTasksOk() (*int32, bool)`

GetPerAppTasksOk returns a tuple with the PerAppTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAppTasks

`func (o *CreateOrganizationQuotaRequestApps) SetPerAppTasks(v int32)`

SetPerAppTasks sets PerAppTasks field to given value.

### HasPerAppTasks

`func (o *CreateOrganizationQuotaRequestApps) HasPerAppTasks() bool`

HasPerAppTasks returns a boolean if a field has been set.

### SetPerAppTasksNil

`func (o *CreateOrganizationQuotaRequestApps) SetPerAppTasksNil(b bool)`

 SetPerAppTasksNil sets the value for PerAppTasks to be an explicit nil

### UnsetPerAppTasks
`func (o *CreateOrganizationQuotaRequestApps) UnsetPerAppTasks()`

UnsetPerAppTasks ensures that no value is present for PerAppTasks, not even an explicit nil
### GetPerProcessMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) GetPerProcessMemoryInMb() int32`

GetPerProcessMemoryInMb returns the PerProcessMemoryInMb field if non-nil, zero value otherwise.

### GetPerProcessMemoryInMbOk

`func (o *CreateOrganizationQuotaRequestApps) GetPerProcessMemoryInMbOk() (*int32, bool)`

GetPerProcessMemoryInMbOk returns a tuple with the PerProcessMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcessMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) SetPerProcessMemoryInMb(v int32)`

SetPerProcessMemoryInMb sets PerProcessMemoryInMb field to given value.

### HasPerProcessMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) HasPerProcessMemoryInMb() bool`

HasPerProcessMemoryInMb returns a boolean if a field has been set.

### SetPerProcessMemoryInMbNil

`func (o *CreateOrganizationQuotaRequestApps) SetPerProcessMemoryInMbNil(b bool)`

 SetPerProcessMemoryInMbNil sets the value for PerProcessMemoryInMb to be an explicit nil

### UnsetPerProcessMemoryInMb
`func (o *CreateOrganizationQuotaRequestApps) UnsetPerProcessMemoryInMb()`

UnsetPerProcessMemoryInMb ensures that no value is present for PerProcessMemoryInMb, not even an explicit nil
### GetTotalInstances

`func (o *CreateOrganizationQuotaRequestApps) GetTotalInstances() int32`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *CreateOrganizationQuotaRequestApps) GetTotalInstancesOk() (*int32, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *CreateOrganizationQuotaRequestApps) SetTotalInstances(v int32)`

SetTotalInstances sets TotalInstances field to given value.

### HasTotalInstances

`func (o *CreateOrganizationQuotaRequestApps) HasTotalInstances() bool`

HasTotalInstances returns a boolean if a field has been set.

### SetTotalInstancesNil

`func (o *CreateOrganizationQuotaRequestApps) SetTotalInstancesNil(b bool)`

 SetTotalInstancesNil sets the value for TotalInstances to be an explicit nil

### UnsetTotalInstances
`func (o *CreateOrganizationQuotaRequestApps) UnsetTotalInstances()`

UnsetTotalInstances ensures that no value is present for TotalInstances, not even an explicit nil
### GetTotalMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) GetTotalMemoryInMb() int32`

GetTotalMemoryInMb returns the TotalMemoryInMb field if non-nil, zero value otherwise.

### GetTotalMemoryInMbOk

`func (o *CreateOrganizationQuotaRequestApps) GetTotalMemoryInMbOk() (*int32, bool)`

GetTotalMemoryInMbOk returns a tuple with the TotalMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) SetTotalMemoryInMb(v int32)`

SetTotalMemoryInMb sets TotalMemoryInMb field to given value.

### HasTotalMemoryInMb

`func (o *CreateOrganizationQuotaRequestApps) HasTotalMemoryInMb() bool`

HasTotalMemoryInMb returns a boolean if a field has been set.

### SetTotalMemoryInMbNil

`func (o *CreateOrganizationQuotaRequestApps) SetTotalMemoryInMbNil(b bool)`

 SetTotalMemoryInMbNil sets the value for TotalMemoryInMb to be an explicit nil

### UnsetTotalMemoryInMb
`func (o *CreateOrganizationQuotaRequestApps) UnsetTotalMemoryInMb()`

UnsetTotalMemoryInMb ensures that no value is present for TotalMemoryInMb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


