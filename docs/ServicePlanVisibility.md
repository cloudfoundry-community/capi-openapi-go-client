# ServicePlanVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ServicePlanVisibilityMetadata**](ServicePlanVisibilityMetadata.md) |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationVisibility**](OrganizationVisibility.md) | List of organizations where the plan is visible (only for organization type) | [optional] 
**Space** | Pointer to [**ServicePlanVisibilitySpace**](ServicePlanVisibilitySpace.md) |  | [optional] 
**Type** | **string** | Denotes the visibility of the plan | 

## Methods

### NewServicePlanVisibility

`func NewServicePlanVisibility(type_ string, ) *ServicePlanVisibility`

NewServicePlanVisibility instantiates a new ServicePlanVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanVisibilityWithDefaults

`func NewServicePlanVisibilityWithDefaults() *ServicePlanVisibility`

NewServicePlanVisibilityWithDefaults instantiates a new ServicePlanVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServicePlanVisibility) GetMetadata() ServicePlanVisibilityMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServicePlanVisibility) GetMetadataOk() (*ServicePlanVisibilityMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServicePlanVisibility) SetMetadata(v ServicePlanVisibilityMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServicePlanVisibility) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrganizations

`func (o *ServicePlanVisibility) GetOrganizations() []OrganizationVisibility`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ServicePlanVisibility) GetOrganizationsOk() (*[]OrganizationVisibility, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ServicePlanVisibility) SetOrganizations(v []OrganizationVisibility)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ServicePlanVisibility) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetSpace

`func (o *ServicePlanVisibility) GetSpace() ServicePlanVisibilitySpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ServicePlanVisibility) GetSpaceOk() (*ServicePlanVisibilitySpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ServicePlanVisibility) SetSpace(v ServicePlanVisibilitySpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *ServicePlanVisibility) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetType

`func (o *ServicePlanVisibility) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePlanVisibility) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePlanVisibility) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


