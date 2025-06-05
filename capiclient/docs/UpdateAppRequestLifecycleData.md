# UpdateAppRequestLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to **[]string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAppRequestLifecycleData

`func NewUpdateAppRequestLifecycleData() *UpdateAppRequestLifecycleData`

NewUpdateAppRequestLifecycleData instantiates a new UpdateAppRequestLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppRequestLifecycleDataWithDefaults

`func NewUpdateAppRequestLifecycleDataWithDefaults() *UpdateAppRequestLifecycleData`

NewUpdateAppRequestLifecycleDataWithDefaults instantiates a new UpdateAppRequestLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *UpdateAppRequestLifecycleData) GetBuildpacks() []string`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *UpdateAppRequestLifecycleData) GetBuildpacksOk() (*[]string, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *UpdateAppRequestLifecycleData) SetBuildpacks(v []string)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *UpdateAppRequestLifecycleData) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetStack

`func (o *UpdateAppRequestLifecycleData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *UpdateAppRequestLifecycleData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *UpdateAppRequestLifecycleData) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *UpdateAppRequestLifecycleData) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


