# ServiceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**LastOperation** | Pointer to [**LastOperation**](LastOperation.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**MaintenanceInfo** | Pointer to [**Get200ResponseLinksCloudControllerV2Meta**](Get200ResponseLinksCloudControllerV2Meta.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceInstanceRelationships**](ServiceInstanceRelationships.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpgradeAvailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceInstance

`func NewServiceInstance() *ServiceInstance`

NewServiceInstance instantiates a new ServiceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceWithDefaults

`func NewServiceInstanceWithDefaults() *ServiceInstance`

NewServiceInstanceWithDefaults instantiates a new ServiceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *ServiceInstance) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *ServiceInstance) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *ServiceInstance) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *ServiceInstance) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceInstance) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceInstance) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceInstance) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceInstance) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastOperation

`func (o *ServiceInstance) GetLastOperation() LastOperation`

GetLastOperation returns the LastOperation field if non-nil, zero value otherwise.

### GetLastOperationOk

`func (o *ServiceInstance) GetLastOperationOk() (*LastOperation, bool)`

GetLastOperationOk returns a tuple with the LastOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperation

`func (o *ServiceInstance) SetLastOperation(v LastOperation)`

SetLastOperation sets LastOperation field to given value.

### HasLastOperation

`func (o *ServiceInstance) HasLastOperation() bool`

HasLastOperation returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceInstance) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceInstance) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceInstance) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *ServiceInstance) GetMaintenanceInfo() Get200ResponseLinksCloudControllerV2Meta`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ServiceInstance) GetMaintenanceInfoOk() (*Get200ResponseLinksCloudControllerV2Meta, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ServiceInstance) SetMaintenanceInfo(v Get200ResponseLinksCloudControllerV2Meta)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ServiceInstance) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceInstance) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceInstance) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceInstance) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServiceInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceInstance) GetRelationships() ServiceInstanceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceInstance) GetRelationshipsOk() (*ServiceInstanceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceInstance) SetRelationships(v ServiceInstanceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceInstance) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *ServiceInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ServiceInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceInstance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpgradeAvailable

`func (o *ServiceInstance) GetUpgradeAvailable() bool`

GetUpgradeAvailable returns the UpgradeAvailable field if non-nil, zero value otherwise.

### GetUpgradeAvailableOk

`func (o *ServiceInstance) GetUpgradeAvailableOk() (*bool, bool)`

GetUpgradeAvailableOk returns a tuple with the UpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeAvailable

`func (o *ServiceInstance) SetUpgradeAvailable(v bool)`

SetUpgradeAvailable sets UpgradeAvailable field to given value.

### HasUpgradeAvailable

`func (o *ServiceInstance) HasUpgradeAvailable() bool`

HasUpgradeAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


