# AppUsageEventProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **NullableString** | Unique identifier of the process that this event pertains to, if applicable | [optional] 
**Type** | Pointer to **NullableString** | Type of the process that this event pertains to, if applicable | [optional] 

## Methods

### NewAppUsageEventProcess

`func NewAppUsageEventProcess() *AppUsageEventProcess`

NewAppUsageEventProcess instantiates a new AppUsageEventProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventProcessWithDefaults

`func NewAppUsageEventProcessWithDefaults() *AppUsageEventProcess`

NewAppUsageEventProcessWithDefaults instantiates a new AppUsageEventProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AppUsageEventProcess) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppUsageEventProcess) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppUsageEventProcess) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppUsageEventProcess) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AppUsageEventProcess) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AppUsageEventProcess) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetType

`func (o *AppUsageEventProcess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppUsageEventProcess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppUsageEventProcess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppUsageEventProcess) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AppUsageEventProcess) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AppUsageEventProcess) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


