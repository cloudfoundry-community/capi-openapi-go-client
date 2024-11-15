# DestinationApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** |  | [optional] 
**Process** | Pointer to [**DestinationAppProcess**](DestinationAppProcess.md) |  | [optional] 

## Methods

### NewDestinationApp

`func NewDestinationApp() *DestinationApp`

NewDestinationApp instantiates a new DestinationApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationAppWithDefaults

`func NewDestinationAppWithDefaults() *DestinationApp`

NewDestinationAppWithDefaults instantiates a new DestinationApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *DestinationApp) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DestinationApp) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DestinationApp) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DestinationApp) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetProcess

`func (o *DestinationApp) GetProcess() DestinationAppProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *DestinationApp) GetProcessOk() (*DestinationAppProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *DestinationApp) SetProcess(v DestinationAppProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *DestinationApp) HasProcess() bool`

HasProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


