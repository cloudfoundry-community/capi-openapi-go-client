# ServicePlanCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**BrokerCatalog** | Pointer to [**BrokerCatalog**](BrokerCatalog.md) |  | [optional] 
**Costs** | Pointer to [**[]Cost**](Cost.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServicePlanRelationships**](ServicePlanRelationships.md) |  | [optional] 
**Schemas** | Pointer to [**PlanSchemas**](PlanSchemas.md) |  | [optional] 
**VisibilityType** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePlanCreate

`func NewServicePlanCreate() *ServicePlanCreate`

NewServicePlanCreate instantiates a new ServicePlanCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanCreateWithDefaults

`func NewServicePlanCreateWithDefaults() *ServicePlanCreate`

NewServicePlanCreateWithDefaults instantiates a new ServicePlanCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ServicePlanCreate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServicePlanCreate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServicePlanCreate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ServicePlanCreate) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBrokerCatalog

`func (o *ServicePlanCreate) GetBrokerCatalog() BrokerCatalog`

GetBrokerCatalog returns the BrokerCatalog field if non-nil, zero value otherwise.

### GetBrokerCatalogOk

`func (o *ServicePlanCreate) GetBrokerCatalogOk() (*BrokerCatalog, bool)`

GetBrokerCatalogOk returns a tuple with the BrokerCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerCatalog

`func (o *ServicePlanCreate) SetBrokerCatalog(v BrokerCatalog)`

SetBrokerCatalog sets BrokerCatalog field to given value.

### HasBrokerCatalog

`func (o *ServicePlanCreate) HasBrokerCatalog() bool`

HasBrokerCatalog returns a boolean if a field has been set.

### GetCosts

`func (o *ServicePlanCreate) GetCosts() []Cost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *ServicePlanCreate) GetCostsOk() (*[]Cost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *ServicePlanCreate) SetCosts(v []Cost)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *ServicePlanCreate) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetDescription

`func (o *ServicePlanCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePlanCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePlanCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePlanCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFree

`func (o *ServicePlanCreate) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *ServicePlanCreate) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *ServicePlanCreate) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *ServicePlanCreate) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetName

`func (o *ServicePlanCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePlanCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePlanCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePlanCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServicePlanCreate) GetRelationships() ServicePlanRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServicePlanCreate) GetRelationshipsOk() (*ServicePlanRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServicePlanCreate) SetRelationships(v ServicePlanRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServicePlanCreate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSchemas

`func (o *ServicePlanCreate) GetSchemas() PlanSchemas`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServicePlanCreate) GetSchemasOk() (*PlanSchemas, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServicePlanCreate) SetSchemas(v PlanSchemas)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServicePlanCreate) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetVisibilityType

`func (o *ServicePlanCreate) GetVisibilityType() string`

GetVisibilityType returns the VisibilityType field if non-nil, zero value otherwise.

### GetVisibilityTypeOk

`func (o *ServicePlanCreate) GetVisibilityTypeOk() (*string, bool)`

GetVisibilityTypeOk returns a tuple with the VisibilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityType

`func (o *ServicePlanCreate) SetVisibilityType(v string)`

SetVisibilityType sets VisibilityType field to given value.

### HasVisibilityType

`func (o *ServicePlanCreate) HasVisibilityType() bool`

HasVisibilityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


