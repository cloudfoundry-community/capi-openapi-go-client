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

// checks if the PackageLinksSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageLinksSelf{}

// PackageLinksSelf struct for PackageLinksSelf
type PackageLinksSelf struct {
	// URL of the package
	Href *string `json:"href,omitempty"`
}

// NewPackageLinksSelf instantiates a new PackageLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLinksSelf() *PackageLinksSelf {
	this := PackageLinksSelf{}
	return &this
}

// NewPackageLinksSelfWithDefaults instantiates a new PackageLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLinksSelfWithDefaults() *PackageLinksSelf {
	this := PackageLinksSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PackageLinksSelf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLinksSelf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PackageLinksSelf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PackageLinksSelf) SetHref(v string) {
	o.Href = &v
}

func (o PackageLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageLinksSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullablePackageLinksSelf struct {
	value *PackageLinksSelf
	isSet bool
}

func (v NullablePackageLinksSelf) Get() *PackageLinksSelf {
	return v.value
}

func (v *NullablePackageLinksSelf) Set(val *PackageLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLinksSelf(val *PackageLinksSelf) *NullablePackageLinksSelf {
	return &NullablePackageLinksSelf{value: val, isSet: true}
}

func (v NullablePackageLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

