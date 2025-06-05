# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomErrorMessage** | Pointer to **NullableString** | Custom error message shown when the feature is disabled | [optional] 
**Enabled** | **bool** | Whether the feature is enabled | 
**Links** | [**FeatureFlagLinks**](FeatureFlagLinks.md) |  | 
**Metadata** | Pointer to [**FeatureFlagMetadata**](FeatureFlagMetadata.md) |  | [optional] 
**Name** | **string** | Unique name of the feature flag | 
**UpdatedAt** | **time.Time** | Timestamp when the feature flag was last updated | 

## Methods

### NewFeatureFlag

`func NewFeatureFlag(enabled bool, links FeatureFlagLinks, name string, updatedAt time.Time, ) *FeatureFlag`

NewFeatureFlag instantiates a new FeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagWithDefaults

`func NewFeatureFlagWithDefaults() *FeatureFlag`

NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomErrorMessage

`func (o *FeatureFlag) GetCustomErrorMessage() string`

GetCustomErrorMessage returns the CustomErrorMessage field if non-nil, zero value otherwise.

### GetCustomErrorMessageOk

`func (o *FeatureFlag) GetCustomErrorMessageOk() (*string, bool)`

GetCustomErrorMessageOk returns a tuple with the CustomErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorMessage

`func (o *FeatureFlag) SetCustomErrorMessage(v string)`

SetCustomErrorMessage sets CustomErrorMessage field to given value.

### HasCustomErrorMessage

`func (o *FeatureFlag) HasCustomErrorMessage() bool`

HasCustomErrorMessage returns a boolean if a field has been set.

### SetCustomErrorMessageNil

`func (o *FeatureFlag) SetCustomErrorMessageNil(b bool)`

 SetCustomErrorMessageNil sets the value for CustomErrorMessage to be an explicit nil

### UnsetCustomErrorMessage
`func (o *FeatureFlag) UnsetCustomErrorMessage()`

UnsetCustomErrorMessage ensures that no value is present for CustomErrorMessage, not even an explicit nil
### GetEnabled

`func (o *FeatureFlag) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureFlag) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureFlag) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLinks

`func (o *FeatureFlag) GetLinks() FeatureFlagLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlag) GetLinksOk() (*FeatureFlagLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlag) SetLinks(v FeatureFlagLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *FeatureFlag) GetMetadata() FeatureFlagMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeatureFlag) GetMetadataOk() (*FeatureFlagMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeatureFlag) SetMetadata(v FeatureFlagMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeatureFlag) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *FeatureFlag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeatureFlag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeatureFlag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


