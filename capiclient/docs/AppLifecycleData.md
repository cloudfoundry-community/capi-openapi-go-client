# AppLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to **[]string** | Buildpacks for buildpack/cnb lifecycle | [optional] 
**Image** | Pointer to **string** | Docker image reference | [optional] 
**Stack** | Pointer to **string** | Stack for buildpack lifecycle | [optional] 

## Methods

### NewAppLifecycleData

`func NewAppLifecycleData() *AppLifecycleData`

NewAppLifecycleData instantiates a new AppLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLifecycleDataWithDefaults

`func NewAppLifecycleDataWithDefaults() *AppLifecycleData`

NewAppLifecycleDataWithDefaults instantiates a new AppLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *AppLifecycleData) GetBuildpacks() []string`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *AppLifecycleData) GetBuildpacksOk() (*[]string, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *AppLifecycleData) SetBuildpacks(v []string)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *AppLifecycleData) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetImage

`func (o *AppLifecycleData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppLifecycleData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppLifecycleData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AppLifecycleData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetStack

`func (o *AppLifecycleData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *AppLifecycleData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *AppLifecycleData) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *AppLifecycleData) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


