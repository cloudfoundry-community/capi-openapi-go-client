# RolesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewRolesList

`func NewRolesList() *RolesList`

NewRolesList instantiates a new RolesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesListWithDefaults

`func NewRolesListWithDefaults() *RolesList`

NewRolesListWithDefaults instantiates a new RolesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RolesList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RolesList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RolesList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RolesList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *RolesList) GetResources() []Role`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RolesList) GetResourcesOk() (*[]Role, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RolesList) SetResources(v []Role)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RolesList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


