# ListOrganizationQuotas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**ListOrganizationQuotas200ResponsePagination**](ListOrganizationQuotas200ResponsePagination.md) |  | [optional] 
**Resources** | Pointer to [**[]OrganizationQuota**](OrganizationQuota.md) |  | [optional] 

## Methods

### NewListOrganizationQuotas200Response

`func NewListOrganizationQuotas200Response() *ListOrganizationQuotas200Response`

NewListOrganizationQuotas200Response instantiates a new ListOrganizationQuotas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationQuotas200ResponseWithDefaults

`func NewListOrganizationQuotas200ResponseWithDefaults() *ListOrganizationQuotas200Response`

NewListOrganizationQuotas200ResponseWithDefaults instantiates a new ListOrganizationQuotas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListOrganizationQuotas200Response) GetPagination() ListOrganizationQuotas200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListOrganizationQuotas200Response) GetPaginationOk() (*ListOrganizationQuotas200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListOrganizationQuotas200Response) SetPagination(v ListOrganizationQuotas200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListOrganizationQuotas200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *ListOrganizationQuotas200Response) GetResources() []OrganizationQuota`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListOrganizationQuotas200Response) GetResourcesOk() (*[]OrganizationQuota, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListOrganizationQuotas200Response) SetResources(v []OrganizationQuota)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ListOrganizationQuotas200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


