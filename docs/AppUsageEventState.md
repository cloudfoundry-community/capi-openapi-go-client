# AppUsageEventState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **NullableString** | Current state of the app that this event pertains to, if applicable | [optional] 
**Previous** | Pointer to **NullableString** | Previous state of the app that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventState

`func NewAppUsageEventState() *AppUsageEventState`

NewAppUsageEventState instantiates a new AppUsageEventState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventStateWithDefaults

`func NewAppUsageEventStateWithDefaults() *AppUsageEventState`

NewAppUsageEventStateWithDefaults instantiates a new AppUsageEventState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *AppUsageEventState) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AppUsageEventState) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AppUsageEventState) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AppUsageEventState) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *AppUsageEventState) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AppUsageEventState) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetPrevious

`func (o *AppUsageEventState) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AppUsageEventState) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AppUsageEventState) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AppUsageEventState) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *AppUsageEventState) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AppUsageEventState) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


