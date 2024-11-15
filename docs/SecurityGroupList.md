# SecurityGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) |  | [optional] 

## Methods

### NewSecurityGroupList

`func NewSecurityGroupList() *SecurityGroupList`

NewSecurityGroupList instantiates a new SecurityGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupListWithDefaults

`func NewSecurityGroupListWithDefaults() *SecurityGroupList`

NewSecurityGroupListWithDefaults instantiates a new SecurityGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SecurityGroupList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SecurityGroupList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SecurityGroupList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SecurityGroupList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *SecurityGroupList) GetResources() []SecurityGroup`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SecurityGroupList) GetResourcesOk() (*[]SecurityGroup, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SecurityGroupList) SetResources(v []SecurityGroup)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SecurityGroupList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


