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

// checks if the ServiceUsageEventServiceInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUsageEventServiceInstance{}

// ServiceUsageEventServiceInstance struct for ServiceUsageEventServiceInstance
type ServiceUsageEventServiceInstance struct {
	Guid *string `json:"guid,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewServiceUsageEventServiceInstance instantiates a new ServiceUsageEventServiceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUsageEventServiceInstance() *ServiceUsageEventServiceInstance {
	this := ServiceUsageEventServiceInstance{}
	return &this
}

// NewServiceUsageEventServiceInstanceWithDefaults instantiates a new ServiceUsageEventServiceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUsageEventServiceInstanceWithDefaults() *ServiceUsageEventServiceInstance {
	this := ServiceUsageEventServiceInstance{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ServiceUsageEventServiceInstance) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsageEventServiceInstance) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ServiceUsageEventServiceInstance) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ServiceUsageEventServiceInstance) SetGuid(v string) {
	o.Guid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceUsageEventServiceInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsageEventServiceInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceUsageEventServiceInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceUsageEventServiceInstance) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceUsageEventServiceInstance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsageEventServiceInstance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceUsageEventServiceInstance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceUsageEventServiceInstance) SetType(v string) {
	o.Type = &v
}

func (o ServiceUsageEventServiceInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceUsageEventServiceInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableServiceUsageEventServiceInstance struct {
	value *ServiceUsageEventServiceInstance
	isSet bool
}

func (v NullableServiceUsageEventServiceInstance) Get() *ServiceUsageEventServiceInstance {
	return v.value
}

func (v *NullableServiceUsageEventServiceInstance) Set(val *ServiceUsageEventServiceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUsageEventServiceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUsageEventServiceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUsageEventServiceInstance(val *ServiceUsageEventServiceInstance) *NullableServiceUsageEventServiceInstance {
	return &NullableServiceUsageEventServiceInstance{value: val, isSet: true}
}

func (v NullableServiceUsageEventServiceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUsageEventServiceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

