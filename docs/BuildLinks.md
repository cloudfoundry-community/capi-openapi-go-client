# BuildLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**BuildLinksApp**](BuildLinksApp.md) |  | [optional] 
**Droplet** | Pointer to [**NullableBuildLinksDroplet**](BuildLinksDroplet.md) |  | [optional] 
**Package** | Pointer to [**BuildLinksPackage**](BuildLinksPackage.md) |  | [optional] 
**Self** | Pointer to [**BuildLinksSelf**](BuildLinksSelf.md) |  | [optional] 

## Methods

### NewBuildLinks

`func NewBuildLinks() *BuildLinks`

NewBuildLinks instantiates a new BuildLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildLinksWithDefaults

`func NewBuildLinksWithDefaults() *BuildLinks`

NewBuildLinksWithDefaults instantiates a new BuildLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *BuildLinks) GetApp() BuildLinksApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *BuildLinks) GetAppOk() (*BuildLinksApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *BuildLinks) SetApp(v BuildLinksApp)`

SetApp sets App field to given value.

### HasApp

`func (o *BuildLinks) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDroplet

`func (o *BuildLinks) GetDroplet() BuildLinksDroplet`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *BuildLinks) GetDropletOk() (*BuildLinksDroplet, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *BuildLinks) SetDroplet(v BuildLinksDroplet)`

SetDroplet sets Droplet field to given value.

### HasDroplet

`func (o *BuildLinks) HasDroplet() bool`

HasDroplet returns a boolean if a field has been set.

### SetDropletNil

`func (o *BuildLinks) SetDropletNil(b bool)`

 SetDropletNil sets the value for Droplet to be an explicit nil

### UnsetDroplet
`func (o *BuildLinks) UnsetDroplet()`

UnsetDroplet ensures that no value is present for Droplet, not even an explicit nil
### GetPackage

`func (o *BuildLinks) GetPackage() BuildLinksPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *BuildLinks) GetPackageOk() (*BuildLinksPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *BuildLinks) SetPackage(v BuildLinksPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *BuildLinks) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSelf

`func (o *BuildLinks) GetSelf() BuildLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *BuildLinks) GetSelfOk() (*BuildLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *BuildLinks) SetSelf(v BuildLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *BuildLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


