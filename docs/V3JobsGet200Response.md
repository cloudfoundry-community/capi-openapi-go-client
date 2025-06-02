# V3JobsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Resources** | [**[]Job**](Job.md) |  | 

## Methods

### NewV3JobsGet200Response

`func NewV3JobsGet200Response(pagination Pagination, resources []Job, ) *V3JobsGet200Response`

NewV3JobsGet200Response instantiates a new V3JobsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3JobsGet200ResponseWithDefaults

`func NewV3JobsGet200ResponseWithDefaults() *V3JobsGet200Response`

NewV3JobsGet200ResponseWithDefaults instantiates a new V3JobsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3JobsGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3JobsGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3JobsGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResources

`func (o *V3JobsGet200Response) GetResources() []Job`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3JobsGet200Response) GetResourcesOk() (*[]Job, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3JobsGet200Response) SetResources(v []Job)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


