# ServicePlanVisibilityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]V3AppsPostRequestRelationshipsSpaceData**](V3AppsPostRequestRelationshipsSpaceData.md) |  | [optional] 
**Type** | Pointer to **string** | Denotes the visibility of the plan; can be public, admin, organization | [optional] 

## Methods

### NewServicePlanVisibilityUpdate

`func NewServicePlanVisibilityUpdate() *ServicePlanVisibilityUpdate`

NewServicePlanVisibilityUpdate instantiates a new ServicePlanVisibilityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanVisibilityUpdateWithDefaults

`func NewServicePlanVisibilityUpdateWithDefaults() *ServicePlanVisibilityUpdate`

NewServicePlanVisibilityUpdateWithDefaults instantiates a new ServicePlanVisibilityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *ServicePlanVisibilityUpdate) GetOrganizations() []V3AppsPostRequestRelationshipsSpaceData`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ServicePlanVisibilityUpdate) GetOrganizationsOk() (*[]V3AppsPostRequestRelationshipsSpaceData, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ServicePlanVisibilityUpdate) SetOrganizations(v []V3AppsPostRequestRelationshipsSpaceData)`

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

### HasType

`func (o *ServicePlanVisibilityUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


