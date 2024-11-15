# EnvironmentVariableGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks**](V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Var** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEnvironmentVariableGroup

`func NewEnvironmentVariableGroup() *EnvironmentVariableGroup`

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

`func (o *EnvironmentVariableGroup) GetLinks() V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentVariableGroup) GetLinksOk() (*V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentVariableGroup) SetLinks(v V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentVariableGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasName

`func (o *EnvironmentVariableGroup) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasUpdatedAt

`func (o *EnvironmentVariableGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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

### HasVar

`func (o *EnvironmentVariableGroup) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


