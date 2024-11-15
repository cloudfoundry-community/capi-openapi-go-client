# RouteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**V3RoutesGet200ResponsePagination**](V3RoutesGet200ResponsePagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Route**](Route.md) |  | [optional] 

## Methods

### NewRouteList

`func NewRouteList() *RouteList`

NewRouteList instantiates a new RouteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteListWithDefaults

`func NewRouteListWithDefaults() *RouteList`

NewRouteListWithDefaults instantiates a new RouteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RouteList) GetPagination() V3RoutesGet200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteList) GetPaginationOk() (*V3RoutesGet200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteList) SetPagination(v V3RoutesGet200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *RouteList) GetResources() []Route`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RouteList) GetResourcesOk() (*[]Route, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RouteList) SetResources(v []Route)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RouteList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


