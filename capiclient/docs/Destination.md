# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**DestinationApp**](DestinationApp.md) |  | 
**Guid** | **string** | Unique identifier for the destination | 
**Port** | **NullableInt32** | Port on the app process | 
**Protocol** | **NullableString** | Protocol for the destination | 
**Weight** | **NullableInt32** | Percentage of traffic for weighted routing | 

## Methods

### NewDestination

`func NewDestination(app DestinationApp, guid string, port NullableInt32, protocol NullableString, weight NullableInt32, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *Destination) GetApp() DestinationApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Destination) GetAppOk() (*DestinationApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Destination) SetApp(v DestinationApp)`

SetApp sets App field to given value.


### GetGuid

`func (o *Destination) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Destination) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Destination) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetPort

`func (o *Destination) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Destination) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Destination) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *Destination) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Destination) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *Destination) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Destination) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Destination) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *Destination) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *Destination) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetWeight

`func (o *Destination) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Destination) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Destination) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### SetWeightNil

`func (o *Destination) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Destination) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


