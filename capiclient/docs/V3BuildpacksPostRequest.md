# V3BuildpacksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether buildpack is available for use | [optional] [default to true]
**Locked** | Pointer to **bool** | Whether buildpack updates are prevented | [optional] [default to false]
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | **string** | Name of the buildpack | 
**Position** | Pointer to **int32** | Priority position for automatic detection | [optional] 
**Stack** | Pointer to **string** | Stack the buildpack runs on | [optional] 

## Methods

### NewV3BuildpacksPostRequest

`func NewV3BuildpacksPostRequest(name string, ) *V3BuildpacksPostRequest`

NewV3BuildpacksPostRequest instantiates a new V3BuildpacksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildpacksPostRequestWithDefaults

`func NewV3BuildpacksPostRequestWithDefaults() *V3BuildpacksPostRequest`

NewV3BuildpacksPostRequestWithDefaults instantiates a new V3BuildpacksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *V3BuildpacksPostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V3BuildpacksPostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V3BuildpacksPostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V3BuildpacksPostRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *V3BuildpacksPostRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *V3BuildpacksPostRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *V3BuildpacksPostRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *V3BuildpacksPostRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMetadata

`func (o *V3BuildpacksPostRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3BuildpacksPostRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3BuildpacksPostRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3BuildpacksPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3BuildpacksPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3BuildpacksPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3BuildpacksPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *V3BuildpacksPostRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *V3BuildpacksPostRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *V3BuildpacksPostRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *V3BuildpacksPostRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStack

`func (o *V3BuildpacksPostRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V3BuildpacksPostRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V3BuildpacksPostRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V3BuildpacksPostRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


