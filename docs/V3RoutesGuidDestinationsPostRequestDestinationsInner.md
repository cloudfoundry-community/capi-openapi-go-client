# V3RoutesGuidDestinationsPostRequestDestinationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**V3RoutesGuidDestinationsPostRequestDestinationsInnerApp**](V3RoutesGuidDestinationsPostRequestDestinationsInnerApp.md) |  | 
**Port** | Pointer to **int32** | Port on the app process | [optional] 
**Protocol** | Pointer to **string** | Protocol for the destination | [optional] 
**Weight** | Pointer to **int32** | Percentage of traffic for weighted routing | [optional] 

## Methods

### NewV3RoutesGuidDestinationsPostRequestDestinationsInner

`func NewV3RoutesGuidDestinationsPostRequestDestinationsInner(app V3RoutesGuidDestinationsPostRequestDestinationsInnerApp, ) *V3RoutesGuidDestinationsPostRequestDestinationsInner`

NewV3RoutesGuidDestinationsPostRequestDestinationsInner instantiates a new V3RoutesGuidDestinationsPostRequestDestinationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3RoutesGuidDestinationsPostRequestDestinationsInnerWithDefaults

`func NewV3RoutesGuidDestinationsPostRequestDestinationsInnerWithDefaults() *V3RoutesGuidDestinationsPostRequestDestinationsInner`

NewV3RoutesGuidDestinationsPostRequestDestinationsInnerWithDefaults instantiates a new V3RoutesGuidDestinationsPostRequestDestinationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetApp() V3RoutesGuidDestinationsPostRequestDestinationsInnerApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetAppOk() (*V3RoutesGuidDestinationsPostRequestDestinationsInnerApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) SetApp(v V3RoutesGuidDestinationsPostRequestDestinationsInnerApp)`

SetApp sets App field to given value.


### GetPort

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetWeight

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V3RoutesGuidDestinationsPostRequestDestinationsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


