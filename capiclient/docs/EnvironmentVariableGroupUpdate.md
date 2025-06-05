# EnvironmentVariableGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**EnvironmentVariableGroupUpdateMetadata**](EnvironmentVariableGroupUpdateMetadata.md) |  | [optional] 
**Var** | Pointer to **map[string]string** | Environment variables to update as key-value pairs. Set a value to null to remove the variable.  | [optional] 

## Methods

### NewEnvironmentVariableGroupUpdate

`func NewEnvironmentVariableGroupUpdate() *EnvironmentVariableGroupUpdate`

NewEnvironmentVariableGroupUpdate instantiates a new EnvironmentVariableGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableGroupUpdateWithDefaults

`func NewEnvironmentVariableGroupUpdateWithDefaults() *EnvironmentVariableGroupUpdate`

NewEnvironmentVariableGroupUpdateWithDefaults instantiates a new EnvironmentVariableGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EnvironmentVariableGroupUpdate) GetMetadata() EnvironmentVariableGroupUpdateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentVariableGroupUpdate) GetMetadataOk() (*EnvironmentVariableGroupUpdateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentVariableGroupUpdate) SetMetadata(v EnvironmentVariableGroupUpdateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentVariableGroupUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVar

`func (o *EnvironmentVariableGroupUpdate) GetVar() map[string]string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *EnvironmentVariableGroupUpdate) GetVarOk() (*map[string]string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *EnvironmentVariableGroupUpdate) SetVar(v map[string]string)`

SetVar sets Var field to given value.

### HasVar

`func (o *EnvironmentVariableGroupUpdate) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


