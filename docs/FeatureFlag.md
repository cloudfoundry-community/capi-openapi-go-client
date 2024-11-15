# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomErrorMessage** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks**](V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFeatureFlag

`func NewFeatureFlag() *FeatureFlag`

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

### HasEnabled

`func (o *FeatureFlag) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLinks

`func (o *FeatureFlag) GetLinks() V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlag) GetLinksOk() (*V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlag) SetLinks(v V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FeatureFlag) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasName

`func (o *FeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasUpdatedAt

`func (o *FeatureFlag) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


