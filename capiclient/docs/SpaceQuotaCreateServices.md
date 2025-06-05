# SpaceQuotaCreateServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidServicesAllowed** | Pointer to **bool** | Whether paid services are allowed | [optional] [default to true]
**TotalServiceInstances** | Pointer to **NullableInt32** | Total service instances allowed (null &#x3D; unlimited) | [optional] 
**TotalServiceKeys** | Pointer to **NullableInt32** | Total service keys allowed (null &#x3D; unlimited) | [optional] 

## Methods

### NewSpaceQuotaCreateServices

`func NewSpaceQuotaCreateServices() *SpaceQuotaCreateServices`

NewSpaceQuotaCreateServices instantiates a new SpaceQuotaCreateServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaCreateServicesWithDefaults

`func NewSpaceQuotaCreateServicesWithDefaults() *SpaceQuotaCreateServices`

NewSpaceQuotaCreateServicesWithDefaults instantiates a new SpaceQuotaCreateServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidServicesAllowed

`func (o *SpaceQuotaCreateServices) GetPaidServicesAllowed() bool`

GetPaidServicesAllowed returns the PaidServicesAllowed field if non-nil, zero value otherwise.

### GetPaidServicesAllowedOk

`func (o *SpaceQuotaCreateServices) GetPaidServicesAllowedOk() (*bool, bool)`

GetPaidServicesAllowedOk returns a tuple with the PaidServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServicesAllowed

`func (o *SpaceQuotaCreateServices) SetPaidServicesAllowed(v bool)`

SetPaidServicesAllowed sets PaidServicesAllowed field to given value.

### HasPaidServicesAllowed

`func (o *SpaceQuotaCreateServices) HasPaidServicesAllowed() bool`

HasPaidServicesAllowed returns a boolean if a field has been set.

### GetTotalServiceInstances

`func (o *SpaceQuotaCreateServices) GetTotalServiceInstances() int32`

GetTotalServiceInstances returns the TotalServiceInstances field if non-nil, zero value otherwise.

### GetTotalServiceInstancesOk

`func (o *SpaceQuotaCreateServices) GetTotalServiceInstancesOk() (*int32, bool)`

GetTotalServiceInstancesOk returns a tuple with the TotalServiceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceInstances

`func (o *SpaceQuotaCreateServices) SetTotalServiceInstances(v int32)`

SetTotalServiceInstances sets TotalServiceInstances field to given value.

### HasTotalServiceInstances

`func (o *SpaceQuotaCreateServices) HasTotalServiceInstances() bool`

HasTotalServiceInstances returns a boolean if a field has been set.

### SetTotalServiceInstancesNil

`func (o *SpaceQuotaCreateServices) SetTotalServiceInstancesNil(b bool)`

 SetTotalServiceInstancesNil sets the value for TotalServiceInstances to be an explicit nil

### UnsetTotalServiceInstances
`func (o *SpaceQuotaCreateServices) UnsetTotalServiceInstances()`

UnsetTotalServiceInstances ensures that no value is present for TotalServiceInstances, not even an explicit nil
### GetTotalServiceKeys

`func (o *SpaceQuotaCreateServices) GetTotalServiceKeys() int32`

GetTotalServiceKeys returns the TotalServiceKeys field if non-nil, zero value otherwise.

### GetTotalServiceKeysOk

`func (o *SpaceQuotaCreateServices) GetTotalServiceKeysOk() (*int32, bool)`

GetTotalServiceKeysOk returns a tuple with the TotalServiceKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceKeys

`func (o *SpaceQuotaCreateServices) SetTotalServiceKeys(v int32)`

SetTotalServiceKeys sets TotalServiceKeys field to given value.

### HasTotalServiceKeys

`func (o *SpaceQuotaCreateServices) HasTotalServiceKeys() bool`

HasTotalServiceKeys returns a boolean if a field has been set.

### SetTotalServiceKeysNil

`func (o *SpaceQuotaCreateServices) SetTotalServiceKeysNil(b bool)`

 SetTotalServiceKeysNil sets the value for TotalServiceKeys to be an explicit nil

### UnsetTotalServiceKeys
`func (o *SpaceQuotaCreateServices) UnsetTotalServiceKeys()`

UnsetTotalServiceKeys ensures that no value is present for TotalServiceKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


