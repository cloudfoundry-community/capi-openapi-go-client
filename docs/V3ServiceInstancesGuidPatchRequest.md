# V3ServiceInstancesGuidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceInfo** | Pointer to [**Get200ResponseLinksCloudControllerV2Meta**](Get200ResponseLinksCloudControllerV2Meta.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**ManagedServiceInstanceUpdateRelationships**](ManagedServiceInstanceUpdateRelationships.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**RouteServiceUrl** | Pointer to **string** |  | [optional] 
**SyslogDrainUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewV3ServiceInstancesGuidPatchRequest

`func NewV3ServiceInstancesGuidPatchRequest() *V3ServiceInstancesGuidPatchRequest`

NewV3ServiceInstancesGuidPatchRequest instantiates a new V3ServiceInstancesGuidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServiceInstancesGuidPatchRequestWithDefaults

`func NewV3ServiceInstancesGuidPatchRequestWithDefaults() *V3ServiceInstancesGuidPatchRequest`

NewV3ServiceInstancesGuidPatchRequestWithDefaults instantiates a new V3ServiceInstancesGuidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceInfo

`func (o *V3ServiceInstancesGuidPatchRequest) GetMaintenanceInfo() Get200ResponseLinksCloudControllerV2Meta`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetMaintenanceInfoOk() (*Get200ResponseLinksCloudControllerV2Meta, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *V3ServiceInstancesGuidPatchRequest) SetMaintenanceInfo(v Get200ResponseLinksCloudControllerV2Meta)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *V3ServiceInstancesGuidPatchRequest) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *V3ServiceInstancesGuidPatchRequest) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3ServiceInstancesGuidPatchRequest) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3ServiceInstancesGuidPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3ServiceInstancesGuidPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3ServiceInstancesGuidPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3ServiceInstancesGuidPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *V3ServiceInstancesGuidPatchRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V3ServiceInstancesGuidPatchRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V3ServiceInstancesGuidPatchRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRelationships

`func (o *V3ServiceInstancesGuidPatchRequest) GetRelationships() ManagedServiceInstanceUpdateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetRelationshipsOk() (*ManagedServiceInstanceUpdateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3ServiceInstancesGuidPatchRequest) SetRelationships(v ManagedServiceInstanceUpdateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3ServiceInstancesGuidPatchRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *V3ServiceInstancesGuidPatchRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V3ServiceInstancesGuidPatchRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V3ServiceInstancesGuidPatchRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCredentials

`func (o *V3ServiceInstancesGuidPatchRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *V3ServiceInstancesGuidPatchRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *V3ServiceInstancesGuidPatchRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRouteServiceUrl

`func (o *V3ServiceInstancesGuidPatchRequest) GetRouteServiceUrl() string`

GetRouteServiceUrl returns the RouteServiceUrl field if non-nil, zero value otherwise.

### GetRouteServiceUrlOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetRouteServiceUrlOk() (*string, bool)`

GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteServiceUrl

`func (o *V3ServiceInstancesGuidPatchRequest) SetRouteServiceUrl(v string)`

SetRouteServiceUrl sets RouteServiceUrl field to given value.

### HasRouteServiceUrl

`func (o *V3ServiceInstancesGuidPatchRequest) HasRouteServiceUrl() bool`

HasRouteServiceUrl returns a boolean if a field has been set.

### GetSyslogDrainUrl

`func (o *V3ServiceInstancesGuidPatchRequest) GetSyslogDrainUrl() string`

GetSyslogDrainUrl returns the SyslogDrainUrl field if non-nil, zero value otherwise.

### GetSyslogDrainUrlOk

`func (o *V3ServiceInstancesGuidPatchRequest) GetSyslogDrainUrlOk() (*string, bool)`

GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDrainUrl

`func (o *V3ServiceInstancesGuidPatchRequest) SetSyslogDrainUrl(v string)`

SetSyslogDrainUrl sets SyslogDrainUrl field to given value.

### HasSyslogDrainUrl

`func (o *V3ServiceInstancesGuidPatchRequest) HasSyslogDrainUrl() bool`

HasSyslogDrainUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


