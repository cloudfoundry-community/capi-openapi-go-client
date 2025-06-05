# V3BuildsPostRequestLifecycleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to **[]string** | List of buildpacks to use (buildpack lifecycle) | [optional] 
**Credentials** | Pointer to [**V3BuildsPostRequestLifecycleDataCredentials**](V3BuildsPostRequestLifecycleDataCredentials.md) |  | [optional] 
**Stack** | Pointer to **string** | Stack to use for staging | [optional] 

## Methods

### NewV3BuildsPostRequestLifecycleData

`func NewV3BuildsPostRequestLifecycleData() *V3BuildsPostRequestLifecycleData`

NewV3BuildsPostRequestLifecycleData instantiates a new V3BuildsPostRequestLifecycleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildsPostRequestLifecycleDataWithDefaults

`func NewV3BuildsPostRequestLifecycleDataWithDefaults() *V3BuildsPostRequestLifecycleData`

NewV3BuildsPostRequestLifecycleDataWithDefaults instantiates a new V3BuildsPostRequestLifecycleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *V3BuildsPostRequestLifecycleData) GetBuildpacks() []string`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *V3BuildsPostRequestLifecycleData) GetBuildpacksOk() (*[]string, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *V3BuildsPostRequestLifecycleData) SetBuildpacks(v []string)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *V3BuildsPostRequestLifecycleData) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetCredentials

`func (o *V3BuildsPostRequestLifecycleData) GetCredentials() V3BuildsPostRequestLifecycleDataCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *V3BuildsPostRequestLifecycleData) GetCredentialsOk() (*V3BuildsPostRequestLifecycleDataCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *V3BuildsPostRequestLifecycleData) SetCredentials(v V3BuildsPostRequestLifecycleDataCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *V3BuildsPostRequestLifecycleData) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetStack

`func (o *V3BuildsPostRequestLifecycleData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V3BuildsPostRequestLifecycleData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V3BuildsPostRequestLifecycleData) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V3BuildsPostRequestLifecycleData) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


