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

// checks if the Get200ResponseLinksCloudControllerV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Get200ResponseLinksCloudControllerV2{}

// Get200ResponseLinksCloudControllerV2 struct for Get200ResponseLinksCloudControllerV2
type Get200ResponseLinksCloudControllerV2 struct {
	Href *string `json:"href,omitempty"`
	Meta *Get200ResponseLinksCloudControllerV2Meta `json:"meta,omitempty"`
}

// NewGet200ResponseLinksCloudControllerV2 instantiates a new Get200ResponseLinksCloudControllerV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGet200ResponseLinksCloudControllerV2() *Get200ResponseLinksCloudControllerV2 {
	this := Get200ResponseLinksCloudControllerV2{}
	return &this
}

// NewGet200ResponseLinksCloudControllerV2WithDefaults instantiates a new Get200ResponseLinksCloudControllerV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGet200ResponseLinksCloudControllerV2WithDefaults() *Get200ResponseLinksCloudControllerV2 {
	this := Get200ResponseLinksCloudControllerV2{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Get200ResponseLinksCloudControllerV2) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Get200ResponseLinksCloudControllerV2) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Get200ResponseLinksCloudControllerV2) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Get200ResponseLinksCloudControllerV2) SetHref(v string) {
	o.Href = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Get200ResponseLinksCloudControllerV2) GetMeta() Get200ResponseLinksCloudControllerV2Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Get200ResponseLinksCloudControllerV2Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Get200ResponseLinksCloudControllerV2) GetMetaOk() (*Get200ResponseLinksCloudControllerV2Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Get200ResponseLinksCloudControllerV2) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Get200ResponseLinksCloudControllerV2Meta and assigns it to the Meta field.
func (o *Get200ResponseLinksCloudControllerV2) SetMeta(v Get200ResponseLinksCloudControllerV2Meta) {
	o.Meta = &v
}

func (o Get200ResponseLinksCloudControllerV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Get200ResponseLinksCloudControllerV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGet200ResponseLinksCloudControllerV2 struct {
	value *Get200ResponseLinksCloudControllerV2
	isSet bool
}

func (v NullableGet200ResponseLinksCloudControllerV2) Get() *Get200ResponseLinksCloudControllerV2 {
	return v.value
}

func (v *NullableGet200ResponseLinksCloudControllerV2) Set(val *Get200ResponseLinksCloudControllerV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGet200ResponseLinksCloudControllerV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGet200ResponseLinksCloudControllerV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGet200ResponseLinksCloudControllerV2(val *Get200ResponseLinksCloudControllerV2) *NullableGet200ResponseLinksCloudControllerV2 {
	return &NullableGet200ResponseLinksCloudControllerV2{value: val, isSet: true}
}

func (v NullableGet200ResponseLinksCloudControllerV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGet200ResponseLinksCloudControllerV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


