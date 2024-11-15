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

// checks if the AppRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationships{}

// AppRelationships struct for AppRelationships
type AppRelationships struct {
	CurrentDroplet *V3DropletsPostRequestRelationshipsApp `json:"current_droplet,omitempty"`
	Space V3DropletsPostRequestRelationshipsApp `json:"space"`
}

type _AppRelationships AppRelationships

// NewAppRelationships instantiates a new AppRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationships(space V3DropletsPostRequestRelationshipsApp) *AppRelationships {
	this := AppRelationships{}
	this.Space = space
	return &this
}

// NewAppRelationshipsWithDefaults instantiates a new AppRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsWithDefaults() *AppRelationships {
	this := AppRelationships{}
	return &this
}

// GetCurrentDroplet returns the CurrentDroplet field value if set, zero value otherwise.
func (o *AppRelationships) GetCurrentDroplet() V3DropletsPostRequestRelationshipsApp {
	if o == nil || IsNil(o.CurrentDroplet) {
		var ret V3DropletsPostRequestRelationshipsApp
		return ret
	}
	return *o.CurrentDroplet
}

// GetCurrentDropletOk returns a tuple with the CurrentDroplet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetCurrentDropletOk() (*V3DropletsPostRequestRelationshipsApp, bool) {
	if o == nil || IsNil(o.CurrentDroplet) {
		return nil, false
	}
	return o.CurrentDroplet, true
}

// HasCurrentDroplet returns a boolean if a field has been set.
func (o *AppRelationships) HasCurrentDroplet() bool {
	if o != nil && !IsNil(o.CurrentDroplet) {
		return true
	}

	return false
}

// SetCurrentDroplet gets a reference to the given V3DropletsPostRequestRelationshipsApp and assigns it to the CurrentDroplet field.
func (o *AppRelationships) SetCurrentDroplet(v V3DropletsPostRequestRelationshipsApp) {
	o.CurrentDroplet = &v
}

// GetSpace returns the Space field value
func (o *AppRelationships) GetSpace() V3DropletsPostRequestRelationshipsApp {
	if o == nil {
		var ret V3DropletsPostRequestRelationshipsApp
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetSpaceOk() (*V3DropletsPostRequestRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *AppRelationships) SetSpace(v V3DropletsPostRequestRelationshipsApp) {
	o.Space = v
}

func (o AppRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentDroplet) {
		toSerialize["current_droplet"] = o.CurrentDroplet
	}
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

func (o *AppRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"space",
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

	varAppRelationships := _AppRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppRelationships)

	if err != nil {
		return err
	}

	*o = AppRelationships(varAppRelationships)

	return err
}

type NullableAppRelationships struct {
	value *AppRelationships
	isSet bool
}

func (v NullableAppRelationships) Get() *AppRelationships {
	return v.value
}

func (v *NullableAppRelationships) Set(val *AppRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationships(val *AppRelationships) *NullableAppRelationships {
	return &NullableAppRelationships{value: val, isSet: true}
}

func (v NullableAppRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

