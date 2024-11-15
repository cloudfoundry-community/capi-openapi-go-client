# V3StacksGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Stack**](Stack.md) |  | [optional] 

## Methods

### NewV3StacksGet200Response

`func NewV3StacksGet200Response() *V3StacksGet200Response`

NewV3StacksGet200Response instantiates a new V3StacksGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3StacksGet200ResponseWithDefaults

`func NewV3StacksGet200ResponseWithDefaults() *V3StacksGet200Response`

NewV3StacksGet200ResponseWithDefaults instantiates a new V3StacksGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3StacksGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3StacksGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3StacksGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3StacksGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3StacksGet200Response) GetResources() []Stack`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3StacksGet200Response) GetResourcesOk() (*[]Stack, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3StacksGet200Response) SetResources(v []Stack)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3StacksGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

