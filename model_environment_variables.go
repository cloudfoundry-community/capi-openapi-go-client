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

// checks if the EnvironmentVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariables{}

// EnvironmentVariables struct for EnvironmentVariables
type EnvironmentVariables struct {
	Links *EnvironmentVariablesLinks `json:"links,omitempty"`
	Var *map[string]string `json:"var,omitempty"`
}

// NewEnvironmentVariables instantiates a new EnvironmentVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariables() *EnvironmentVariables {
	this := EnvironmentVariables{}
	return &this
}

// NewEnvironmentVariablesWithDefaults instantiates a new EnvironmentVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariablesWithDefaults() *EnvironmentVariables {
	this := EnvironmentVariables{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentVariables) GetLinks() EnvironmentVariablesLinks {
	if o == nil || IsNil(o.Links) {
		var ret EnvironmentVariablesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariables) GetLinksOk() (*EnvironmentVariablesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentVariables) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EnvironmentVariablesLinks and assigns it to the Links field.
func (o *EnvironmentVariables) SetLinks(v EnvironmentVariablesLinks) {
	o.Links = &v
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *EnvironmentVariables) GetVar() map[string]string {
	if o == nil || IsNil(o.Var) {
		var ret map[string]string
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariables) GetVarOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Var) {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *EnvironmentVariables) HasVar() bool {
	if o != nil && !IsNil(o.Var) {
		return true
	}

	return false
}

// SetVar gets a reference to the given map[string]string and assigns it to the Var field.
func (o *EnvironmentVariables) SetVar(v map[string]string) {
	o.Var = &v
}

func (o EnvironmentVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Var) {
		toSerialize["var"] = o.Var
	}
	return toSerialize, nil
}

type NullableEnvironmentVariables struct {
	value *EnvironmentVariables
	isSet bool
}

func (v NullableEnvironmentVariables) Get() *EnvironmentVariables {
	return v.value
}

func (v *NullableEnvironmentVariables) Set(val *EnvironmentVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariables(val *EnvironmentVariables) *NullableEnvironmentVariables {
	return &NullableEnvironmentVariables{value: val, isSet: true}
}

func (v NullableEnvironmentVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


