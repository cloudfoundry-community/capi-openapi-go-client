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

// checks if the ServiceUsageEventServiceBroker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUsageEventServiceBroker{}

// ServiceUsageEventServiceBroker struct for ServiceUsageEventServiceBroker
type ServiceUsageEventServiceBroker struct {
	Guid *string `json:"guid,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewServiceUsageEventServiceBroker instantiates a new ServiceUsageEventServiceBroker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUsageEventServiceBroker() *ServiceUsageEventServiceBroker {
	this := ServiceUsageEventServiceBroker{}
	return &this
}

// NewServiceUsageEventServiceBrokerWithDefaults instantiates a new ServiceUsageEventServiceBroker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUsageEventServiceBrokerWithDefaults() *ServiceUsageEventServiceBroker {
	this := ServiceUsageEventServiceBroker{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ServiceUsageEventServiceBroker) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsageEventServiceBroker) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ServiceUsageEventServiceBroker) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ServiceUsageEventServiceBroker) SetGuid(v string) {
	o.Guid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceUsageEventServiceBroker) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsageEventServiceBroker) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceUsageEventServiceBroker) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceUsageEventServiceBroker) SetName(v string) {
	o.Name = &v
}

func (o ServiceUsageEventServiceBroker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceUsageEventServiceBroker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableServiceUsageEventServiceBroker struct {
	value *ServiceUsageEventServiceBroker
	isSet bool
}

func (v NullableServiceUsageEventServiceBroker) Get() *ServiceUsageEventServiceBroker {
	return v.value
}

func (v *NullableServiceUsageEventServiceBroker) Set(val *ServiceUsageEventServiceBroker) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUsageEventServiceBroker) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUsageEventServiceBroker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUsageEventServiceBroker(val *ServiceUsageEventServiceBroker) *NullableServiceUsageEventServiceBroker {
	return &NullableServiceUsageEventServiceBroker{value: val, isSet: true}
}

func (v NullableServiceUsageEventServiceBroker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUsageEventServiceBroker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


