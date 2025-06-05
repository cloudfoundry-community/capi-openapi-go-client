# PlatformInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **string** |  | [optional] 
**CliVersion** | Pointer to [**PlatformInfoCliVersion**](PlatformInfoCliVersion.md) |  | [optional] 
**Custom** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**PlatformInfoLinks**](PlatformInfoLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlatformInfo

`func NewPlatformInfo() *PlatformInfo`

NewPlatformInfo instantiates a new PlatformInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInfoWithDefaults

`func NewPlatformInfoWithDefaults() *PlatformInfo`

NewPlatformInfoWithDefaults instantiates a new PlatformInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *PlatformInfo) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *PlatformInfo) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *PlatformInfo) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *PlatformInfo) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetCliVersion

`func (o *PlatformInfo) GetCliVersion() PlatformInfoCliVersion`

GetCliVersion returns the CliVersion field if non-nil, zero value otherwise.

### GetCliVersionOk

`func (o *PlatformInfo) GetCliVersionOk() (*PlatformInfoCliVersion, bool)`

GetCliVersionOk returns a tuple with the CliVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVersion

`func (o *PlatformInfo) SetCliVersion(v PlatformInfoCliVersion)`

SetCliVersion sets CliVersion field to given value.

### HasCliVersion

`func (o *PlatformInfo) HasCliVersion() bool`

HasCliVersion returns a boolean if a field has been set.

### GetCustom

`func (o *PlatformInfo) GetCustom() map[string]string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *PlatformInfo) GetCustomOk() (*map[string]string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *PlatformInfo) SetCustom(v map[string]string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *PlatformInfo) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *PlatformInfo) GetLinks() PlatformInfoLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PlatformInfo) GetLinksOk() (*PlatformInfoLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PlatformInfo) SetLinks(v PlatformInfoLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PlatformInfo) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *PlatformInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PlatformInfo) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlatformInfo) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlatformInfo) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlatformInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


