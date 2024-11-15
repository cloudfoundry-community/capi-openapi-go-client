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

// checks if the AuditEventLinksSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEventLinksSelf{}

// AuditEventLinksSelf struct for AuditEventLinksSelf
type AuditEventLinksSelf struct {
	// Link to the audit event resource
	Href *string `json:"href,omitempty"`
}

// NewAuditEventLinksSelf instantiates a new AuditEventLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEventLinksSelf() *AuditEventLinksSelf {
	this := AuditEventLinksSelf{}
	return &this
}

// NewAuditEventLinksSelfWithDefaults instantiates a new AuditEventLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEventLinksSelfWithDefaults() *AuditEventLinksSelf {
	this := AuditEventLinksSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AuditEventLinksSelf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventLinksSelf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AuditEventLinksSelf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AuditEventLinksSelf) SetHref(v string) {
	o.Href = &v
}

func (o AuditEventLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEventLinksSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableAuditEventLinksSelf struct {
	value *AuditEventLinksSelf
	isSet bool
}

func (v NullableAuditEventLinksSelf) Get() *AuditEventLinksSelf {
	return v.value
}

func (v *NullableAuditEventLinksSelf) Set(val *AuditEventLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEventLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEventLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEventLinksSelf(val *AuditEventLinksSelf) *NullableAuditEventLinksSelf {
	return &NullableAuditEventLinksSelf{value: val, isSet: true}
}

func (v NullableAuditEventLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEventLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


