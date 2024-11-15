# V3AppsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to **map[string]interface{}** | Environment variables for the app. | [optional] 
**Lifecycle** | Pointer to [**V3AppsPostRequestLifecycle**](V3AppsPostRequestLifecycle.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata of the app. | [optional] 
**Name** | **string** | Name of the app. | 
**Relationships** | Pointer to [**V3AppsPostRequestRelationships**](V3AppsPostRequestRelationships.md) |  | [optional] 
**SpaceGuid** | **string** | GUID of the space. | 
**State** | Pointer to **string** | Desired state of the app. | [optional] 

## Methods

### NewV3AppsPostRequest

`func NewV3AppsPostRequest(name string, spaceGuid string, ) *V3AppsPostRequest`

NewV3AppsPostRequest instantiates a new V3AppsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsPostRequestWithDefaults

`func NewV3AppsPostRequestWithDefaults() *V3AppsPostRequest`

NewV3AppsPostRequestWithDefaults instantiates a new V3AppsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *V3AppsPostRequest) GetEnvironmentVariables() map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *V3AppsPostRequest) GetEnvironmentVariablesOk() (*map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *V3AppsPostRequest) SetEnvironmentVariables(v map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *V3AppsPostRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLifecycle

`func (o *V3AppsPostRequest) GetLifecycle() V3AppsPostRequestLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3AppsPostRequest) GetLifecycleOk() (*V3AppsPostRequestLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3AppsPostRequest) SetLifecycle(v V3AppsPostRequestLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3AppsPostRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMetadata

`func (o *V3AppsPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3AppsPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3AppsPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3AppsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3AppsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3AppsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3AppsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *V3AppsPostRequest) GetRelationships() V3AppsPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3AppsPostRequest) GetRelationshipsOk() (*V3AppsPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3AppsPostRequest) SetRelationships(v V3AppsPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3AppsPostRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSpaceGuid

`func (o *V3AppsPostRequest) GetSpaceGuid() string`

GetSpaceGuid returns the SpaceGuid field if non-nil, zero value otherwise.

### GetSpaceGuidOk

`func (o *V3AppsPostRequest) GetSpaceGuidOk() (*string, bool)`

GetSpaceGuidOk returns a tuple with the SpaceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceGuid

`func (o *V3AppsPostRequest) SetSpaceGuid(v string)`

SetSpaceGuid sets SpaceGuid field to given value.


### GetState

`func (o *V3AppsPostRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V3AppsPostRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V3AppsPostRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V3AppsPostRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


