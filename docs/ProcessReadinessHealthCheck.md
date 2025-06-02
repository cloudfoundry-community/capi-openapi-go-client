# ProcessReadinessHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**V3ProcessesGuidPatchRequestReadinessHealthCheckData**](V3ProcessesGuidPatchRequestReadinessHealthCheckData.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewProcessReadinessHealthCheck

`func NewProcessReadinessHealthCheck() *ProcessReadinessHealthCheck`

NewProcessReadinessHealthCheck instantiates a new ProcessReadinessHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessReadinessHealthCheckWithDefaults

`func NewProcessReadinessHealthCheckWithDefaults() *ProcessReadinessHealthCheck`

NewProcessReadinessHealthCheckWithDefaults instantiates a new ProcessReadinessHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProcessReadinessHealthCheck) GetData() V3ProcessesGuidPatchRequestReadinessHealthCheckData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessReadinessHealthCheck) GetDataOk() (*V3ProcessesGuidPatchRequestReadinessHealthCheckData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessReadinessHealthCheck) SetData(v V3ProcessesGuidPatchRequestReadinessHealthCheckData)`

SetData sets Data field to given value.

### HasData

`func (o *ProcessReadinessHealthCheck) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *ProcessReadinessHealthCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessReadinessHealthCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessReadinessHealthCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProcessReadinessHealthCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


