# ServiceUsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ServiceUsageEventLinks**](ServiceUsageEventLinks.md) |  | [optional] 
**Organization** | Pointer to [**NullableRoleRelationshipsSpaceData**](RoleRelationshipsSpaceData.md) |  | [optional] 
**ServiceBroker** | Pointer to [**NullableServiceUsageEventServiceBroker**](ServiceUsageEventServiceBroker.md) |  | [optional] 
**ServiceInstance** | Pointer to [**NullableServiceUsageEventServiceInstance**](ServiceUsageEventServiceInstance.md) |  | [optional] 
**ServiceOffering** | Pointer to [**NullableServiceUsageEventServiceBroker**](ServiceUsageEventServiceBroker.md) |  | [optional] 
**ServicePlan** | Pointer to [**NullableServiceUsageEventServiceBroker**](ServiceUsageEventServiceBroker.md) |  | [optional] 
**Space** | Pointer to [**NullableServiceUsageEventServiceBroker**](ServiceUsageEventServiceBroker.md) |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceUsageEvent

`func NewServiceUsageEvent() *ServiceUsageEvent`

NewServiceUsageEvent instantiates a new ServiceUsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUsageEventWithDefaults

`func NewServiceUsageEventWithDefaults() *ServiceUsageEvent`

NewServiceUsageEventWithDefaults instantiates a new ServiceUsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceUsageEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceUsageEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceUsageEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceUsageEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceUsageEvent) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceUsageEvent) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceUsageEvent) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceUsageEvent) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceUsageEvent) GetLinks() ServiceUsageEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceUsageEvent) GetLinksOk() (*ServiceUsageEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceUsageEvent) SetLinks(v ServiceUsageEventLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceUsageEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrganization

`func (o *ServiceUsageEvent) GetOrganization() RoleRelationshipsSpaceData`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServiceUsageEvent) GetOrganizationOk() (*RoleRelationshipsSpaceData, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServiceUsageEvent) SetOrganization(v RoleRelationshipsSpaceData)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServiceUsageEvent) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ServiceUsageEvent) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ServiceUsageEvent) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetServiceBroker

`func (o *ServiceUsageEvent) GetServiceBroker() ServiceUsageEventServiceBroker`

GetServiceBroker returns the ServiceBroker field if non-nil, zero value otherwise.

### GetServiceBrokerOk

`func (o *ServiceUsageEvent) GetServiceBrokerOk() (*ServiceUsageEventServiceBroker, bool)`

GetServiceBrokerOk returns a tuple with the ServiceBroker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBroker

`func (o *ServiceUsageEvent) SetServiceBroker(v ServiceUsageEventServiceBroker)`

SetServiceBroker sets ServiceBroker field to given value.

### HasServiceBroker

`func (o *ServiceUsageEvent) HasServiceBroker() bool`

HasServiceBroker returns a boolean if a field has been set.

### SetServiceBrokerNil

`func (o *ServiceUsageEvent) SetServiceBrokerNil(b bool)`

 SetServiceBrokerNil sets the value for ServiceBroker to be an explicit nil

### UnsetServiceBroker
`func (o *ServiceUsageEvent) UnsetServiceBroker()`

UnsetServiceBroker ensures that no value is present for ServiceBroker, not even an explicit nil
### GetServiceInstance

`func (o *ServiceUsageEvent) GetServiceInstance() ServiceUsageEventServiceInstance`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *ServiceUsageEvent) GetServiceInstanceOk() (*ServiceUsageEventServiceInstance, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *ServiceUsageEvent) SetServiceInstance(v ServiceUsageEventServiceInstance)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *ServiceUsageEvent) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.

### SetServiceInstanceNil

`func (o *ServiceUsageEvent) SetServiceInstanceNil(b bool)`

 SetServiceInstanceNil sets the value for ServiceInstance to be an explicit nil

### UnsetServiceInstance
`func (o *ServiceUsageEvent) UnsetServiceInstance()`

UnsetServiceInstance ensures that no value is present for ServiceInstance, not even an explicit nil
### GetServiceOffering

`func (o *ServiceUsageEvent) GetServiceOffering() ServiceUsageEventServiceBroker`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ServiceUsageEvent) GetServiceOfferingOk() (*ServiceUsageEventServiceBroker, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ServiceUsageEvent) SetServiceOffering(v ServiceUsageEventServiceBroker)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ServiceUsageEvent) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### SetServiceOfferingNil

`func (o *ServiceUsageEvent) SetServiceOfferingNil(b bool)`

 SetServiceOfferingNil sets the value for ServiceOffering to be an explicit nil

### UnsetServiceOffering
`func (o *ServiceUsageEvent) UnsetServiceOffering()`

UnsetServiceOffering ensures that no value is present for ServiceOffering, not even an explicit nil
### GetServicePlan

`func (o *ServiceUsageEvent) GetServicePlan() ServiceUsageEventServiceBroker`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *ServiceUsageEvent) GetServicePlanOk() (*ServiceUsageEventServiceBroker, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *ServiceUsageEvent) SetServicePlan(v ServiceUsageEventServiceBroker)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *ServiceUsageEvent) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### SetServicePlanNil

`func (o *ServiceUsageEvent) SetServicePlanNil(b bool)`

 SetServicePlanNil sets the value for ServicePlan to be an explicit nil

### UnsetServicePlan
`func (o *ServiceUsageEvent) UnsetServicePlan()`

UnsetServicePlan ensures that no value is present for ServicePlan, not even an explicit nil
### GetSpace

`func (o *ServiceUsageEvent) GetSpace() ServiceUsageEventServiceBroker`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ServiceUsageEvent) GetSpaceOk() (*ServiceUsageEventServiceBroker, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ServiceUsageEvent) SetSpace(v ServiceUsageEventServiceBroker)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *ServiceUsageEvent) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### SetSpaceNil

`func (o *ServiceUsageEvent) SetSpaceNil(b bool)`

 SetSpaceNil sets the value for Space to be an explicit nil

### UnsetSpace
`func (o *ServiceUsageEvent) UnsetSpace()`

UnsetSpace ensures that no value is present for Space, not even an explicit nil
### GetState

`func (o *ServiceUsageEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceUsageEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceUsageEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ServiceUsageEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ServiceUsageEvent) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ServiceUsageEvent) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUpdatedAt

`func (o *ServiceUsageEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceUsageEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceUsageEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceUsageEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


