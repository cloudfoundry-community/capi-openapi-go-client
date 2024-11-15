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

// checks if the AppUsageEventSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUsageEventSpace{}

// AppUsageEventSpace struct for AppUsageEventSpace
type AppUsageEventSpace struct {
	// Unique identifier of the space that this event pertains to, if applicable
	Guid NullableString `json:"guid,omitempty"`
	// Name of the space that this event pertains to, if applicable
	Name NullableString `json:"name,omitempty"`
}

// NewAppUsageEventSpace instantiates a new AppUsageEventSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUsageEventSpace() *AppUsageEventSpace {
	this := AppUsageEventSpace{}
	return &this
}

// NewAppUsageEventSpaceWithDefaults instantiates a new AppUsageEventSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUsageEventSpaceWithDefaults() *AppUsageEventSpace {
	this := AppUsageEventSpace{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUsageEventSpace) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUsageEventSpace) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *AppUsageEventSpace) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *AppUsageEventSpace) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *AppUsageEventSpace) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *AppUsageEventSpace) UnsetGuid() {
	o.Guid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUsageEventSpace) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUsageEventSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AppUsageEventSpace) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AppUsageEventSpace) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AppUsageEventSpace) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AppUsageEventSpace) UnsetName() {
	o.Name.Unset()
}

func (o AppUsageEventSpace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUsageEventSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableAppUsageEventSpace struct {
	value *AppUsageEventSpace
	isSet bool
}

func (v NullableAppUsageEventSpace) Get() *AppUsageEventSpace {
	return v.value
}

func (v *NullableAppUsageEventSpace) Set(val *AppUsageEventSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUsageEventSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUsageEventSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUsageEventSpace(val *AppUsageEventSpace) *NullableAppUsageEventSpace {
	return &NullableAppUsageEventSpace{value: val, isSet: true}
}

func (v NullableAppUsageEventSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUsageEventSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

