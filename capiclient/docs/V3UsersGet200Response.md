# V3UsersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**V3UsersGet200ResponsePagination**](V3UsersGet200ResponsePagination.md) |  | [optional] 
**Resources** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewV3UsersGet200Response

`func NewV3UsersGet200Response() *V3UsersGet200Response`

NewV3UsersGet200Response instantiates a new V3UsersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3UsersGet200ResponseWithDefaults

`func NewV3UsersGet200ResponseWithDefaults() *V3UsersGet200Response`

NewV3UsersGet200ResponseWithDefaults instantiates a new V3UsersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3UsersGet200Response) GetPagination() V3UsersGet200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3UsersGet200Response) GetPaginationOk() (*V3UsersGet200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3UsersGet200Response) SetPagination(v V3UsersGet200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3UsersGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3UsersGet200Response) GetResources() []User`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3UsersGet200Response) GetResourcesOk() (*[]User, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3UsersGet200Response) SetResources(v []User)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3UsersGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


