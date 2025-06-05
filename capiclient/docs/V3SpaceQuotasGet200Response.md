# V3SpaceQuotasGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**V3SpaceQuotasGet200ResponseIncluded**](V3SpaceQuotasGet200ResponseIncluded.md) |  | [optional] 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Resources** | [**[]SpaceQuota**](SpaceQuota.md) |  | 

## Methods

### NewV3SpaceQuotasGet200Response

`func NewV3SpaceQuotasGet200Response(pagination Pagination, resources []SpaceQuota, ) *V3SpaceQuotasGet200Response`

NewV3SpaceQuotasGet200Response instantiates a new V3SpaceQuotasGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SpaceQuotasGet200ResponseWithDefaults

`func NewV3SpaceQuotasGet200ResponseWithDefaults() *V3SpaceQuotasGet200Response`

NewV3SpaceQuotasGet200ResponseWithDefaults instantiates a new V3SpaceQuotasGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *V3SpaceQuotasGet200Response) GetIncluded() V3SpaceQuotasGet200ResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *V3SpaceQuotasGet200Response) GetIncludedOk() (*V3SpaceQuotasGet200ResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *V3SpaceQuotasGet200Response) SetIncluded(v V3SpaceQuotasGet200ResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *V3SpaceQuotasGet200Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetPagination

`func (o *V3SpaceQuotasGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3SpaceQuotasGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3SpaceQuotasGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResources

`func (o *V3SpaceQuotasGet200Response) GetResources() []SpaceQuota`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3SpaceQuotasGet200Response) GetResourcesOk() (*[]SpaceQuota, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3SpaceQuotasGet200Response) SetResources(v []SpaceQuota)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


