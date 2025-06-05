# DeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]interface{}** | Additional status details | [optional] 
**Reason** | Pointer to **NullableString** | Reason for the current status | [optional] 
**Value** | Pointer to **string** | Status value | [optional] 

## Methods

### NewDeploymentStatus

`func NewDeploymentStatus() *DeploymentStatus`

NewDeploymentStatus instantiates a new DeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStatusWithDefaults

`func NewDeploymentStatusWithDefaults() *DeploymentStatus`

NewDeploymentStatusWithDefaults instantiates a new DeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DeploymentStatus) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DeploymentStatus) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DeploymentStatus) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DeploymentStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *DeploymentStatus) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *DeploymentStatus) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetReason

`func (o *DeploymentStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeploymentStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeploymentStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeploymentStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DeploymentStatus) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DeploymentStatus) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetValue

`func (o *DeploymentStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeploymentStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeploymentStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeploymentStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


