# V3SecurityGroupsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Resources** | [**[]SecurityGroup**](SecurityGroup.md) |  | 

## Methods

### NewV3SecurityGroupsGet200Response

`func NewV3SecurityGroupsGet200Response(pagination Pagination, resources []SecurityGroup, ) *V3SecurityGroupsGet200Response`

NewV3SecurityGroupsGet200Response instantiates a new V3SecurityGroupsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SecurityGroupsGet200ResponseWithDefaults

`func NewV3SecurityGroupsGet200ResponseWithDefaults() *V3SecurityGroupsGet200Response`

NewV3SecurityGroupsGet200ResponseWithDefaults instantiates a new V3SecurityGroupsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3SecurityGroupsGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3SecurityGroupsGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3SecurityGroupsGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResources

`func (o *V3SecurityGroupsGet200Response) GetResources() []SecurityGroup`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3SecurityGroupsGet200Response) GetResourcesOk() (*[]SecurityGroup, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3SecurityGroupsGet200Response) SetResources(v []SecurityGroup)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


