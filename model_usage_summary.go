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

// checks if the UsageSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageSummary{}

// UsageSummary struct for UsageSummary
type UsageSummary struct {
	Links *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks `json:"links,omitempty"`
	UsageSummary *[]map[string]interface{} `json:"usage_summary,omitempty"`
}

// NewUsageSummary instantiates a new UsageSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSummary() *UsageSummary {
	this := UsageSummary{}
	return &this
}

// NewUsageSummaryWithDefaults instantiates a new UsageSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSummaryWithDefaults() *UsageSummary {
	this := UsageSummary{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UsageSummary) GetLinks() V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummary) GetLinksOk() (*V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UsageSummary) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks and assigns it to the Links field.
func (o *UsageSummary) SetLinks(v V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) {
	o.Links = &v
}

// GetUsageSummary returns the UsageSummary field value if set, zero value otherwise.
func (o *UsageSummary) GetUsageSummary() []map[string]interface{} {
	if o == nil || IsNil(o.UsageSummary) {
		var ret []map[string]interface{}
		return ret
	}
	return *o.UsageSummary
}

// GetUsageSummaryOk returns a tuple with the UsageSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummary) GetUsageSummaryOk() (*[]map[string]interface{}, bool) {
	if o == nil || IsNil(o.UsageSummary) {
		return nil, false
	}
	return o.UsageSummary, true
}

// HasUsageSummary returns a boolean if a field has been set.
func (o *UsageSummary) HasUsageSummary() bool {
	if o != nil && !IsNil(o.UsageSummary) {
		return true
	}

	return false
}

// SetUsageSummary gets a reference to the given []map[string]interface{} and assigns it to the UsageSummary field.
func (o *UsageSummary) SetUsageSummary(v []map[string]interface{}) {
	o.UsageSummary = &v
}

func (o UsageSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.UsageSummary) {
		toSerialize["usage_summary"] = o.UsageSummary
	}
	return toSerialize, nil
}

type NullableUsageSummary struct {
	value *UsageSummary
	isSet bool
}

func (v NullableUsageSummary) Get() *UsageSummary {
	return v.value
}

func (v *NullableUsageSummary) Set(val *UsageSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSummary(val *UsageSummary) *NullableUsageSummary {
	return &NullableUsageSummary{value: val, isSet: true}
}

func (v NullableUsageSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

