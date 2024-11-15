# AppUsageEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**AppUsageEventListPagination**](AppUsageEventListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]AppUsageEvent**](AppUsageEvent.md) |  | [optional] 

## Methods

### NewAppUsageEventList

`func NewAppUsageEventList() *AppUsageEventList`

NewAppUsageEventList instantiates a new AppUsageEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventListWithDefaults

`func NewAppUsageEventListWithDefaults() *AppUsageEventList`

NewAppUsageEventListWithDefaults instantiates a new AppUsageEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AppUsageEventList) GetPagination() AppUsageEventListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppUsageEventList) GetPaginationOk() (*AppUsageEventListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppUsageEventList) SetPagination(v AppUsageEventListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppUsageEventList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *AppUsageEventList) GetResources() []AppUsageEvent`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AppUsageEventList) GetResourcesOk() (*[]AppUsageEvent, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AppUsageEventList) SetResources(v []AppUsageEvent)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AppUsageEventList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


