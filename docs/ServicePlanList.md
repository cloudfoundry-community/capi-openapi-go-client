# ServicePlanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]ServicePlan**](ServicePlan.md) |  | [optional] 

## Methods

### NewServicePlanList

`func NewServicePlanList() *ServicePlanList`

NewServicePlanList instantiates a new ServicePlanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanListWithDefaults

`func NewServicePlanListWithDefaults() *ServicePlanList`

NewServicePlanListWithDefaults instantiates a new ServicePlanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ServicePlanList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServicePlanList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServicePlanList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServicePlanList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *ServicePlanList) GetResources() []ServicePlan`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServicePlanList) GetResourcesOk() (*[]ServicePlan, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServicePlanList) SetResources(v []ServicePlan)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServicePlanList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

