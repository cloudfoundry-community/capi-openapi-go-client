# ResourceMatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ResourceMatchObject**](ResourceMatchObject.md) |  | [optional] 

## Methods

### NewResourceMatchRequest

`func NewResourceMatchRequest() *ResourceMatchRequest`

NewResourceMatchRequest instantiates a new ResourceMatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMatchRequestWithDefaults

`func NewResourceMatchRequestWithDefaults() *ResourceMatchRequest`

NewResourceMatchRequestWithDefaults instantiates a new ResourceMatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourceMatchRequest) GetResources() []ResourceMatchObject`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceMatchRequest) GetResourcesOk() (*[]ResourceMatchObject, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceMatchRequest) SetResources(v []ResourceMatchObject)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourceMatchRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


