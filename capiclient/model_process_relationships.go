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

// checks if the ProcessRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessRelationships{}

// ProcessRelationships struct for ProcessRelationships
type ProcessRelationships struct {
	App *ToOneRelationship `json:"app,omitempty"`
	Revision *ToOneRelationship `json:"revision,omitempty"`
}

// NewProcessRelationships instantiates a new ProcessRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessRelationships() *ProcessRelationships {
	this := ProcessRelationships{}
	return &this
}

// NewProcessRelationshipsWithDefaults instantiates a new ProcessRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessRelationshipsWithDefaults() *ProcessRelationships {
	this := ProcessRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ProcessRelationships) GetApp() ToOneRelationship {
	if o == nil || IsNil(o.App) {
		var ret ToOneRelationship
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRelationships) GetAppOk() (*ToOneRelationship, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ProcessRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ToOneRelationship and assigns it to the App field.
func (o *ProcessRelationships) SetApp(v ToOneRelationship) {
	o.App = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ProcessRelationships) GetRevision() ToOneRelationship {
	if o == nil || IsNil(o.Revision) {
		var ret ToOneRelationship
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRelationships) GetRevisionOk() (*ToOneRelationship, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ProcessRelationships) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given ToOneRelationship and assigns it to the Revision field.
func (o *ProcessRelationships) SetRevision(v ToOneRelationship) {
	o.Revision = &v
}

func (o ProcessRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	return toSerialize, nil
}

type NullableProcessRelationships struct {
	value *ProcessRelationships
	isSet bool
}

func (v NullableProcessRelationships) Get() *ProcessRelationships {
	return v.value
}

func (v *NullableProcessRelationships) Set(val *ProcessRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessRelationships(val *ProcessRelationships) *NullableProcessRelationships {
	return &NullableProcessRelationships{value: val, isSet: true}
}

func (v NullableProcessRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


