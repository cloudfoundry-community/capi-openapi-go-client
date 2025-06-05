# AppUsageEventApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **NullableString** | Unique identifier of the app that this event pertains to, if applicable | [optional] 
**Name** | Pointer to **NullableString** | Name of the app that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventApp

`func NewAppUsageEventApp() *AppUsageEventApp`

NewAppUsageEventApp instantiates a new AppUsageEventApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventAppWithDefaults

`func NewAppUsageEventAppWithDefaults() *AppUsageEventApp`

NewAppUsageEventAppWithDefaults instantiates a new AppUsageEventApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AppUsageEventApp) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppUsageEventApp) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppUsageEventApp) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppUsageEventApp) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AppUsageEventApp) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AppUsageEventApp) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *AppUsageEventApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUsageEventApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUsageEventApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppUsageEventApp) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppUsageEventApp) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppUsageEventApp) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


