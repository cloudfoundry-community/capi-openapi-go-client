# ManagedServiceInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceInfo** | Pointer to [**Get200ResponseLinksCloudControllerV2Meta**](Get200ResponseLinksCloudControllerV2Meta.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**ManagedServiceInstanceUpdateRelationships**](ManagedServiceInstanceUpdateRelationships.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManagedServiceInstanceUpdate

`func NewManagedServiceInstanceUpdate() *ManagedServiceInstanceUpdate`

NewManagedServiceInstanceUpdate instantiates a new ManagedServiceInstanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedServiceInstanceUpdateWithDefaults

`func NewManagedServiceInstanceUpdateWithDefaults() *ManagedServiceInstanceUpdate`

NewManagedServiceInstanceUpdateWithDefaults instantiates a new ManagedServiceInstanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceInfo

`func (o *ManagedServiceInstanceUpdate) GetMaintenanceInfo() Get200ResponseLinksCloudControllerV2Meta`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ManagedServiceInstanceUpdate) GetMaintenanceInfoOk() (*Get200ResponseLinksCloudControllerV2Meta, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ManagedServiceInstanceUpdate) SetMaintenanceInfo(v Get200ResponseLinksCloudControllerV2Meta)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ManagedServiceInstanceUpdate) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagedServiceInstanceUpdate) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagedServiceInstanceUpdate) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagedServiceInstanceUpdate) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagedServiceInstanceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ManagedServiceInstanceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedServiceInstanceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedServiceInstanceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedServiceInstanceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ManagedServiceInstanceUpdate) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ManagedServiceInstanceUpdate) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ManagedServiceInstanceUpdate) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ManagedServiceInstanceUpdate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRelationships

`func (o *ManagedServiceInstanceUpdate) GetRelationships() ManagedServiceInstanceUpdateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManagedServiceInstanceUpdate) GetRelationshipsOk() (*ManagedServiceInstanceUpdateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManagedServiceInstanceUpdate) SetRelationships(v ManagedServiceInstanceUpdateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManagedServiceInstanceUpdate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *ManagedServiceInstanceUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagedServiceInstanceUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagedServiceInstanceUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagedServiceInstanceUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


