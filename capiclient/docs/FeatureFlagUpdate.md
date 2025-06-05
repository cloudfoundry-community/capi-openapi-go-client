# FeatureFlagUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomErrorMessage** | Pointer to **NullableString** | Custom error message shown when the feature is disabled | [optional] 
**Enabled** | Pointer to **bool** | Whether to enable or disable the feature | [optional] 
**Metadata** | Pointer to [**FeatureFlagUpdateMetadata**](FeatureFlagUpdateMetadata.md) |  | [optional] 

## Methods

### NewFeatureFlagUpdate

`func NewFeatureFlagUpdate() *FeatureFlagUpdate`

NewFeatureFlagUpdate instantiates a new FeatureFlagUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagUpdateWithDefaults

`func NewFeatureFlagUpdateWithDefaults() *FeatureFlagUpdate`

NewFeatureFlagUpdateWithDefaults instantiates a new FeatureFlagUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomErrorMessage

`func (o *FeatureFlagUpdate) GetCustomErrorMessage() string`

GetCustomErrorMessage returns the CustomErrorMessage field if non-nil, zero value otherwise.

### GetCustomErrorMessageOk

`func (o *FeatureFlagUpdate) GetCustomErrorMessageOk() (*string, bool)`

GetCustomErrorMessageOk returns a tuple with the CustomErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorMessage

`func (o *FeatureFlagUpdate) SetCustomErrorMessage(v string)`

SetCustomErrorMessage sets CustomErrorMessage field to given value.

### HasCustomErrorMessage

`func (o *FeatureFlagUpdate) HasCustomErrorMessage() bool`

HasCustomErrorMessage returns a boolean if a field has been set.

### SetCustomErrorMessageNil

`func (o *FeatureFlagUpdate) SetCustomErrorMessageNil(b bool)`

 SetCustomErrorMessageNil sets the value for CustomErrorMessage to be an explicit nil

### UnsetCustomErrorMessage
`func (o *FeatureFlagUpdate) UnsetCustomErrorMessage()`

UnsetCustomErrorMessage ensures that no value is present for CustomErrorMessage, not even an explicit nil
### GetEnabled

`func (o *FeatureFlagUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureFlagUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureFlagUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FeatureFlagUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *FeatureFlagUpdate) GetMetadata() FeatureFlagUpdateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeatureFlagUpdate) GetMetadataOk() (*FeatureFlagUpdateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeatureFlagUpdate) SetMetadata(v FeatureFlagUpdateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeatureFlagUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


