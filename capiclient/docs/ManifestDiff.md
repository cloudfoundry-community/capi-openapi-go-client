# ManifestDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diff** | [**[]ManifestDiffDiffInner**](ManifestDiffDiffInner.md) | List of differences between manifest and current state | 

## Methods

### NewManifestDiff

`func NewManifestDiff(diff []ManifestDiffDiffInner, ) *ManifestDiff`

NewManifestDiff instantiates a new ManifestDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestDiffWithDefaults

`func NewManifestDiffWithDefaults() *ManifestDiff`

NewManifestDiffWithDefaults instantiates a new ManifestDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiff

`func (o *ManifestDiff) GetDiff() []ManifestDiffDiffInner`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *ManifestDiff) GetDiffOk() (*[]ManifestDiffDiffInner, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *ManifestDiff) SetDiff(v []ManifestDiffDiffInner)`

SetDiff sets Diff field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


