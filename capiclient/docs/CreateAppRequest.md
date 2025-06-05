# CreateAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to **map[string]string** | Environment variables for the app | [optional] 
**Lifecycle** | Pointer to [**CreateAppRequestLifecycle**](CreateAppRequestLifecycle.md) |  | [optional] 
**Metadata** | Pointer to [**CreateAppRequestMetadata**](CreateAppRequestMetadata.md) |  | [optional] 
**Name** | **string** | Name of the app (unique within space) | 
**Relationships** | [**CreateAppRequestRelationships**](CreateAppRequestRelationships.md) |  | 
**State** | Pointer to **string** | Initial desired state of the app | [optional] [default to "STOPPED"]

## Methods

### NewCreateAppRequest

`func NewCreateAppRequest(name string, relationships CreateAppRequestRelationships, ) *CreateAppRequest`

NewCreateAppRequest instantiates a new CreateAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppRequestWithDefaults

`func NewCreateAppRequestWithDefaults() *CreateAppRequest`

NewCreateAppRequestWithDefaults instantiates a new CreateAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *CreateAppRequest) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *CreateAppRequest) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *CreateAppRequest) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *CreateAppRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLifecycle

`func (o *CreateAppRequest) GetLifecycle() CreateAppRequestLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *CreateAppRequest) GetLifecycleOk() (*CreateAppRequestLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *CreateAppRequest) SetLifecycle(v CreateAppRequestLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *CreateAppRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateAppRequest) GetMetadata() CreateAppRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAppRequest) GetMetadataOk() (*CreateAppRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAppRequest) SetMetadata(v CreateAppRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateAppRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateAppRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *CreateAppRequest) GetRelationships() CreateAppRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CreateAppRequest) GetRelationshipsOk() (*CreateAppRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CreateAppRequest) SetRelationships(v CreateAppRequestRelationships)`

SetRelationships sets Relationships field to given value.


### GetState

`func (o *CreateAppRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateAppRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateAppRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateAppRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


