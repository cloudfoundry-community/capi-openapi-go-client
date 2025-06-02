# BuildLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**BuildLifecycleData**](BuildLifecycleData.md) |  | [optional] 
**Type** | Pointer to **string** | Type of lifecycle used | [optional] 

## Methods

### NewBuildLifecycle

`func NewBuildLifecycle() *BuildLifecycle`

NewBuildLifecycle instantiates a new BuildLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildLifecycleWithDefaults

`func NewBuildLifecycleWithDefaults() *BuildLifecycle`

NewBuildLifecycleWithDefaults instantiates a new BuildLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildLifecycle) GetData() BuildLifecycleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildLifecycle) GetDataOk() (*BuildLifecycleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildLifecycle) SetData(v BuildLifecycleData)`

SetData sets Data field to given value.

### HasData

`func (o *BuildLifecycle) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *BuildLifecycle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BuildLifecycle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BuildLifecycle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BuildLifecycle) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


