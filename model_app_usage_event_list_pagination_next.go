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

// checks if the AppUsageEventListPaginationNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUsageEventListPaginationNext{}

// AppUsageEventListPaginationNext struct for AppUsageEventListPaginationNext
type AppUsageEventListPaginationNext struct {
	// Link to the next page
	Href *string `json:"href,omitempty"`
}

// NewAppUsageEventListPaginationNext instantiates a new AppUsageEventListPaginationNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUsageEventListPaginationNext() *AppUsageEventListPaginationNext {
	this := AppUsageEventListPaginationNext{}
	return &this
}

// NewAppUsageEventListPaginationNextWithDefaults instantiates a new AppUsageEventListPaginationNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUsageEventListPaginationNextWithDefaults() *AppUsageEventListPaginationNext {
	this := AppUsageEventListPaginationNext{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AppUsageEventListPaginationNext) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPaginationNext) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AppUsageEventListPaginationNext) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AppUsageEventListPaginationNext) SetHref(v string) {
	o.Href = &v
}

func (o AppUsageEventListPaginationNext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUsageEventListPaginationNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableAppUsageEventListPaginationNext struct {
	value *AppUsageEventListPaginationNext
	isSet bool
}

func (v NullableAppUsageEventListPaginationNext) Get() *AppUsageEventListPaginationNext {
	return v.value
}

func (v *NullableAppUsageEventListPaginationNext) Set(val *AppUsageEventListPaginationNext) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUsageEventListPaginationNext) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUsageEventListPaginationNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUsageEventListPaginationNext(val *AppUsageEventListPaginationNext) *NullableAppUsageEventListPaginationNext {
	return &NullableAppUsageEventListPaginationNext{value: val, isSet: true}
}

func (v NullableAppUsageEventListPaginationNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUsageEventListPaginationNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

