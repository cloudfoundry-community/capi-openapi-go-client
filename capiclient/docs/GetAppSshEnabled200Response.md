# GetAppSshEnabled200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetAppSshEnabled200Response

`func NewGetAppSshEnabled200Response() *GetAppSshEnabled200Response`

NewGetAppSshEnabled200Response instantiates a new GetAppSshEnabled200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppSshEnabled200ResponseWithDefaults

`func NewGetAppSshEnabled200ResponseWithDefaults() *GetAppSshEnabled200Response`

NewGetAppSshEnabled200ResponseWithDefaults instantiates a new GetAppSshEnabled200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetAppSshEnabled200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAppSshEnabled200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAppSshEnabled200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetAppSshEnabled200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReason

`func (o *GetAppSshEnabled200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetAppSshEnabled200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetAppSshEnabled200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetAppSshEnabled200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GetAppSshEnabled200Response) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetAppSshEnabled200Response) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


