# SecurityGroupGloballyEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | **bool** | Applied globally to all running apps | 
**Staging** | **bool** | Applied globally during app staging | 

## Methods

### NewSecurityGroupGloballyEnabled

`func NewSecurityGroupGloballyEnabled(running bool, staging bool, ) *SecurityGroupGloballyEnabled`

NewSecurityGroupGloballyEnabled instantiates a new SecurityGroupGloballyEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupGloballyEnabledWithDefaults

`func NewSecurityGroupGloballyEnabledWithDefaults() *SecurityGroupGloballyEnabled`

NewSecurityGroupGloballyEnabledWithDefaults instantiates a new SecurityGroupGloballyEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *SecurityGroupGloballyEnabled) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *SecurityGroupGloballyEnabled) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *SecurityGroupGloballyEnabled) SetRunning(v bool)`

SetRunning sets Running field to given value.


### GetStaging

`func (o *SecurityGroupGloballyEnabled) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *SecurityGroupGloballyEnabled) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *SecurityGroupGloballyEnabled) SetStaging(v bool)`

SetStaging sets Staging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


