# V3BuildsGuidPatchRequestLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildpackLifecycleBuildpacks** | Pointer to [**[]V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner**](V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner.md) |  | [optional] 
**Image** | Pointer to **string** | Resulting image reference | [optional] 

## Methods

### NewV3BuildsGuidPatchRequestLifecycleData

`func NewV3BuildsGuidPatchRequestLifecycleData() *V3BuildsGuidPatchRequestLifecycleData`

NewV3BuildsGuidPatchRequestLifecycleData instantiates a new V3BuildsGuidPatchRequestLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildsGuidPatchRequestLifecycleDataWithDefaults

`func NewV3BuildsGuidPatchRequestLifecycleDataWithDefaults() *V3BuildsGuidPatchRequestLifecycleData`

NewV3BuildsGuidPatchRequestLifecycleDataWithDefaults instantiates a new V3BuildsGuidPatchRequestLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpackLifecycleBuildpacks

`func (o *V3BuildsGuidPatchRequestLifecycleData) GetBuildpackLifecycleBuildpacks() []V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner`

GetBuildpackLifecycleBuildpacks returns the BuildpackLifecycleBuildpacks field if non-nil, zero value otherwise.

### GetBuildpackLifecycleBuildpacksOk

`func (o *V3BuildsGuidPatchRequestLifecycleData) GetBuildpackLifecycleBuildpacksOk() (*[]V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner, bool)`

GetBuildpackLifecycleBuildpacksOk returns a tuple with the BuildpackLifecycleBuildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLifecycleBuildpacks

`func (o *V3BuildsGuidPatchRequestLifecycleData) SetBuildpackLifecycleBuildpacks(v []V3BuildsGuidPatchRequestLifecycleDataBuildpackLifecycleBuildpacksInner)`

SetBuildpackLifecycleBuildpacks sets BuildpackLifecycleBuildpacks field to given value.

### HasBuildpackLifecycleBuildpacks

`func (o *V3BuildsGuidPatchRequestLifecycleData) HasBuildpackLifecycleBuildpacks() bool`

HasBuildpackLifecycleBuildpacks returns a boolean if a field has been set.

### GetImage

`func (o *V3BuildsGuidPatchRequestLifecycleData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V3BuildsGuidPatchRequestLifecycleData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V3BuildsGuidPatchRequestLifecycleData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V3BuildsGuidPatchRequestLifecycleData) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


