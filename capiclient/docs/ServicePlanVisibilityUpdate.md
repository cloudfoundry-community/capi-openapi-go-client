# ServicePlanVisibilityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ServicePlanVisibilityUpdateMetadata**](ServicePlanVisibilityUpdateMetadata.md) |  | [optional] 
**Organizations** | Pointer to [**[]ApplyOrganizationQuotaToOrganizationsRequestDataInner**](ApplyOrganizationQuotaToOrganizationsRequestDataInner.md) | List of organizations where the plan should be visible (required for organization type) | [optional] 
**Type** | **string** | Denotes the visibility of the plan | 

## Methods

### NewServicePlanVisibilityUpdate

`func NewServicePlanVisibilityUpdate(type_ string, ) *ServicePlanVisibilityUpdate`

NewServicePlanVisibilityUpdate instantiates a new ServicePlanVisibilityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanVisibilityUpdateWithDefaults

`func NewServicePlanVisibilityUpdateWithDefaults() *ServicePlanVisibilityUpdate`

NewServicePlanVisibilityUpdateWithDefaults instantiates a new ServicePlanVisibilityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServicePlanVisibilityUpdate) GetMetadata() ServicePlanVisibilityUpdateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServicePlanVisibilityUpdate) GetMetadataOk() (*ServicePlanVisibilityUpdateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServicePlanVisibilityUpdate) SetMetadata(v ServicePlanVisibilityUpdateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServicePlanVisibilityUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrganizations

`func (o *ServicePlanVisibilityUpdate) GetOrganizations() []ApplyOrganizationQuotaToOrganizationsRequestDataInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ServicePlanVisibilityUpdate) GetOrganizationsOk() (*[]ApplyOrganizationQuotaToOrganizationsRequestDataInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ServicePlanVisibilityUpdate) SetOrganizations(v []ApplyOrganizationQuotaToOrganizationsRequestDataInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ServicePlanVisibilityUpdate) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetType

`func (o *ServicePlanVisibilityUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePlanVisibilityUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePlanVisibilityUpdate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


