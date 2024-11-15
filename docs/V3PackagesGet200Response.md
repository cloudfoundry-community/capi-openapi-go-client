# V3PackagesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**V3PackagesGet200ResponsePagination**](V3PackagesGet200ResponsePagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Package**](Package.md) |  | [optional] 

## Methods

### NewV3PackagesGet200Response

`func NewV3PackagesGet200Response() *V3PackagesGet200Response`

NewV3PackagesGet200Response instantiates a new V3PackagesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PackagesGet200ResponseWithDefaults

`func NewV3PackagesGet200ResponseWithDefaults() *V3PackagesGet200Response`

NewV3PackagesGet200ResponseWithDefaults instantiates a new V3PackagesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3PackagesGet200Response) GetPagination() V3PackagesGet200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3PackagesGet200Response) GetPaginationOk() (*V3PackagesGet200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3PackagesGet200Response) SetPagination(v V3PackagesGet200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3PackagesGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3PackagesGet200Response) GetResources() []Package`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3PackagesGet200Response) GetResourcesOk() (*[]Package, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3PackagesGet200Response) SetResources(v []Package)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3PackagesGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


