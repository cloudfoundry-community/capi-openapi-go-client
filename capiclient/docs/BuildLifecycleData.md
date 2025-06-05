# BuildLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildpackLifecycleBuildpacks** | Pointer to [**[]V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner**](V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner.md) | Detected buildpacks with versions | [optional] 
**Buildpacks** | Pointer to **[]string** | Buildpacks used (buildpack lifecycle) | [optional] 
**Image** | Pointer to **string** | Docker image (docker lifecycle) or resulting image (cnb) | [optional] 
**Stack** | Pointer to **string** | Stack used for staging | [optional] 

## Methods

### NewBuildLifecycleData

`func NewBuildLifecycleData() *BuildLifecycleData`

NewBuildLifecycleData instantiates a new BuildLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildLifecycleDataWithDefaults

`func NewBuildLifecycleDataWithDefaults() *BuildLifecycleData`

NewBuildLifecycleDataWithDefaults instantiates a new BuildLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpackLifecycleBuildpacks

`func (o *BuildLifecycleData) GetBuildpackLifecycleBuildpacks() []V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner`

GetBuildpackLifecycleBuildpacks returns the BuildpackLifecycleBuildpacks field if non-nil, zero value otherwise.

### GetBuildpackLifecycleBuildpacksOk

`func (o *BuildLifecycleData) GetBuildpackLifecycleBuildpacksOk() (*[]V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner, bool)`

GetBuildpackLifecycleBuildpacksOk returns a tuple with the BuildpackLifecycleBuildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLifecycleBuildpacks

`func (o *BuildLifecycleData) SetBuildpackLifecycleBuildpacks(v []V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner)`

SetBuildpackLifecycleBuildpacks sets BuildpackLifecycleBuildpacks field to given value.

### HasBuildpackLifecycleBuildpacks

`func (o *BuildLifecycleData) HasBuildpackLifecycleBuildpacks() bool`

HasBuildpackLifecycleBuildpacks returns a boolean if a field has been set.

### GetBuildpacks

`func (o *BuildLifecycleData) GetBuildpacks() []string`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *BuildLifecycleData) GetBuildpacksOk() (*[]string, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *BuildLifecycleData) SetBuildpacks(v []string)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *BuildLifecycleData) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetImage

`func (o *BuildLifecycleData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BuildLifecycleData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BuildLifecycleData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BuildLifecycleData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetStack

`func (o *BuildLifecycleData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *BuildLifecycleData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *BuildLifecycleData) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *BuildLifecycleData) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


