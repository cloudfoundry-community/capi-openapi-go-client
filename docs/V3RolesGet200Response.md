# V3RolesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**V3RolesGet200ResponseIncluded**](V3RolesGet200ResponseIncluded.md) |  | [optional] 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Resources** | [**[]Role**](Role.md) |  | 

## Methods

### NewV3RolesGet200Response

`func NewV3RolesGet200Response(pagination Pagination, resources []Role, ) *V3RolesGet200Response`

NewV3RolesGet200Response instantiates a new V3RolesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3RolesGet200ResponseWithDefaults

`func NewV3RolesGet200ResponseWithDefaults() *V3RolesGet200Response`

NewV3RolesGet200ResponseWithDefaults instantiates a new V3RolesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *V3RolesGet200Response) GetIncluded() V3RolesGet200ResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *V3RolesGet200Response) GetIncludedOk() (*V3RolesGet200ResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *V3RolesGet200Response) SetIncluded(v V3RolesGet200ResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *V3RolesGet200Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetPagination

`func (o *V3RolesGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3RolesGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3RolesGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResources

`func (o *V3RolesGet200Response) GetResources() []Role`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3RolesGet200Response) GetResourcesOk() (*[]Role, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3RolesGet200Response) SetResources(v []Role)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


