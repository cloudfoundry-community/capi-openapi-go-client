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

// checks if the V3AppsGuidDropletsCurrentGet200ResponseLifecycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3AppsGuidDropletsCurrentGet200ResponseLifecycle{}

// V3AppsGuidDropletsCurrentGet200ResponseLifecycle struct for V3AppsGuidDropletsCurrentGet200ResponseLifecycle
type V3AppsGuidDropletsCurrentGet200ResponseLifecycle struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV3AppsGuidDropletsCurrentGet200ResponseLifecycle instantiates a new V3AppsGuidDropletsCurrentGet200ResponseLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AppsGuidDropletsCurrentGet200ResponseLifecycle() *V3AppsGuidDropletsCurrentGet200ResponseLifecycle {
	this := V3AppsGuidDropletsCurrentGet200ResponseLifecycle{}
	return &this
}

// NewV3AppsGuidDropletsCurrentGet200ResponseLifecycleWithDefaults instantiates a new V3AppsGuidDropletsCurrentGet200ResponseLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AppsGuidDropletsCurrentGet200ResponseLifecycleWithDefaults() *V3AppsGuidDropletsCurrentGet200ResponseLifecycle {
	this := V3AppsGuidDropletsCurrentGet200ResponseLifecycle{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) SetType(v string) {
	o.Type = &v
}

func (o V3AppsGuidDropletsCurrentGet200ResponseLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3AppsGuidDropletsCurrentGet200ResponseLifecycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle struct {
	value *V3AppsGuidDropletsCurrentGet200ResponseLifecycle
	isSet bool
}

func (v NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) Get() *V3AppsGuidDropletsCurrentGet200ResponseLifecycle {
	return v.value
}

func (v *NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) Set(val *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle(val *V3AppsGuidDropletsCurrentGet200ResponseLifecycle) *NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle {
	return &NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle{value: val, isSet: true}
}

func (v NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AppsGuidDropletsCurrentGet200ResponseLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


