# AppUsageEventMemoryInMbPerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **NullableInt32** | Current memory in MB of the app that this event pertains to, if applicable | [optional] 
**Previous** | Pointer to **NullableInt32** | Previous memory in MB of the app that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventMemoryInMbPerInstance

`func NewAppUsageEventMemoryInMbPerInstance() *AppUsageEventMemoryInMbPerInstance`

NewAppUsageEventMemoryInMbPerInstance instantiates a new AppUsageEventMemoryInMbPerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventMemoryInMbPerInstanceWithDefaults

`func NewAppUsageEventMemoryInMbPerInstanceWithDefaults() *AppUsageEventMemoryInMbPerInstance`

NewAppUsageEventMemoryInMbPerInstanceWithDefaults instantiates a new AppUsageEventMemoryInMbPerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *AppUsageEventMemoryInMbPerInstance) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AppUsageEventMemoryInMbPerInstance) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AppUsageEventMemoryInMbPerInstance) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AppUsageEventMemoryInMbPerInstance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *AppUsageEventMemoryInMbPerInstance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AppUsageEventMemoryInMbPerInstance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetPrevious

`func (o *AppUsageEventMemoryInMbPerInstance) GetPrevious() int32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AppUsageEventMemoryInMbPerInstance) GetPreviousOk() (*int32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AppUsageEventMemoryInMbPerInstance) SetPrevious(v int32)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AppUsageEventMemoryInMbPerInstance) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *AppUsageEventMemoryInMbPerInstance) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AppUsageEventMemoryInMbPerInstance) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


