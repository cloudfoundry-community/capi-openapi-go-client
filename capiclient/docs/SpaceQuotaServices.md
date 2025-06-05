# SpaceQuotaServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidServicesAllowed** | **bool** | Whether paid services are allowed | 
**TotalServiceInstances** | **NullableInt32** | Total service instances allowed (null &#x3D; unlimited) | 
**TotalServiceKeys** | **NullableInt32** | Total service keys allowed (null &#x3D; unlimited) | 

## Methods

### NewSpaceQuotaServices

`func NewSpaceQuotaServices(paidServicesAllowed bool, totalServiceInstances NullableInt32, totalServiceKeys NullableInt32, ) *SpaceQuotaServices`

NewSpaceQuotaServices instantiates a new SpaceQuotaServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaServicesWithDefaults

`func NewSpaceQuotaServicesWithDefaults() *SpaceQuotaServices`

NewSpaceQuotaServicesWithDefaults instantiates a new SpaceQuotaServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidServicesAllowed

`func (o *SpaceQuotaServices) GetPaidServicesAllowed() bool`

GetPaidServicesAllowed returns the PaidServicesAllowed field if non-nil, zero value otherwise.

### GetPaidServicesAllowedOk

`func (o *SpaceQuotaServices) GetPaidServicesAllowedOk() (*bool, bool)`

GetPaidServicesAllowedOk returns a tuple with the PaidServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServicesAllowed

`func (o *SpaceQuotaServices) SetPaidServicesAllowed(v bool)`

SetPaidServicesAllowed sets PaidServicesAllowed field to given value.


### GetTotalServiceInstances

`func (o *SpaceQuotaServices) GetTotalServiceInstances() int32`

GetTotalServiceInstances returns the TotalServiceInstances field if non-nil, zero value otherwise.

### GetTotalServiceInstancesOk

`func (o *SpaceQuotaServices) GetTotalServiceInstancesOk() (*int32, bool)`

GetTotalServiceInstancesOk returns a tuple with the TotalServiceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceInstances

`func (o *SpaceQuotaServices) SetTotalServiceInstances(v int32)`

SetTotalServiceInstances sets TotalServiceInstances field to given value.


### SetTotalServiceInstancesNil

`func (o *SpaceQuotaServices) SetTotalServiceInstancesNil(b bool)`

 SetTotalServiceInstancesNil sets the value for TotalServiceInstances to be an explicit nil

### UnsetTotalServiceInstances
`func (o *SpaceQuotaServices) UnsetTotalServiceInstances()`

UnsetTotalServiceInstances ensures that no value is present for TotalServiceInstances, not even an explicit nil
### GetTotalServiceKeys

`func (o *SpaceQuotaServices) GetTotalServiceKeys() int32`

GetTotalServiceKeys returns the TotalServiceKeys field if non-nil, zero value otherwise.

### GetTotalServiceKeysOk

`func (o *SpaceQuotaServices) GetTotalServiceKeysOk() (*int32, bool)`

GetTotalServiceKeysOk returns a tuple with the TotalServiceKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceKeys

`func (o *SpaceQuotaServices) SetTotalServiceKeys(v int32)`

SetTotalServiceKeys sets TotalServiceKeys field to given value.


### SetTotalServiceKeysNil

`func (o *SpaceQuotaServices) SetTotalServiceKeysNil(b bool)`

 SetTotalServiceKeysNil sets the value for TotalServiceKeys to be an explicit nil

### UnsetTotalServiceKeys
`func (o *SpaceQuotaServices) UnsetTotalServiceKeys()`

UnsetTotalServiceKeys ensures that no value is present for TotalServiceKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


