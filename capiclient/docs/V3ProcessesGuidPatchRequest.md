# V3ProcessesGuidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **NullableString** | Start command for the process | [optional] 
**HealthCheck** | Pointer to [**V3ProcessesGuidPatchRequestHealthCheck**](V3ProcessesGuidPatchRequestHealthCheck.md) |  | [optional] 
**Metadata** | Pointer to [**UpdateAppRequestMetadata**](UpdateAppRequestMetadata.md) |  | [optional] 
**ReadinessHealthCheck** | Pointer to [**V3ProcessesGuidPatchRequestReadinessHealthCheck**](V3ProcessesGuidPatchRequestReadinessHealthCheck.md) |  | [optional] 

## Methods

### NewV3ProcessesGuidPatchRequest

`func NewV3ProcessesGuidPatchRequest() *V3ProcessesGuidPatchRequest`

NewV3ProcessesGuidPatchRequest instantiates a new V3ProcessesGuidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ProcessesGuidPatchRequestWithDefaults

`func NewV3ProcessesGuidPatchRequestWithDefaults() *V3ProcessesGuidPatchRequest`

NewV3ProcessesGuidPatchRequestWithDefaults instantiates a new V3ProcessesGuidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *V3ProcessesGuidPatchRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V3ProcessesGuidPatchRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V3ProcessesGuidPatchRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V3ProcessesGuidPatchRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *V3ProcessesGuidPatchRequest) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *V3ProcessesGuidPatchRequest) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetHealthCheck

`func (o *V3ProcessesGuidPatchRequest) GetHealthCheck() V3ProcessesGuidPatchRequestHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *V3ProcessesGuidPatchRequest) GetHealthCheckOk() (*V3ProcessesGuidPatchRequestHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *V3ProcessesGuidPatchRequest) SetHealthCheck(v V3ProcessesGuidPatchRequestHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *V3ProcessesGuidPatchRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetMetadata

`func (o *V3ProcessesGuidPatchRequest) GetMetadata() UpdateAppRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3ProcessesGuidPatchRequest) GetMetadataOk() (*UpdateAppRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3ProcessesGuidPatchRequest) SetMetadata(v UpdateAppRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3ProcessesGuidPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReadinessHealthCheck

`func (o *V3ProcessesGuidPatchRequest) GetReadinessHealthCheck() V3ProcessesGuidPatchRequestReadinessHealthCheck`

GetReadinessHealthCheck returns the ReadinessHealthCheck field if non-nil, zero value otherwise.

### GetReadinessHealthCheckOk

`func (o *V3ProcessesGuidPatchRequest) GetReadinessHealthCheckOk() (*V3ProcessesGuidPatchRequestReadinessHealthCheck, bool)`

GetReadinessHealthCheckOk returns a tuple with the ReadinessHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessHealthCheck

`func (o *V3ProcessesGuidPatchRequest) SetReadinessHealthCheck(v V3ProcessesGuidPatchRequestReadinessHealthCheck)`

SetReadinessHealthCheck sets ReadinessHealthCheck field to given value.

### HasReadinessHealthCheck

`func (o *V3ProcessesGuidPatchRequest) HasReadinessHealthCheck() bool`

HasReadinessHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


