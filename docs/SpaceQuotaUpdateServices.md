# SpaceQuotaUpdateServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidServicesAllowed** | Pointer to **bool** | Whether paid services are allowed | [optional] 
**TotalServiceInstances** | Pointer to **NullableInt32** | Total service instances allowed (null &#x3D; unlimited) | [optional] 
**TotalServiceKeys** | Pointer to **NullableInt32** | Total service keys allowed (null &#x3D; unlimited) | [optional] 

## Methods

### NewSpaceQuotaUpdateServices

`func NewSpaceQuotaUpdateServices() *SpaceQuotaUpdateServices`

NewSpaceQuotaUpdateServices instantiates a new SpaceQuotaUpdateServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaUpdateServicesWithDefaults

`func NewSpaceQuotaUpdateServicesWithDefaults() *SpaceQuotaUpdateServices`

NewSpaceQuotaUpdateServicesWithDefaults instantiates a new SpaceQuotaUpdateServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidServicesAllowed

`func (o *SpaceQuotaUpdateServices) GetPaidServicesAllowed() bool`

GetPaidServicesAllowed returns the PaidServicesAllowed field if non-nil, zero value otherwise.

### GetPaidServicesAllowedOk

`func (o *SpaceQuotaUpdateServices) GetPaidServicesAllowedOk() (*bool, bool)`

GetPaidServicesAllowedOk returns a tuple with the PaidServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServicesAllowed

`func (o *SpaceQuotaUpdateServices) SetPaidServicesAllowed(v bool)`

SetPaidServicesAllowed sets PaidServicesAllowed field to given value.

### HasPaidServicesAllowed

`func (o *SpaceQuotaUpdateServices) HasPaidServicesAllowed() bool`

HasPaidServicesAllowed returns a boolean if a field has been set.

### GetTotalServiceInstances

`func (o *SpaceQuotaUpdateServices) GetTotalServiceInstances() int32`

GetTotalServiceInstances returns the TotalServiceInstances field if non-nil, zero value otherwise.

### GetTotalServiceInstancesOk

`func (o *SpaceQuotaUpdateServices) GetTotalServiceInstancesOk() (*int32, bool)`

GetTotalServiceInstancesOk returns a tuple with the TotalServiceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceInstances

`func (o *SpaceQuotaUpdateServices) SetTotalServiceInstances(v int32)`

SetTotalServiceInstances sets TotalServiceInstances field to given value.

### HasTotalServiceInstances

`func (o *SpaceQuotaUpdateServices) HasTotalServiceInstances() bool`

HasTotalServiceInstances returns a boolean if a field has been set.

### SetTotalServiceInstancesNil

`func (o *SpaceQuotaUpdateServices) SetTotalServiceInstancesNil(b bool)`

 SetTotalServiceInstancesNil sets the value for TotalServiceInstances to be an explicit nil

### UnsetTotalServiceInstances
`func (o *SpaceQuotaUpdateServices) UnsetTotalServiceInstances()`

UnsetTotalServiceInstances ensures that no value is present for TotalServiceInstances, not even an explicit nil
### GetTotalServiceKeys

`func (o *SpaceQuotaUpdateServices) GetTotalServiceKeys() int32`

GetTotalServiceKeys returns the TotalServiceKeys field if non-nil, zero value otherwise.

### GetTotalServiceKeysOk

`func (o *SpaceQuotaUpdateServices) GetTotalServiceKeysOk() (*int32, bool)`

GetTotalServiceKeysOk returns a tuple with the TotalServiceKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceKeys

`func (o *SpaceQuotaUpdateServices) SetTotalServiceKeys(v int32)`

SetTotalServiceKeys sets TotalServiceKeys field to given value.

### HasTotalServiceKeys

`func (o *SpaceQuotaUpdateServices) HasTotalServiceKeys() bool`

HasTotalServiceKeys returns a boolean if a field has been set.

### SetTotalServiceKeysNil

`func (o *SpaceQuotaUpdateServices) SetTotalServiceKeysNil(b bool)`

 SetTotalServiceKeysNil sets the value for TotalServiceKeys to be an explicit nil

### UnsetTotalServiceKeys
`func (o *SpaceQuotaUpdateServices) UnsetTotalServiceKeys()`

UnsetTotalServiceKeys ensures that no value is present for TotalServiceKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


