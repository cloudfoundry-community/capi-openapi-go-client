# PlanSchemas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceBinding** | Pointer to [**PlanSchemasServiceBinding**](PlanSchemasServiceBinding.md) |  | [optional] 
**ServiceInstance** | Pointer to [**PlanSchemasServiceInstance**](PlanSchemasServiceInstance.md) |  | [optional] 

## Methods

### NewPlanSchemas

`func NewPlanSchemas() *PlanSchemas`

NewPlanSchemas instantiates a new PlanSchemas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSchemasWithDefaults

`func NewPlanSchemasWithDefaults() *PlanSchemas`

NewPlanSchemasWithDefaults instantiates a new PlanSchemas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceBinding

`func (o *PlanSchemas) GetServiceBinding() PlanSchemasServiceBinding`

GetServiceBinding returns the ServiceBinding field if non-nil, zero value otherwise.

### GetServiceBindingOk

`func (o *PlanSchemas) GetServiceBindingOk() (*PlanSchemasServiceBinding, bool)`

GetServiceBindingOk returns a tuple with the ServiceBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBinding

`func (o *PlanSchemas) SetServiceBinding(v PlanSchemasServiceBinding)`

SetServiceBinding sets ServiceBinding field to given value.

### HasServiceBinding

`func (o *PlanSchemas) HasServiceBinding() bool`

HasServiceBinding returns a boolean if a field has been set.

### GetServiceInstance

`func (o *PlanSchemas) GetServiceInstance() PlanSchemasServiceInstance`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *PlanSchemas) GetServiceInstanceOk() (*PlanSchemasServiceInstance, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *PlanSchemas) SetServiceInstance(v PlanSchemasServiceInstance)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *PlanSchemas) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


