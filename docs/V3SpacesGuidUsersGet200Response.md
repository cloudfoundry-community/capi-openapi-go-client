# V3SpacesGuidUsersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Resources** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewV3SpacesGuidUsersGet200Response

`func NewV3SpacesGuidUsersGet200Response() *V3SpacesGuidUsersGet200Response`

NewV3SpacesGuidUsersGet200Response instantiates a new V3SpacesGuidUsersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SpacesGuidUsersGet200ResponseWithDefaults

`func NewV3SpacesGuidUsersGet200ResponseWithDefaults() *V3SpacesGuidUsersGet200Response`

NewV3SpacesGuidUsersGet200ResponseWithDefaults instantiates a new V3SpacesGuidUsersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *V3SpacesGuidUsersGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V3SpacesGuidUsersGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V3SpacesGuidUsersGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V3SpacesGuidUsersGet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *V3SpacesGuidUsersGet200Response) GetResources() []User`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V3SpacesGuidUsersGet200Response) GetResourcesOk() (*[]User, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V3SpacesGuidUsersGet200Response) SetResources(v []User)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V3SpacesGuidUsersGet200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


