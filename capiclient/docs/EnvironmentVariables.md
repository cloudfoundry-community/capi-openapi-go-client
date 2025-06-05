# EnvironmentVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**EnvironmentVariablesLinks**](EnvironmentVariablesLinks.md) |  | [optional] 
**Var** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEnvironmentVariables

`func NewEnvironmentVariables() *EnvironmentVariables`

NewEnvironmentVariables instantiates a new EnvironmentVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariablesWithDefaults

`func NewEnvironmentVariablesWithDefaults() *EnvironmentVariables`

NewEnvironmentVariablesWithDefaults instantiates a new EnvironmentVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentVariables) GetLinks() EnvironmentVariablesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentVariables) GetLinksOk() (*EnvironmentVariablesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentVariables) SetLinks(v EnvironmentVariablesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentVariables) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVar

`func (o *EnvironmentVariables) GetVar() map[string]string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *EnvironmentVariables) GetVarOk() (*map[string]string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *EnvironmentVariables) SetVar(v map[string]string)`

SetVar sets Var field to given value.

### HasVar

`func (o *EnvironmentVariables) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


