# ProcessUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **NullableString** |  | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**ReadinessHealthCheck** | Pointer to [**ReadinessHealthCheck**](ReadinessHealthCheck.md) |  | [optional] 

## Methods

### NewProcessUpdate

`func NewProcessUpdate() *ProcessUpdate`

NewProcessUpdate instantiates a new ProcessUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessUpdateWithDefaults

`func NewProcessUpdateWithDefaults() *ProcessUpdate`

NewProcessUpdateWithDefaults instantiates a new ProcessUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ProcessUpdate) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ProcessUpdate) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ProcessUpdate) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ProcessUpdate) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *ProcessUpdate) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *ProcessUpdate) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetHealthCheck

`func (o *ProcessUpdate) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *ProcessUpdate) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *ProcessUpdate) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *ProcessUpdate) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetMetadata

`func (o *ProcessUpdate) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProcessUpdate) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProcessUpdate) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProcessUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReadinessHealthCheck

`func (o *ProcessUpdate) GetReadinessHealthCheck() ReadinessHealthCheck`

GetReadinessHealthCheck returns the ReadinessHealthCheck field if non-nil, zero value otherwise.

### GetReadinessHealthCheckOk

`func (o *ProcessUpdate) GetReadinessHealthCheckOk() (*ReadinessHealthCheck, bool)`

GetReadinessHealthCheckOk returns a tuple with the ReadinessHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessHealthCheck

`func (o *ProcessUpdate) SetReadinessHealthCheck(v ReadinessHealthCheck)`

SetReadinessHealthCheck sets ReadinessHealthCheck field to given value.

### HasReadinessHealthCheck

`func (o *ProcessUpdate) HasReadinessHealthCheck() bool`

HasReadinessHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


