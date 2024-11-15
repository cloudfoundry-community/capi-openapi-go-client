# ServiceOfferingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]ServiceOffering**](ServiceOffering.md) |  | [optional] 

## Methods

### NewServiceOfferingList

`func NewServiceOfferingList() *ServiceOfferingList`

NewServiceOfferingList instantiates a new ServiceOfferingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferingListWithDefaults

`func NewServiceOfferingListWithDefaults() *ServiceOfferingList`

NewServiceOfferingListWithDefaults instantiates a new ServiceOfferingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ServiceOfferingList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceOfferingList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceOfferingList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceOfferingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *ServiceOfferingList) GetResources() []ServiceOffering`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServiceOfferingList) GetResourcesOk() (*[]ServiceOffering, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServiceOfferingList) SetResources(v []ServiceOffering)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServiceOfferingList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


