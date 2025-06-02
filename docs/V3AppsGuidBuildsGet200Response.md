# V3AppsGuidBuildsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Build**](Build.md) |  | [optional] 

## Methods

### NewV3AppsGuidBuildsGet200Response

`func NewV3AppsGuidBuildsGet200Response() *V3AppsGuidBuildsGet200Response`

NewV3AppsGuidBuildsGet200Response instantiates a new V3AppsGuidBuildsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidBuildsGet200ResponseWithDefaults

`func NewV3AppsGuidBuildsGet200ResponseWithDefaults() *V3AppsGuidBuildsGet200Response`

NewV3AppsGuidBuildsGet200ResponseWithDefaults instantiates a new V3AppsGuidBuildsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3AppsGuidBuildsGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3AppsGuidBuildsGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3AppsGuidBuildsGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3AppsGuidBuildsGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3AppsGuidBuildsGet200Response) GetResources() []Build`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3AppsGuidBuildsGet200Response) GetResourcesOk() (*[]Build, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3AppsGuidBuildsGet200Response) SetResources(v []Build)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3AppsGuidBuildsGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


