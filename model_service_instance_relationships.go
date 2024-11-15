/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
)

// checks if the ServiceInstanceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceInstanceRelationships{}

// ServiceInstanceRelationships struct for ServiceInstanceRelationships
type ServiceInstanceRelationships struct {
	ServicePlan *ToOneRelationship `json:"service_plan,omitempty"`
	Space *ToOneRelationship `json:"space,omitempty"`
}

// NewServiceInstanceRelationships instantiates a new ServiceInstanceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInstanceRelationships() *ServiceInstanceRelationships {
	this := ServiceInstanceRelationships{}
	return &this
}

// NewServiceInstanceRelationshipsWithDefaults instantiates a new ServiceInstanceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInstanceRelationshipsWithDefaults() *ServiceInstanceRelationships {
	this := ServiceInstanceRelationships{}
	return &this
}

// GetServicePlan returns the ServicePlan field value if set, zero value otherwise.
func (o *ServiceInstanceRelationships) GetServicePlan() ToOneRelationship {
	if o == nil || IsNil(o.ServicePlan) {
		var ret ToOneRelationship
		return ret
	}
	return *o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceRelationships) GetServicePlanOk() (*ToOneRelationship, bool) {
	if o == nil || IsNil(o.ServicePlan) {
		return nil, false
	}
	return o.ServicePlan, true
}

// HasServicePlan returns a boolean if a field has been set.
func (o *ServiceInstanceRelationships) HasServicePlan() bool {
	if o != nil && !IsNil(o.ServicePlan) {
		return true
	}

	return false
}

// SetServicePlan gets a reference to the given ToOneRelationship and assigns it to the ServicePlan field.
func (o *ServiceInstanceRelationships) SetServicePlan(v ToOneRelationship) {
	o.ServicePlan = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *ServiceInstanceRelationships) GetSpace() ToOneRelationship {
	if o == nil || IsNil(o.Space) {
		var ret ToOneRelationship
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceRelationships) GetSpaceOk() (*ToOneRelationship, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *ServiceInstanceRelationships) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given ToOneRelationship and assigns it to the Space field.
func (o *ServiceInstanceRelationships) SetSpace(v ToOneRelationship) {
	o.Space = &v
}

func (o ServiceInstanceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceInstanceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServicePlan) {
		toSerialize["service_plan"] = o.ServicePlan
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	return toSerialize, nil
}

type NullableServiceInstanceRelationships struct {
	value *ServiceInstanceRelationships
	isSet bool
}

func (v NullableServiceInstanceRelationships) Get() *ServiceInstanceRelationships {
	return v.value
}

func (v *NullableServiceInstanceRelationships) Set(val *ServiceInstanceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInstanceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInstanceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInstanceRelationships(val *ServiceInstanceRelationships) *NullableServiceInstanceRelationships {
	return &NullableServiceInstanceRelationships{value: val, isSet: true}
}

func (v NullableServiceInstanceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInstanceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


