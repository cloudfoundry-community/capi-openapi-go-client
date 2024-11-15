# Buildpack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildpackName** | Pointer to **string** |  | [optional] 
**DetectOutput** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildpack

`func NewBuildpack() *Buildpack`

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

### HasName

`func (o *Buildpack) HasName() bool`

HasName returns a boolean if a field has been set.

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


