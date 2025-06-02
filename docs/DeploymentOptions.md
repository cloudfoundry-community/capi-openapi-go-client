# DeploymentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInFlight** | Pointer to **int32** | Maximum number of instances to update simultaneously | [optional] 

## Methods

### NewDeploymentOptions

`func NewDeploymentOptions() *DeploymentOptions`

NewDeploymentOptions instantiates a new DeploymentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentOptionsWithDefaults

`func NewDeploymentOptionsWithDefaults() *DeploymentOptions`

NewDeploymentOptionsWithDefaults instantiates a new DeploymentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInFlight

`func (o *DeploymentOptions) GetMaxInFlight() int32`

GetMaxInFlight returns the MaxInFlight field if non-nil, zero value otherwise.

### GetMaxInFlightOk

`func (o *DeploymentOptions) GetMaxInFlightOk() (*int32, bool)`

GetMaxInFlightOk returns a tuple with the MaxInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInFlight

`func (o *DeploymentOptions) SetMaxInFlight(v int32)`

SetMaxInFlight sets MaxInFlight field to given value.

### HasMaxInFlight

`func (o *DeploymentOptions) HasMaxInFlight() bool`

HasMaxInFlight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


