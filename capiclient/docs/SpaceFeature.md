# SpaceFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human-readable description of the feature | 
**Enabled** | **bool** | Whether the feature is enabled for this space | 
**Metadata** | Pointer to [**SpaceFeatureMetadata**](SpaceFeatureMetadata.md) |  | [optional] 
**Name** | **string** | The feature identifier | 

## Methods

### NewSpaceFeature

`func NewSpaceFeature(description string, enabled bool, name string, ) *SpaceFeature`

NewSpaceFeature instantiates a new SpaceFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceFeatureWithDefaults

`func NewSpaceFeatureWithDefaults() *SpaceFeature`

NewSpaceFeatureWithDefaults instantiates a new SpaceFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SpaceFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpaceFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpaceFeature) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *SpaceFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpaceFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpaceFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetadata

`func (o *SpaceFeature) GetMetadata() SpaceFeatureMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpaceFeature) GetMetadataOk() (*SpaceFeatureMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpaceFeature) SetMetadata(v SpaceFeatureMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpaceFeature) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SpaceFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceFeature) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


