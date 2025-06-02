# Buildpack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildpackName** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** | The time the buildpack was created | 
**DetectOutput** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** | Whether buildpack is available for use | 
**Filename** | **NullableString** | Filename of uploaded buildpack | 
**Guid** | **string** | Unique identifier for the buildpack | 
**Lifecycle** | Pointer to [**BuildpackLifecycle**](BuildpackLifecycle.md) |  | [optional] 
**Links** | [**BuildpackLinks**](BuildpackLinks.md) |  | 
**Locked** | **bool** | Whether buildpack updates are prevented | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**Name** | **string** | Name of the buildpack | 
**Position** | **int32** | Priority position for automatic detection | 
**Stack** | **NullableString** | Stack the buildpack runs on | 
**State** | **string** | Current state of the buildpack | 
**UpdatedAt** | **time.Time** | The time the buildpack was last updated | 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildpack

`func NewBuildpack(createdAt time.Time, enabled bool, filename NullableString, guid string, links BuildpackLinks, locked bool, metadata V3AppsGuidTasksPostRequestMetadata, name string, position int32, stack NullableString, state string, updatedAt time.Time, ) *Buildpack`

NewBuildpack instantiates a new Buildpack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildpackWithDefaults

`func NewBuildpackWithDefaults() *Buildpack`

NewBuildpackWithDefaults instantiates a new Buildpack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpackName

`func (o *Buildpack) GetBuildpackName() string`

GetBuildpackName returns the BuildpackName field if non-nil, zero value otherwise.

### GetBuildpackNameOk

`func (o *Buildpack) GetBuildpackNameOk() (*string, bool)`

GetBuildpackNameOk returns a tuple with the BuildpackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackName

`func (o *Buildpack) SetBuildpackName(v string)`

SetBuildpackName sets BuildpackName field to given value.

### HasBuildpackName

`func (o *Buildpack) HasBuildpackName() bool`

HasBuildpackName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Buildpack) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Buildpack) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Buildpack) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetectOutput

`func (o *Buildpack) GetDetectOutput() string`

GetDetectOutput returns the DetectOutput field if non-nil, zero value otherwise.

### GetDetectOutputOk

`func (o *Buildpack) GetDetectOutputOk() (*string, bool)`

GetDetectOutputOk returns a tuple with the DetectOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectOutput

`func (o *Buildpack) SetDetectOutput(v string)`

SetDetectOutput sets DetectOutput field to given value.

### HasDetectOutput

`func (o *Buildpack) HasDetectOutput() bool`

HasDetectOutput returns a boolean if a field has been set.

### GetEnabled

`func (o *Buildpack) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Buildpack) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Buildpack) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFilename

`func (o *Buildpack) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Buildpack) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Buildpack) SetFilename(v string)`

SetFilename sets Filename field to given value.


### SetFilenameNil

`func (o *Buildpack) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *Buildpack) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetGuid

`func (o *Buildpack) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Buildpack) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Buildpack) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLifecycle

`func (o *Buildpack) GetLifecycle() BuildpackLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *Buildpack) GetLifecycleOk() (*BuildpackLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *Buildpack) SetLifecycle(v BuildpackLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *Buildpack) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLinks

`func (o *Buildpack) GetLinks() BuildpackLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Buildpack) GetLinksOk() (*BuildpackLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Buildpack) SetLinks(v BuildpackLinks)`

SetLinks sets Links field to given value.


### GetLocked

`func (o *Buildpack) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Buildpack) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Buildpack) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMetadata

`func (o *Buildpack) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Buildpack) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Buildpack) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *Buildpack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Buildpack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Buildpack) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *Buildpack) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Buildpack) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Buildpack) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetStack

`func (o *Buildpack) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *Buildpack) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *Buildpack) SetStack(v string)`

SetStack sets Stack field to given value.


### SetStackNil

`func (o *Buildpack) SetStackNil(b bool)`

 SetStackNil sets the value for Stack to be an explicit nil

### UnsetStack
`func (o *Buildpack) UnsetStack()`

UnsetStack ensures that no value is present for Stack, not even an explicit nil
### GetState

`func (o *Buildpack) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Buildpack) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Buildpack) SetState(v string)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *Buildpack) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Buildpack) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Buildpack) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Buildpack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Buildpack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Buildpack) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Buildpack) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


