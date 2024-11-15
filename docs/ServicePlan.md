# ServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**BrokerCatalog** | Pointer to [**BrokerCatalog**](BrokerCatalog.md) |  | [optional] 
**Costs** | Pointer to [**[]Cost**](Cost.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ServicePlanLinks**](ServicePlanLinks.md) |  | [optional] 
**MaintenanceInfo** | Pointer to [**MaintenanceInfo**](MaintenanceInfo.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServicePlanRelationships**](ServicePlanRelationships.md) |  | [optional] 
**Schemas** | Pointer to [**PlanSchemas**](PlanSchemas.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VisibilityType** | Pointer to **string** | Denotes the visibility of the plan; can be public, admin, organization, space | [optional] 

## Methods

### NewServicePlan

`func NewServicePlan() *ServicePlan`

NewServicePlan instantiates a new ServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanWithDefaults

`func NewServicePlanWithDefaults() *ServicePlan`

NewServicePlanWithDefaults instantiates a new ServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ServicePlan) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServicePlan) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServicePlan) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ServicePlan) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBrokerCatalog

`func (o *ServicePlan) GetBrokerCatalog() BrokerCatalog`

GetBrokerCatalog returns the BrokerCatalog field if non-nil, zero value otherwise.

### GetBrokerCatalogOk

`func (o *ServicePlan) GetBrokerCatalogOk() (*BrokerCatalog, bool)`

GetBrokerCatalogOk returns a tuple with the BrokerCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerCatalog

`func (o *ServicePlan) SetBrokerCatalog(v BrokerCatalog)`

SetBrokerCatalog sets BrokerCatalog field to given value.

### HasBrokerCatalog

`func (o *ServicePlan) HasBrokerCatalog() bool`

HasBrokerCatalog returns a boolean if a field has been set.

### GetCosts

`func (o *ServicePlan) GetCosts() []Cost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *ServicePlan) GetCostsOk() (*[]Cost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *ServicePlan) SetCosts(v []Cost)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *ServicePlan) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServicePlan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicePlan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicePlan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServicePlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFree

`func (o *ServicePlan) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *ServicePlan) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *ServicePlan) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *ServicePlan) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetGuid

`func (o *ServicePlan) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServicePlan) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServicePlan) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServicePlan) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *ServicePlan) GetLinks() ServicePlanLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServicePlan) GetLinksOk() (*ServicePlanLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServicePlan) SetLinks(v ServicePlanLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServicePlan) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *ServicePlan) GetMaintenanceInfo() MaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ServicePlan) GetMaintenanceInfoOk() (*MaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ServicePlan) SetMaintenanceInfo(v MaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ServicePlan) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *ServicePlan) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServicePlan) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServicePlan) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServicePlan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServicePlan) GetRelationships() ServicePlanRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServicePlan) GetRelationshipsOk() (*ServicePlanRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServicePlan) SetRelationships(v ServicePlanRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServicePlan) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSchemas

`func (o *ServicePlan) GetSchemas() PlanSchemas`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServicePlan) GetSchemasOk() (*PlanSchemas, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServicePlan) SetSchemas(v PlanSchemas)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServicePlan) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServicePlan) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServicePlan) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServicePlan) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServicePlan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibilityType

`func (o *ServicePlan) GetVisibilityType() string`

GetVisibilityType returns the VisibilityType field if non-nil, zero value otherwise.

### GetVisibilityTypeOk

`func (o *ServicePlan) GetVisibilityTypeOk() (*string, bool)`

GetVisibilityTypeOk returns a tuple with the VisibilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityType

`func (o *ServicePlan) SetVisibilityType(v string)`

SetVisibilityType sets VisibilityType field to given value.

### HasVisibilityType

`func (o *ServicePlan) HasVisibilityType() bool`

HasVisibilityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


