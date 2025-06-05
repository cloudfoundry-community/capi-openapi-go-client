# ManifestDiffDiffInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the application | 
**Changes** | [**[]ManifestDiffDiffInnerChangesInner**](ManifestDiffDiffInnerChangesInner.md) | List of changes for this application | 

## Methods

### NewManifestDiffDiffInner

`func NewManifestDiffDiffInner(appName string, changes []ManifestDiffDiffInnerChangesInner, ) *ManifestDiffDiffInner`

NewManifestDiffDiffInner instantiates a new ManifestDiffDiffInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestDiffDiffInnerWithDefaults

`func NewManifestDiffDiffInnerWithDefaults() *ManifestDiffDiffInner`

NewManifestDiffDiffInnerWithDefaults instantiates a new ManifestDiffDiffInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ManifestDiffDiffInner) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ManifestDiffDiffInner) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ManifestDiffDiffInner) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetChanges

`func (o *ManifestDiffDiffInner) GetChanges() []ManifestDiffDiffInnerChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ManifestDiffDiffInner) GetChangesOk() (*[]ManifestDiffDiffInnerChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ManifestDiffDiffInner) SetChanges(v []ManifestDiffDiffInnerChangesInner)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


