# ResourceMatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ResourceMatchObject**](ResourceMatchObject.md) |  | [optional] 

## Methods

### NewResourceMatchResponse

`func NewResourceMatchResponse() *ResourceMatchResponse`

NewResourceMatchResponse instantiates a new ResourceMatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMatchResponseWithDefaults

`func NewResourceMatchResponseWithDefaults() *ResourceMatchResponse`

NewResourceMatchResponseWithDefaults instantiates a new ResourceMatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourceMatchResponse) GetResources() []ResourceMatchObject`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceMatchResponse) GetResourcesOk() (*[]ResourceMatchObject, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceMatchResponse) SetResources(v []ResourceMatchObject)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourceMatchResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


