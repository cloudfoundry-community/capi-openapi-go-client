# EnvironmentVariableGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**EnvironmentVariableGroupLinks**](EnvironmentVariableGroupLinks.md) |  | 
**Metadata** | Pointer to [**EnvironmentVariableGroupMetadata**](EnvironmentVariableGroupMetadata.md) |  | [optional] 
**Name** | **string** | Name of the environment variable group | 
**UpdatedAt** | **time.Time** | Timestamp when the environment variable group was last updated | 
**Var** | **map[string]string** | Environment variables as key-value pairs | 

## Methods

### NewEnvironmentVariableGroup

`func NewEnvironmentVariableGroup(links EnvironmentVariableGroupLinks, name string, updatedAt time.Time, var_ map[string]string, ) *EnvironmentVariableGroup`

NewEnvironmentVariableGroup instantiates a new EnvironmentVariableGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableGroupWithDefaults

`func NewEnvironmentVariableGroupWithDefaults() *EnvironmentVariableGroup`

NewEnvironmentVariableGroupWithDefaults instantiates a new EnvironmentVariableGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentVariableGroup) GetLinks() EnvironmentVariableGroupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentVariableGroup) GetLinksOk() (*EnvironmentVariableGroupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentVariableGroup) SetLinks(v EnvironmentVariableGroupLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *EnvironmentVariableGroup) GetMetadata() EnvironmentVariableGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentVariableGroup) GetMetadataOk() (*EnvironmentVariableGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentVariableGroup) SetMetadata(v EnvironmentVariableGroupMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentVariableGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentVariableGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentVariableGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentVariableGroup) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *EnvironmentVariableGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentVariableGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentVariableGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVar

`func (o *EnvironmentVariableGroup) GetVar() map[string]string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *EnvironmentVariableGroup) GetVarOk() (*map[string]string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *EnvironmentVariableGroup) SetVar(v map[string]string)`

SetVar sets Var field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


