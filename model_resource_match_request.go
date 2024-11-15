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

// checks if the ResourceMatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceMatchRequest{}

// ResourceMatchRequest struct for ResourceMatchRequest
type ResourceMatchRequest struct {
	Resources []ResourceMatchObject `json:"resources,omitempty"`
}

// NewResourceMatchRequest instantiates a new ResourceMatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMatchRequest() *ResourceMatchRequest {
	this := ResourceMatchRequest{}
	return &this
}

// NewResourceMatchRequestWithDefaults instantiates a new ResourceMatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMatchRequestWithDefaults() *ResourceMatchRequest {
	this := ResourceMatchRequest{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceMatchRequest) GetResources() []ResourceMatchObject {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourceMatchObject
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMatchRequest) GetResourcesOk() ([]ResourceMatchObject, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceMatchRequest) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceMatchObject and assigns it to the Resources field.
func (o *ResourceMatchRequest) SetResources(v []ResourceMatchObject) {
	o.Resources = v
}

func (o ResourceMatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceMatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableResourceMatchRequest struct {
	value *ResourceMatchRequest
	isSet bool
}

func (v NullableResourceMatchRequest) Get() *ResourceMatchRequest {
	return v.value
}

func (v *NullableResourceMatchRequest) Set(val *ResourceMatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMatchRequest(val *ResourceMatchRequest) *NullableResourceMatchRequest {
	return &NullableResourceMatchRequest{value: val, isSet: true}
}

func (v NullableResourceMatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


