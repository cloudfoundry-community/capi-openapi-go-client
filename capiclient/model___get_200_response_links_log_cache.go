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

// checks if the Get200ResponseLinksLogCache type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Get200ResponseLinksLogCache{}

// Get200ResponseLinksLogCache struct for Get200ResponseLinksLogCache
type Get200ResponseLinksLogCache struct {
	Href *string `json:"href,omitempty"`
}

// NewGet200ResponseLinksLogCache instantiates a new Get200ResponseLinksLogCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGet200ResponseLinksLogCache() *Get200ResponseLinksLogCache {
	this := Get200ResponseLinksLogCache{}
	return &this
}

// NewGet200ResponseLinksLogCacheWithDefaults instantiates a new Get200ResponseLinksLogCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGet200ResponseLinksLogCacheWithDefaults() *Get200ResponseLinksLogCache {
	this := Get200ResponseLinksLogCache{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Get200ResponseLinksLogCache) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Get200ResponseLinksLogCache) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Get200ResponseLinksLogCache) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Get200ResponseLinksLogCache) SetHref(v string) {
	o.Href = &v
}

func (o Get200ResponseLinksLogCache) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Get200ResponseLinksLogCache) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableGet200ResponseLinksLogCache struct {
	value *Get200ResponseLinksLogCache
	isSet bool
}

func (v NullableGet200ResponseLinksLogCache) Get() *Get200ResponseLinksLogCache {
	return v.value
}

func (v *NullableGet200ResponseLinksLogCache) Set(val *Get200ResponseLinksLogCache) {
	v.value = val
	v.isSet = true
}

func (v NullableGet200ResponseLinksLogCache) IsSet() bool {
	return v.isSet
}

func (v *NullableGet200ResponseLinksLogCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGet200ResponseLinksLogCache(val *Get200ResponseLinksLogCache) *NullableGet200ResponseLinksLogCache {
	return &NullableGet200ResponseLinksLogCache{value: val, isSet: true}
}

func (v NullableGet200ResponseLinksLogCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGet200ResponseLinksLogCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

