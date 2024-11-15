# HealthCheckData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **NullableInt32** |  | [optional] 
**InvocationTimeout** | Pointer to **NullableInt32** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewHealthCheckData

`func NewHealthCheckData() *HealthCheckData`

NewHealthCheckData instantiates a new HealthCheckData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckDataWithDefaults

`func NewHealthCheckDataWithDefaults() *HealthCheckData`

NewHealthCheckDataWithDefaults instantiates a new HealthCheckData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *HealthCheckData) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HealthCheckData) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HealthCheckData) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HealthCheckData) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *HealthCheckData) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *HealthCheckData) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetInterval

`func (o *HealthCheckData) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthCheckData) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthCheckData) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthCheckData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *HealthCheckData) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *HealthCheckData) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetInvocationTimeout

`func (o *HealthCheckData) GetInvocationTimeout() int32`

GetInvocationTimeout returns the InvocationTimeout field if non-nil, zero value otherwise.

### GetInvocationTimeoutOk

`func (o *HealthCheckData) GetInvocationTimeoutOk() (*int32, bool)`

GetInvocationTimeoutOk returns a tuple with the InvocationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationTimeout

`func (o *HealthCheckData) SetInvocationTimeout(v int32)`

SetInvocationTimeout sets InvocationTimeout field to given value.

### HasInvocationTimeout

`func (o *HealthCheckData) HasInvocationTimeout() bool`

HasInvocationTimeout returns a boolean if a field has been set.

### SetInvocationTimeoutNil

`func (o *HealthCheckData) SetInvocationTimeoutNil(b bool)`

 SetInvocationTimeoutNil sets the value for InvocationTimeout to be an explicit nil

### UnsetInvocationTimeout
`func (o *HealthCheckData) UnsetInvocationTimeout()`

UnsetInvocationTimeout ensures that no value is present for InvocationTimeout, not even an explicit nil
### GetTimeout

`func (o *HealthCheckData) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthCheckData) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthCheckData) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HealthCheckData) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *HealthCheckData) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *HealthCheckData) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


