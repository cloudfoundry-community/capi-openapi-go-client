# V3AppsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**V3AppsGet200ResponseIncluded**](V3AppsGet200ResponseIncluded.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Resources** | Pointer to [**[]App**](App.md) |  | [optional] 

## Methods

### NewV3AppsGet200Response

`func NewV3AppsGet200Response() *V3AppsGet200Response`

NewV3AppsGet200Response instantiates a new V3AppsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGet200ResponseWithDefaults

`func NewV3AppsGet200ResponseWithDefaults() *V3AppsGet200Response`

NewV3AppsGet200ResponseWithDefaults instantiates a new V3AppsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *V3AppsGet200Response) GetIncluded() V3AppsGet200ResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *V3AppsGet200Response) GetIncludedOk() (*V3AppsGet200ResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *V3AppsGet200Response) SetIncluded(v V3AppsGet200ResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *V3AppsGet200Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetPagination

`func (o *V3AppsGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3AppsGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3AppsGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3AppsGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3AppsGet200Response) GetResources() []App`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3AppsGet200Response) GetResourcesOk() (*[]App, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3AppsGet200Response) SetResources(v []App)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3AppsGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


