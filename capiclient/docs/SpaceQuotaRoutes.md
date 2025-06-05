# SpaceQuotaRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalReservedPorts** | **NullableInt32** | Total reserved ports allowed (null &#x3D; unlimited) | 
**TotalRoutes** | **NullableInt32** | Total routes allowed (null &#x3D; unlimited) | 

## Methods

### NewSpaceQuotaRoutes

`func NewSpaceQuotaRoutes(totalReservedPorts NullableInt32, totalRoutes NullableInt32, ) *SpaceQuotaRoutes`

NewSpaceQuotaRoutes instantiates a new SpaceQuotaRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaRoutesWithDefaults

`func NewSpaceQuotaRoutesWithDefaults() *SpaceQuotaRoutes`

NewSpaceQuotaRoutesWithDefaults instantiates a new SpaceQuotaRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalReservedPorts

`func (o *SpaceQuotaRoutes) GetTotalReservedPorts() int32`

GetTotalReservedPorts returns the TotalReservedPorts field if non-nil, zero value otherwise.

### GetTotalReservedPortsOk

`func (o *SpaceQuotaRoutes) GetTotalReservedPortsOk() (*int32, bool)`

GetTotalReservedPortsOk returns a tuple with the TotalReservedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservedPorts

`func (o *SpaceQuotaRoutes) SetTotalReservedPorts(v int32)`

SetTotalReservedPorts sets TotalReservedPorts field to given value.


### SetTotalReservedPortsNil

`func (o *SpaceQuotaRoutes) SetTotalReservedPortsNil(b bool)`

 SetTotalReservedPortsNil sets the value for TotalReservedPorts to be an explicit nil

### UnsetTotalReservedPorts
`func (o *SpaceQuotaRoutes) UnsetTotalReservedPorts()`

UnsetTotalReservedPorts ensures that no value is present for TotalReservedPorts, not even an explicit nil
### GetTotalRoutes

`func (o *SpaceQuotaRoutes) GetTotalRoutes() int32`

GetTotalRoutes returns the TotalRoutes field if non-nil, zero value otherwise.

### GetTotalRoutesOk

`func (o *SpaceQuotaRoutes) GetTotalRoutesOk() (*int32, bool)`

GetTotalRoutesOk returns a tuple with the TotalRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRoutes

`func (o *SpaceQuotaRoutes) SetTotalRoutes(v int32)`

SetTotalRoutes sets TotalRoutes field to given value.


### SetTotalRoutesNil

`func (o *SpaceQuotaRoutes) SetTotalRoutesNil(b bool)`

 SetTotalRoutesNil sets the value for TotalRoutes to be an explicit nil

### UnsetTotalRoutes
`func (o *SpaceQuotaRoutes) UnsetTotalRoutes()`

UnsetTotalRoutes ensures that no value is present for TotalRoutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


