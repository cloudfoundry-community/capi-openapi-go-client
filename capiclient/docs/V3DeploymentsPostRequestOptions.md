# V3DeploymentsPostRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInFlight** | Pointer to **int32** | Maximum number of instances to update simultaneously (rolling strategy) | [optional] [default to 1]

## Methods

### NewV3DeploymentsPostRequestOptions

`func NewV3DeploymentsPostRequestOptions() *V3DeploymentsPostRequestOptions`

NewV3DeploymentsPostRequestOptions instantiates a new V3DeploymentsPostRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3DeploymentsPostRequestOptionsWithDefaults

`func NewV3DeploymentsPostRequestOptionsWithDefaults() *V3DeploymentsPostRequestOptions`

NewV3DeploymentsPostRequestOptionsWithDefaults instantiates a new V3DeploymentsPostRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInFlight

`func (o *V3DeploymentsPostRequestOptions) GetMaxInFlight() int32`

GetMaxInFlight returns the MaxInFlight field if non-nil, zero value otherwise.

### GetMaxInFlightOk

`func (o *V3DeploymentsPostRequestOptions) GetMaxInFlightOk() (*int32, bool)`

GetMaxInFlightOk returns a tuple with the MaxInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInFlight

`func (o *V3DeploymentsPostRequestOptions) SetMaxInFlight(v int32)`

SetMaxInFlight sets MaxInFlight field to given value.

### HasMaxInFlight

`func (o *V3DeploymentsPostRequestOptions) HasMaxInFlight() bool`

HasMaxInFlight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


