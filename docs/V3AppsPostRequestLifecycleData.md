# V3AppsPostRequestLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to **[]string** | List of buildpacks for buildpack/cnb lifecycle | [optional] 
**Stack** | Pointer to **string** | Stack to use for buildpack lifecycle | [optional] 

## Methods

### NewV3AppsPostRequestLifecycleData

`func NewV3AppsPostRequestLifecycleData() *V3AppsPostRequestLifecycleData`

NewV3AppsPostRequestLifecycleData instantiates a new V3AppsPostRequestLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsPostRequestLifecycleDataWithDefaults

`func NewV3AppsPostRequestLifecycleDataWithDefaults() *V3AppsPostRequestLifecycleData`

NewV3AppsPostRequestLifecycleDataWithDefaults instantiates a new V3AppsPostRequestLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *V3AppsPostRequestLifecycleData) GetBuildpacks() []string`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *V3AppsPostRequestLifecycleData) GetBuildpacksOk() (*[]string, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *V3AppsPostRequestLifecycleData) SetBuildpacks(v []string)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *V3AppsPostRequestLifecycleData) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetStack

`func (o *V3AppsPostRequestLifecycleData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V3AppsPostRequestLifecycleData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V3AppsPostRequestLifecycleData) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V3AppsPostRequestLifecycleData) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


