# JobResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | GUID of the resource | [optional] 
**Name** | Pointer to **string** | Name of the resource | [optional] 
**Type** | Pointer to **string** | Type of the resource | [optional] 

## Methods

### NewJobResource

`func NewJobResource() *JobResource`

NewJobResource instantiates a new JobResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResourceWithDefaults

`func NewJobResourceWithDefaults() *JobResource`

NewJobResourceWithDefaults instantiates a new JobResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *JobResource) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *JobResource) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *JobResource) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *JobResource) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *JobResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *JobResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


