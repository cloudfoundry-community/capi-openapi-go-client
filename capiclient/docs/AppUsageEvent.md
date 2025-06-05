# AppUsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppUsageEventApp**](AppUsageEventApp.md) |  | [optional] 
**Buildpack** | Pointer to [**AppUsageEventBuildpack**](AppUsageEventBuildpack.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time with zone when the event occurred | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the event | [optional] 
**InstanceCount** | Pointer to [**AppUsageEventInstanceCount**](AppUsageEventInstanceCount.md) |  | [optional] 
**Links** | Pointer to [**AppUsageEventLinks**](AppUsageEventLinks.md) |  | [optional] 
**MemoryInMbPerInstance** | Pointer to [**AppUsageEventMemoryInMbPerInstance**](AppUsageEventMemoryInMbPerInstance.md) |  | [optional] 
**Organization** | Pointer to [**AppUsageEventOrganization**](AppUsageEventOrganization.md) |  | [optional] 
**Process** | Pointer to [**AppUsageEventProcess**](AppUsageEventProcess.md) |  | [optional] 
**Space** | Pointer to [**AppUsageEventSpace**](AppUsageEventSpace.md) |  | [optional] 
**State** | Pointer to [**AppUsageEventState**](AppUsageEventState.md) |  | [optional] 
**Task** | Pointer to [**AppUsageEventTask**](AppUsageEventTask.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Identical to created_at (events are created, never updated) | [optional] 

## Methods

### NewAppUsageEvent

`func NewAppUsageEvent() *AppUsageEvent`

NewAppUsageEvent instantiates a new AppUsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventWithDefaults

`func NewAppUsageEventWithDefaults() *AppUsageEvent`

NewAppUsageEventWithDefaults instantiates a new AppUsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AppUsageEvent) GetApp() AppUsageEventApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppUsageEvent) GetAppOk() (*AppUsageEventApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppUsageEvent) SetApp(v AppUsageEventApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AppUsageEvent) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBuildpack

`func (o *AppUsageEvent) GetBuildpack() AppUsageEventBuildpack`

GetBuildpack returns the Buildpack field if non-nil, zero value otherwise.

### GetBuildpackOk

`func (o *AppUsageEvent) GetBuildpackOk() (*AppUsageEventBuildpack, bool)`

GetBuildpackOk returns a tuple with the Buildpack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpack

`func (o *AppUsageEvent) SetBuildpack(v AppUsageEventBuildpack)`

SetBuildpack sets Buildpack field to given value.

### HasBuildpack

`func (o *AppUsageEvent) HasBuildpack() bool`

HasBuildpack returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppUsageEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppUsageEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppUsageEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppUsageEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *AppUsageEvent) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppUsageEvent) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppUsageEvent) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppUsageEvent) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AppUsageEvent) GetInstanceCount() AppUsageEventInstanceCount`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AppUsageEvent) GetInstanceCountOk() (*AppUsageEventInstanceCount, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AppUsageEvent) SetInstanceCount(v AppUsageEventInstanceCount)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AppUsageEvent) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetLinks

`func (o *AppUsageEvent) GetLinks() AppUsageEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppUsageEvent) GetLinksOk() (*AppUsageEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppUsageEvent) SetLinks(v AppUsageEventLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppUsageEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMemoryInMbPerInstance

`func (o *AppUsageEvent) GetMemoryInMbPerInstance() AppUsageEventMemoryInMbPerInstance`

GetMemoryInMbPerInstance returns the MemoryInMbPerInstance field if non-nil, zero value otherwise.

### GetMemoryInMbPerInstanceOk

`func (o *AppUsageEvent) GetMemoryInMbPerInstanceOk() (*AppUsageEventMemoryInMbPerInstance, bool)`

GetMemoryInMbPerInstanceOk returns a tuple with the MemoryInMbPerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMbPerInstance

`func (o *AppUsageEvent) SetMemoryInMbPerInstance(v AppUsageEventMemoryInMbPerInstance)`

SetMemoryInMbPerInstance sets MemoryInMbPerInstance field to given value.

### HasMemoryInMbPerInstance

`func (o *AppUsageEvent) HasMemoryInMbPerInstance() bool`

HasMemoryInMbPerInstance returns a boolean if a field has been set.

### GetOrganization

`func (o *AppUsageEvent) GetOrganization() AppUsageEventOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AppUsageEvent) GetOrganizationOk() (*AppUsageEventOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AppUsageEvent) SetOrganization(v AppUsageEventOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AppUsageEvent) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProcess

`func (o *AppUsageEvent) GetProcess() AppUsageEventProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *AppUsageEvent) GetProcessOk() (*AppUsageEventProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *AppUsageEvent) SetProcess(v AppUsageEventProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *AppUsageEvent) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetSpace

`func (o *AppUsageEvent) GetSpace() AppUsageEventSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *AppUsageEvent) GetSpaceOk() (*AppUsageEventSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *AppUsageEvent) SetSpace(v AppUsageEventSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *AppUsageEvent) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetState

`func (o *AppUsageEvent) GetState() AppUsageEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppUsageEvent) GetStateOk() (*AppUsageEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppUsageEvent) SetState(v AppUsageEventState)`

SetState sets State field to given value.

### HasState

`func (o *AppUsageEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTask

`func (o *AppUsageEvent) GetTask() AppUsageEventTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AppUsageEvent) GetTaskOk() (*AppUsageEventTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AppUsageEvent) SetTask(v AppUsageEventTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *AppUsageEvent) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppUsageEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppUsageEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppUsageEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppUsageEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


