# SecurityGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GloballyEnabled** | Pointer to [**SecurityGroupUpdateGloballyEnabled**](SecurityGroupUpdateGloballyEnabled.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Rules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 

## Methods

### NewSecurityGroupUpdate

`func NewSecurityGroupUpdate() *SecurityGroupUpdate`

NewSecurityGroupUpdate instantiates a new SecurityGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupUpdateWithDefaults

`func NewSecurityGroupUpdateWithDefaults() *SecurityGroupUpdate`

NewSecurityGroupUpdateWithDefaults instantiates a new SecurityGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGloballyEnabled

`func (o *SecurityGroupUpdate) GetGloballyEnabled() SecurityGroupUpdateGloballyEnabled`

GetGloballyEnabled returns the GloballyEnabled field if non-nil, zero value otherwise.

### GetGloballyEnabledOk

`func (o *SecurityGroupUpdate) GetGloballyEnabledOk() (*SecurityGroupUpdateGloballyEnabled, bool)`

GetGloballyEnabledOk returns a tuple with the GloballyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyEnabled

`func (o *SecurityGroupUpdate) SetGloballyEnabled(v SecurityGroupUpdateGloballyEnabled)`

SetGloballyEnabled sets GloballyEnabled field to given value.

### HasGloballyEnabled

`func (o *SecurityGroupUpdate) HasGloballyEnabled() bool`

HasGloballyEnabled returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityGroupUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityGroupUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRules

`func (o *SecurityGroupUpdate) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupUpdate) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupUpdate) SetRules(v []Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroupUpdate) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *SecurityGroupUpdate) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *SecurityGroupUpdate) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


