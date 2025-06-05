# Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The time the build was created | 
**CreatedBy** | [**BuildCreatedBy**](BuildCreatedBy.md) |  | 
**Droplet** | [**NullableBuildDroplet**](BuildDroplet.md) |  | 
**Error** | **NullableString** | Error message if build failed | 
**Guid** | **string** | Unique identifier for the build | 
**Lifecycle** | [**BuildLifecycle**](BuildLifecycle.md) |  | 
**Links** | [**BuildLinks**](BuildLinks.md) |  | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**Package** | [**BuildPackage**](BuildPackage.md) |  | 
**StagingDiskInMb** | **int32** | Disk space allocated for staging | 
**StagingLogRateLimitBytesPerSecond** | **int32** | Log rate limit during staging (-1 for unlimited) | 
**StagingMemoryInMb** | **int32** | Memory allocated for staging | 
**State** | **string** | Current state of the build | 
**UpdatedAt** | **time.Time** | The time the build was last updated | 

## Methods

### NewBuild

`func NewBuild(createdAt time.Time, createdBy BuildCreatedBy, droplet NullableBuildDroplet, error_ NullableString, guid string, lifecycle BuildLifecycle, links BuildLinks, metadata V3AppsGuidTasksPostRequestMetadata, package_ BuildPackage, stagingDiskInMb int32, stagingLogRateLimitBytesPerSecond int32, stagingMemoryInMb int32, state string, updatedAt time.Time, ) *Build`

NewBuild instantiates a new Build object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildWithDefaults

`func NewBuildWithDefaults() *Build`

NewBuildWithDefaults instantiates a new Build object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Build) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Build) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Build) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Build) GetCreatedBy() BuildCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Build) GetCreatedByOk() (*BuildCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Build) SetCreatedBy(v BuildCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.


### GetDroplet

`func (o *Build) GetDroplet() BuildDroplet`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *Build) GetDropletOk() (*BuildDroplet, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *Build) SetDroplet(v BuildDroplet)`

SetDroplet sets Droplet field to given value.


### SetDropletNil

`func (o *Build) SetDropletNil(b bool)`

 SetDropletNil sets the value for Droplet to be an explicit nil

### UnsetDroplet
`func (o *Build) UnsetDroplet()`

UnsetDroplet ensures that no value is present for Droplet, not even an explicit nil
### GetError

`func (o *Build) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Build) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Build) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *Build) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Build) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetGuid

`func (o *Build) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Build) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Build) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLifecycle

`func (o *Build) GetLifecycle() BuildLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *Build) GetLifecycleOk() (*BuildLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *Build) SetLifecycle(v BuildLifecycle)`

SetLifecycle sets Lifecycle field to given value.


### GetLinks

`func (o *Build) GetLinks() BuildLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Build) GetLinksOk() (*BuildLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Build) SetLinks(v BuildLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *Build) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Build) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Build) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.


### GetPackage

`func (o *Build) GetPackage() BuildPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *Build) GetPackageOk() (*BuildPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *Build) SetPackage(v BuildPackage)`

SetPackage sets Package field to given value.


### GetStagingDiskInMb

`func (o *Build) GetStagingDiskInMb() int32`

GetStagingDiskInMb returns the StagingDiskInMb field if non-nil, zero value otherwise.

### GetStagingDiskInMbOk

`func (o *Build) GetStagingDiskInMbOk() (*int32, bool)`

GetStagingDiskInMbOk returns a tuple with the StagingDiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDiskInMb

`func (o *Build) SetStagingDiskInMb(v int32)`

SetStagingDiskInMb sets StagingDiskInMb field to given value.


### GetStagingLogRateLimitBytesPerSecond

`func (o *Build) GetStagingLogRateLimitBytesPerSecond() int32`

GetStagingLogRateLimitBytesPerSecond returns the StagingLogRateLimitBytesPerSecond field if non-nil, zero value otherwise.

### GetStagingLogRateLimitBytesPerSecondOk

`func (o *Build) GetStagingLogRateLimitBytesPerSecondOk() (*int32, bool)`

GetStagingLogRateLimitBytesPerSecondOk returns a tuple with the StagingLogRateLimitBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingLogRateLimitBytesPerSecond

`func (o *Build) SetStagingLogRateLimitBytesPerSecond(v int32)`

SetStagingLogRateLimitBytesPerSecond sets StagingLogRateLimitBytesPerSecond field to given value.


### GetStagingMemoryInMb

`func (o *Build) GetStagingMemoryInMb() int32`

GetStagingMemoryInMb returns the StagingMemoryInMb field if non-nil, zero value otherwise.

### GetStagingMemoryInMbOk

`func (o *Build) GetStagingMemoryInMbOk() (*int32, bool)`

GetStagingMemoryInMbOk returns a tuple with the StagingMemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMemoryInMb

`func (o *Build) SetStagingMemoryInMb(v int32)`

SetStagingMemoryInMb sets StagingMemoryInMb field to given value.


### GetState

`func (o *Build) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Build) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Build) SetState(v string)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *Build) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Build) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Build) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


