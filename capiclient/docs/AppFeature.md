# AppFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human-readable description of the feature | 
**Enabled** | **bool** | Whether the feature is enabled for this app | 
**Metadata** | Pointer to [**AppFeatureMetadata**](AppFeatureMetadata.md) |  | [optional] 
**Name** | **string** | The feature identifier | 

## Methods

### NewAppFeature

`func NewAppFeature(description string, enabled bool, name string, ) *AppFeature`

NewAppFeature instantiates a new AppFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppFeatureWithDefaults

`func NewAppFeatureWithDefaults() *AppFeature`

NewAppFeatureWithDefaults instantiates a new AppFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppFeature) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *AppFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetadata

`func (o *AppFeature) GetMetadata() AppFeatureMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppFeature) GetMetadataOk() (*AppFeatureMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppFeature) SetMetadata(v AppFeatureMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AppFeature) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AppFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppFeature) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


