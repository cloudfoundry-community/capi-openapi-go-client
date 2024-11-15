# ServiceBrokerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]ServiceBroker**](ServiceBroker.md) |  | [optional] 

## Methods

### NewServiceBrokerList

`func NewServiceBrokerList() *ServiceBrokerList`

NewServiceBrokerList instantiates a new ServiceBrokerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBrokerListWithDefaults

`func NewServiceBrokerListWithDefaults() *ServiceBrokerList`

NewServiceBrokerListWithDefaults instantiates a new ServiceBrokerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ServiceBrokerList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceBrokerList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceBrokerList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceBrokerList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *ServiceBrokerList) GetResources() []ServiceBroker`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServiceBrokerList) GetResourcesOk() (*[]ServiceBroker, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServiceBrokerList) SetResources(v []ServiceBroker)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServiceBrokerList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


