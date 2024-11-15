# SecurityGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GloballyEnabled** | Pointer to [**SecurityGroupGloballyEnabled**](SecurityGroupGloballyEnabled.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 

## Methods

### NewSecurityGroupCreate

`func NewSecurityGroupCreate() *SecurityGroupCreate`

NewSecurityGroupCreate instantiates a new SecurityGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupCreateWithDefaults

`func NewSecurityGroupCreateWithDefaults() *SecurityGroupCreate`

NewSecurityGroupCreateWithDefaults instantiates a new SecurityGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGloballyEnabled

`func (o *SecurityGroupCreate) GetGloballyEnabled() SecurityGroupGloballyEnabled`

GetGloballyEnabled returns the GloballyEnabled field if non-nil, zero value otherwise.

### GetGloballyEnabledOk

`func (o *SecurityGroupCreate) GetGloballyEnabledOk() (*SecurityGroupGloballyEnabled, bool)`

GetGloballyEnabledOk returns a tuple with the GloballyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyEnabled

`func (o *SecurityGroupCreate) SetGloballyEnabled(v SecurityGroupGloballyEnabled)`

SetGloballyEnabled sets GloballyEnabled field to given value.

### HasGloballyEnabled

`func (o *SecurityGroupCreate) HasGloballyEnabled() bool`

HasGloballyEnabled returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *SecurityGroupCreate) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupCreate) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupCreate) SetRules(v []Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroupCreate) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


