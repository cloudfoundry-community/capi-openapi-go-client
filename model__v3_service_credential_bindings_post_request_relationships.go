/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V3ServiceCredentialBindingsPostRequestRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3ServiceCredentialBindingsPostRequestRelationships{}

// V3ServiceCredentialBindingsPostRequestRelationships struct for V3ServiceCredentialBindingsPostRequestRelationships
type V3ServiceCredentialBindingsPostRequestRelationships struct {
	App *V3AppsPostRequestRelationshipsSpace `json:"app,omitempty"`
	ServiceInstance V3AppsPostRequestRelationshipsSpace `json:"service_instance"`
}

type _V3ServiceCredentialBindingsPostRequestRelationships V3ServiceCredentialBindingsPostRequestRelationships

// NewV3ServiceCredentialBindingsPostRequestRelationships instantiates a new V3ServiceCredentialBindingsPostRequestRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3ServiceCredentialBindingsPostRequestRelationships(serviceInstance V3AppsPostRequestRelationshipsSpace) *V3ServiceCredentialBindingsPostRequestRelationships {
	this := V3ServiceCredentialBindingsPostRequestRelationships{}
	this.ServiceInstance = serviceInstance
	return &this
}

// NewV3ServiceCredentialBindingsPostRequestRelationshipsWithDefaults instantiates a new V3ServiceCredentialBindingsPostRequestRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ServiceCredentialBindingsPostRequestRelationshipsWithDefaults() *V3ServiceCredentialBindingsPostRequestRelationships {
	this := V3ServiceCredentialBindingsPostRequestRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *V3ServiceCredentialBindingsPostRequestRelationships) GetApp() V3AppsPostRequestRelationshipsSpace {
	if o == nil || IsNil(o.App) {
		var ret V3AppsPostRequestRelationshipsSpace
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3ServiceCredentialBindingsPostRequestRelationships) GetAppOk() (*V3AppsPostRequestRelationshipsSpace, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *V3ServiceCredentialBindingsPostRequestRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given V3AppsPostRequestRelationshipsSpace and assigns it to the App field.
func (o *V3ServiceCredentialBindingsPostRequestRelationships) SetApp(v V3AppsPostRequestRelationshipsSpace) {
	o.App = &v
}

// GetServiceInstance returns the ServiceInstance field value
func (o *V3ServiceCredentialBindingsPostRequestRelationships) GetServiceInstance() V3AppsPostRequestRelationshipsSpace {
	if o == nil {
		var ret V3AppsPostRequestRelationshipsSpace
		return ret
	}

	return o.ServiceInstance
}

// GetServiceInstanceOk returns a tuple with the ServiceInstance field value
// and a boolean to check if the value has been set.
func (o *V3ServiceCredentialBindingsPostRequestRelationships) GetServiceInstanceOk() (*V3AppsPostRequestRelationshipsSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceInstance, true
}

// SetServiceInstance sets field value
func (o *V3ServiceCredentialBindingsPostRequestRelationships) SetServiceInstance(v V3AppsPostRequestRelationshipsSpace) {
	o.ServiceInstance = v
}

func (o V3ServiceCredentialBindingsPostRequestRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3ServiceCredentialBindingsPostRequestRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	toSerialize["service_instance"] = o.ServiceInstance
	return toSerialize, nil
}

func (o *V3ServiceCredentialBindingsPostRequestRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_instance",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV3ServiceCredentialBindingsPostRequestRelationships := _V3ServiceCredentialBindingsPostRequestRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV3ServiceCredentialBindingsPostRequestRelationships)

	if err != nil {
		return err
	}

	*o = V3ServiceCredentialBindingsPostRequestRelationships(varV3ServiceCredentialBindingsPostRequestRelationships)

	return err
}

type NullableV3ServiceCredentialBindingsPostRequestRelationships struct {
	value *V3ServiceCredentialBindingsPostRequestRelationships
	isSet bool
}

func (v NullableV3ServiceCredentialBindingsPostRequestRelationships) Get() *V3ServiceCredentialBindingsPostRequestRelationships {
	return v.value
}

func (v *NullableV3ServiceCredentialBindingsPostRequestRelationships) Set(val *V3ServiceCredentialBindingsPostRequestRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableV3ServiceCredentialBindingsPostRequestRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableV3ServiceCredentialBindingsPostRequestRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3ServiceCredentialBindingsPostRequestRelationships(val *V3ServiceCredentialBindingsPostRequestRelationships) *NullableV3ServiceCredentialBindingsPostRequestRelationships {
	return &NullableV3ServiceCredentialBindingsPostRequestRelationships{value: val, isSet: true}
}

func (v NullableV3ServiceCredentialBindingsPostRequestRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3ServiceCredentialBindingsPostRequestRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

