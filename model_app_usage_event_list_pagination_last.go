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

// checks if the AppUsageEventListPaginationLast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUsageEventListPaginationLast{}

// AppUsageEventListPaginationLast struct for AppUsageEventListPaginationLast
type AppUsageEventListPaginationLast struct {
	// Link to the last page
	Href *string `json:"href,omitempty"`
}

// NewAppUsageEventListPaginationLast instantiates a new AppUsageEventListPaginationLast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUsageEventListPaginationLast() *AppUsageEventListPaginationLast {
	this := AppUsageEventListPaginationLast{}
	return &this
}

// NewAppUsageEventListPaginationLastWithDefaults instantiates a new AppUsageEventListPaginationLast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUsageEventListPaginationLastWithDefaults() *AppUsageEventListPaginationLast {
	this := AppUsageEventListPaginationLast{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AppUsageEventListPaginationLast) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPaginationLast) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AppUsageEventListPaginationLast) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AppUsageEventListPaginationLast) SetHref(v string) {
	o.Href = &v
}

func (o AppUsageEventListPaginationLast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUsageEventListPaginationLast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableAppUsageEventListPaginationLast struct {
	value *AppUsageEventListPaginationLast
	isSet bool
}

func (v NullableAppUsageEventListPaginationLast) Get() *AppUsageEventListPaginationLast {
	return v.value
}

func (v *NullableAppUsageEventListPaginationLast) Set(val *AppUsageEventListPaginationLast) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUsageEventListPaginationLast) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUsageEventListPaginationLast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUsageEventListPaginationLast(val *AppUsageEventListPaginationLast) *NullableAppUsageEventListPaginationLast {
	return &NullableAppUsageEventListPaginationLast{value: val, isSet: true}
}

func (v NullableAppUsageEventListPaginationLast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUsageEventListPaginationLast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


