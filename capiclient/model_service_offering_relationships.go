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

// checks if the ServiceOfferingRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOfferingRelationships{}

// ServiceOfferingRelationships struct for ServiceOfferingRelationships
type ServiceOfferingRelationships struct {
	ServiceBroker *V3AppsPostRequestRelationshipsSpace `json:"service_broker,omitempty"`
}

// NewServiceOfferingRelationships instantiates a new ServiceOfferingRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOfferingRelationships() *ServiceOfferingRelationships {
	this := ServiceOfferingRelationships{}
	return &this
}

// NewServiceOfferingRelationshipsWithDefaults instantiates a new ServiceOfferingRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOfferingRelationshipsWithDefaults() *ServiceOfferingRelationships {
	this := ServiceOfferingRelationships{}
	return &this
}

// GetServiceBroker returns the ServiceBroker field value if set, zero value otherwise.
func (o *ServiceOfferingRelationships) GetServiceBroker() V3AppsPostRequestRelationshipsSpace {
	if o == nil || IsNil(o.ServiceBroker) {
		var ret V3AppsPostRequestRelationshipsSpace
		return ret
	}
	return *o.ServiceBroker
}

// GetServiceBrokerOk returns a tuple with the ServiceBroker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingRelationships) GetServiceBrokerOk() (*V3AppsPostRequestRelationshipsSpace, bool) {
	if o == nil || IsNil(o.ServiceBroker) {
		return nil, false
	}
	return o.ServiceBroker, true
}

// HasServiceBroker returns a boolean if a field has been set.
func (o *ServiceOfferingRelationships) HasServiceBroker() bool {
	if o != nil && !IsNil(o.ServiceBroker) {
		return true
	}

	return false
}

// SetServiceBroker gets a reference to the given V3AppsPostRequestRelationshipsSpace and assigns it to the ServiceBroker field.
func (o *ServiceOfferingRelationships) SetServiceBroker(v V3AppsPostRequestRelationshipsSpace) {
	o.ServiceBroker = &v
}

func (o ServiceOfferingRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOfferingRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceBroker) {
		toSerialize["service_broker"] = o.ServiceBroker
	}
	return toSerialize, nil
}

type NullableServiceOfferingRelationships struct {
	value *ServiceOfferingRelationships
	isSet bool
}

func (v NullableServiceOfferingRelationships) Get() *ServiceOfferingRelationships {
	return v.value
}

func (v *NullableServiceOfferingRelationships) Set(val *ServiceOfferingRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOfferingRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOfferingRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOfferingRelationships(val *ServiceOfferingRelationships) *NullableServiceOfferingRelationships {
	return &NullableServiceOfferingRelationships{value: val, isSet: true}
}

func (v NullableServiceOfferingRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOfferingRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


