# SecurityGroupCreateGloballyEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to **bool** | Apply globally to all running apps | [optional] [default to false]
**Staging** | Pointer to **bool** | Apply globally during app staging | [optional] [default to false]

## Methods

### NewSecurityGroupCreateGloballyEnabled

`func NewSecurityGroupCreateGloballyEnabled() *SecurityGroupCreateGloballyEnabled`

NewSecurityGroupCreateGloballyEnabled instantiates a new SecurityGroupCreateGloballyEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupCreateGloballyEnabledWithDefaults

`func NewSecurityGroupCreateGloballyEnabledWithDefaults() *SecurityGroupCreateGloballyEnabled`

NewSecurityGroupCreateGloballyEnabledWithDefaults instantiates a new SecurityGroupCreateGloballyEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *SecurityGroupCreateGloballyEnabled) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *SecurityGroupCreateGloballyEnabled) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *SecurityGroupCreateGloballyEnabled) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *SecurityGroupCreateGloballyEnabled) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetStaging

`func (o *SecurityGroupCreateGloballyEnabled) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *SecurityGroupCreateGloballyEnabled) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *SecurityGroupCreateGloballyEnabled) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *SecurityGroupCreateGloballyEnabled) HasStaging() bool`

HasStaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


