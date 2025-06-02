# OrganizationQuotaCreateRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalReservedPorts** | Pointer to **NullableInt32** | Total reserved ports allowed (null &#x3D; unlimited) | [optional] 
**TotalRoutes** | Pointer to **NullableInt32** | Total routes allowed (null &#x3D; unlimited) | [optional] 

## Methods

### NewOrganizationQuotaCreateRoutes

`func NewOrganizationQuotaCreateRoutes() *OrganizationQuotaCreateRoutes`

NewOrganizationQuotaCreateRoutes instantiates a new OrganizationQuotaCreateRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaCreateRoutesWithDefaults

`func NewOrganizationQuotaCreateRoutesWithDefaults() *OrganizationQuotaCreateRoutes`

NewOrganizationQuotaCreateRoutesWithDefaults instantiates a new OrganizationQuotaCreateRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalReservedPorts

`func (o *OrganizationQuotaCreateRoutes) GetTotalReservedPorts() int32`

GetTotalReservedPorts returns the TotalReservedPorts field if non-nil, zero value otherwise.

### GetTotalReservedPortsOk

`func (o *OrganizationQuotaCreateRoutes) GetTotalReservedPortsOk() (*int32, bool)`

GetTotalReservedPortsOk returns a tuple with the TotalReservedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservedPorts

`func (o *OrganizationQuotaCreateRoutes) SetTotalReservedPorts(v int32)`

SetTotalReservedPorts sets TotalReservedPorts field to given value.

### HasTotalReservedPorts

`func (o *OrganizationQuotaCreateRoutes) HasTotalReservedPorts() bool`

HasTotalReservedPorts returns a boolean if a field has been set.

### SetTotalReservedPortsNil

`func (o *OrganizationQuotaCreateRoutes) SetTotalReservedPortsNil(b bool)`

 SetTotalReservedPortsNil sets the value for TotalReservedPorts to be an explicit nil

### UnsetTotalReservedPorts
`func (o *OrganizationQuotaCreateRoutes) UnsetTotalReservedPorts()`

UnsetTotalReservedPorts ensures that no value is present for TotalReservedPorts, not even an explicit nil
### GetTotalRoutes

`func (o *OrganizationQuotaCreateRoutes) GetTotalRoutes() int32`

GetTotalRoutes returns the TotalRoutes field if non-nil, zero value otherwise.

### GetTotalRoutesOk

`func (o *OrganizationQuotaCreateRoutes) GetTotalRoutesOk() (*int32, bool)`

GetTotalRoutesOk returns a tuple with the TotalRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRoutes

`func (o *OrganizationQuotaCreateRoutes) SetTotalRoutes(v int32)`

SetTotalRoutes sets TotalRoutes field to given value.

### HasTotalRoutes

`func (o *OrganizationQuotaCreateRoutes) HasTotalRoutes() bool`

HasTotalRoutes returns a boolean if a field has been set.

### SetTotalRoutesNil

`func (o *OrganizationQuotaCreateRoutes) SetTotalRoutesNil(b bool)`

 SetTotalRoutesNil sets the value for TotalRoutes to be an explicit nil

### UnsetTotalRoutes
`func (o *OrganizationQuotaCreateRoutes) UnsetTotalRoutes()`

UnsetTotalRoutes ensures that no value is present for TotalRoutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


