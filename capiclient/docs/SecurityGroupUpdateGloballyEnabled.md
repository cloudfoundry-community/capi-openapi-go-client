# SecurityGroupUpdateGloballyEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to **bool** | Apply globally to all running apps | [optional] 
**Staging** | Pointer to **bool** | Apply globally during app staging | [optional] 

## Methods

### NewSecurityGroupUpdateGloballyEnabled

`func NewSecurityGroupUpdateGloballyEnabled() *SecurityGroupUpdateGloballyEnabled`

NewSecurityGroupUpdateGloballyEnabled instantiates a new SecurityGroupUpdateGloballyEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupUpdateGloballyEnabledWithDefaults

`func NewSecurityGroupUpdateGloballyEnabledWithDefaults() *SecurityGroupUpdateGloballyEnabled`

NewSecurityGroupUpdateGloballyEnabledWithDefaults instantiates a new SecurityGroupUpdateGloballyEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *SecurityGroupUpdateGloballyEnabled) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *SecurityGroupUpdateGloballyEnabled) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *SecurityGroupUpdateGloballyEnabled) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *SecurityGroupUpdateGloballyEnabled) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetStaging

`func (o *SecurityGroupUpdateGloballyEnabled) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *SecurityGroupUpdateGloballyEnabled) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *SecurityGroupUpdateGloballyEnabled) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *SecurityGroupUpdateGloballyEnabled) HasStaging() bool`

HasStaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


