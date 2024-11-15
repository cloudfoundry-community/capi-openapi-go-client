# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**HealthCheckData**](HealthCheckData.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthCheck

`func NewHealthCheck() *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HealthCheck) GetData() HealthCheckData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HealthCheck) GetDataOk() (*HealthCheckData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HealthCheck) SetData(v HealthCheckData)`

SetData sets Data field to given value.

### HasData

`func (o *HealthCheck) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *HealthCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HealthCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


