# ServiceInstanceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlan** | Pointer to [**ToOneRelationship**](ToOneRelationship.md) |  | [optional] 
**Space** | Pointer to [**ToOneRelationship**](ToOneRelationship.md) |  | [optional] 

## Methods

### NewServiceInstanceRelationships

`func NewServiceInstanceRelationships() *ServiceInstanceRelationships`

NewServiceInstanceRelationships instantiates a new ServiceInstanceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceRelationshipsWithDefaults

`func NewServiceInstanceRelationshipsWithDefaults() *ServiceInstanceRelationships`

NewServiceInstanceRelationshipsWithDefaults instantiates a new ServiceInstanceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePlan

`func (o *ServiceInstanceRelationships) GetServicePlan() ToOneRelationship`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *ServiceInstanceRelationships) GetServicePlanOk() (*ToOneRelationship, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *ServiceInstanceRelationships) SetServicePlan(v ToOneRelationship)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *ServiceInstanceRelationships) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### GetSpace

`func (o *ServiceInstanceRelationships) GetSpace() ToOneRelationship`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ServiceInstanceRelationships) GetSpaceOk() (*ToOneRelationship, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ServiceInstanceRelationships) SetSpace(v ToOneRelationship)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *ServiceInstanceRelationships) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


