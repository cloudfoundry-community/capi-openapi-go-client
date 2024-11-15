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

// checks if the V3SpaceQuotasPostRequestRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3SpaceQuotasPostRequestRelationships{}

// V3SpaceQuotasPostRequestRelationships struct for V3SpaceQuotasPostRequestRelationships
type V3SpaceQuotasPostRequestRelationships struct {
	Organization *V3AppsPostRequestRelationshipsSpace `json:"organization,omitempty"`
}

// NewV3SpaceQuotasPostRequestRelationships instantiates a new V3SpaceQuotasPostRequestRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3SpaceQuotasPostRequestRelationships() *V3SpaceQuotasPostRequestRelationships {
	this := V3SpaceQuotasPostRequestRelationships{}
	return &this
}

// NewV3SpaceQuotasPostRequestRelationshipsWithDefaults instantiates a new V3SpaceQuotasPostRequestRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3SpaceQuotasPostRequestRelationshipsWithDefaults() *V3SpaceQuotasPostRequestRelationships {
	this := V3SpaceQuotasPostRequestRelationships{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *V3SpaceQuotasPostRequestRelationships) GetOrganization() V3AppsPostRequestRelationshipsSpace {
	if o == nil || IsNil(o.Organization) {
		var ret V3AppsPostRequestRelationshipsSpace
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3SpaceQuotasPostRequestRelationships) GetOrganizationOk() (*V3AppsPostRequestRelationshipsSpace, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestRelationships) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given V3AppsPostRequestRelationshipsSpace and assigns it to the Organization field.
func (o *V3SpaceQuotasPostRequestRelationships) SetOrganization(v V3AppsPostRequestRelationshipsSpace) {
	o.Organization = &v
}

func (o V3SpaceQuotasPostRequestRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3SpaceQuotasPostRequestRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableV3SpaceQuotasPostRequestRelationships struct {
	value *V3SpaceQuotasPostRequestRelationships
	isSet bool
}

func (v NullableV3SpaceQuotasPostRequestRelationships) Get() *V3SpaceQuotasPostRequestRelationships {
	return v.value
}

func (v *NullableV3SpaceQuotasPostRequestRelationships) Set(val *V3SpaceQuotasPostRequestRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableV3SpaceQuotasPostRequestRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableV3SpaceQuotasPostRequestRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3SpaceQuotasPostRequestRelationships(val *V3SpaceQuotasPostRequestRelationships) *NullableV3SpaceQuotasPostRequestRelationships {
	return &NullableV3SpaceQuotasPostRequestRelationships{value: val, isSet: true}
}

func (v NullableV3SpaceQuotasPostRequestRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3SpaceQuotasPostRequestRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


