# SpaceFeatureUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether to enable or disable the feature | 
**Metadata** | Pointer to [**SpaceFeatureUpdateMetadata**](SpaceFeatureUpdateMetadata.md) |  | [optional] 

## Methods

### NewSpaceFeatureUpdate

`func NewSpaceFeatureUpdate(enabled bool, ) *SpaceFeatureUpdate`

NewSpaceFeatureUpdate instantiates a new SpaceFeatureUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceFeatureUpdateWithDefaults

`func NewSpaceFeatureUpdateWithDefaults() *SpaceFeatureUpdate`

NewSpaceFeatureUpdateWithDefaults instantiates a new SpaceFeatureUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SpaceFeatureUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpaceFeatureUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpaceFeatureUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetadata

`func (o *SpaceFeatureUpdate) GetMetadata() SpaceFeatureUpdateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpaceFeatureUpdate) GetMetadataOk() (*SpaceFeatureUpdateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpaceFeatureUpdate) SetMetadata(v SpaceFeatureUpdateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpaceFeatureUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


