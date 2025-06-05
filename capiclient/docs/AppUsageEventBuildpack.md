# AppUsageEventBuildpack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **NullableString** | Unique identifier of the buildpack that this event pertains to, if applicable | [optional] 
**Name** | Pointer to **NullableString** | Name of the buildpack that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventBuildpack

`func NewAppUsageEventBuildpack() *AppUsageEventBuildpack`

NewAppUsageEventBuildpack instantiates a new AppUsageEventBuildpack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventBuildpackWithDefaults

`func NewAppUsageEventBuildpackWithDefaults() *AppUsageEventBuildpack`

NewAppUsageEventBuildpackWithDefaults instantiates a new AppUsageEventBuildpack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AppUsageEventBuildpack) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppUsageEventBuildpack) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppUsageEventBuildpack) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppUsageEventBuildpack) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AppUsageEventBuildpack) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AppUsageEventBuildpack) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *AppUsageEventBuildpack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUsageEventBuildpack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUsageEventBuildpack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppUsageEventBuildpack) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppUsageEventBuildpack) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppUsageEventBuildpack) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


