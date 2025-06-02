# BuildpackLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Image reference for Cloud Native Buildpack | [optional] 
**Version** | Pointer to **string** | Version of the buildpack | [optional] 

## Methods

### NewBuildpackLifecycleData

`func NewBuildpackLifecycleData() *BuildpackLifecycleData`

NewBuildpackLifecycleData instantiates a new BuildpackLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildpackLifecycleDataWithDefaults

`func NewBuildpackLifecycleDataWithDefaults() *BuildpackLifecycleData`

NewBuildpackLifecycleDataWithDefaults instantiates a new BuildpackLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *BuildpackLifecycleData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BuildpackLifecycleData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BuildpackLifecycleData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BuildpackLifecycleData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVersion

`func (o *BuildpackLifecycleData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BuildpackLifecycleData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BuildpackLifecycleData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BuildpackLifecycleData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


