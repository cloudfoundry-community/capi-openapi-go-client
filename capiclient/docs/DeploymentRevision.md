# DeploymentRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | GUID of the revision being deployed | [optional] 
**Version** | Pointer to **int32** | Version number of the revision | [optional] 

## Methods

### NewDeploymentRevision

`func NewDeploymentRevision() *DeploymentRevision`

NewDeploymentRevision instantiates a new DeploymentRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRevisionWithDefaults

`func NewDeploymentRevisionWithDefaults() *DeploymentRevision`

NewDeploymentRevisionWithDefaults instantiates a new DeploymentRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *DeploymentRevision) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DeploymentRevision) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DeploymentRevision) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DeploymentRevision) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentRevision) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentRevision) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentRevision) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentRevision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


