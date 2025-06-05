# AppUsageEventInstanceCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **NullableInt32** | Current instance count of the app that this event pertains to, if applicable | [optional] 
**Previous** | Pointer to **NullableInt32** | Previous instance count of the app that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventInstanceCount

`func NewAppUsageEventInstanceCount() *AppUsageEventInstanceCount`

NewAppUsageEventInstanceCount instantiates a new AppUsageEventInstanceCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventInstanceCountWithDefaults

`func NewAppUsageEventInstanceCountWithDefaults() *AppUsageEventInstanceCount`

NewAppUsageEventInstanceCountWithDefaults instantiates a new AppUsageEventInstanceCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *AppUsageEventInstanceCount) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AppUsageEventInstanceCount) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AppUsageEventInstanceCount) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AppUsageEventInstanceCount) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *AppUsageEventInstanceCount) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AppUsageEventInstanceCount) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetPrevious

`func (o *AppUsageEventInstanceCount) GetPrevious() int32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AppUsageEventInstanceCount) GetPreviousOk() (*int32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AppUsageEventInstanceCount) SetPrevious(v int32)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AppUsageEventInstanceCount) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *AppUsageEventInstanceCount) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AppUsageEventInstanceCount) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


