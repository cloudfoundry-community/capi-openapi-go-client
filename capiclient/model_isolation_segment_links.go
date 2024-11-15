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

// checks if the IsolationSegmentLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsolationSegmentLinks{}

// IsolationSegmentLinks struct for IsolationSegmentLinks
type IsolationSegmentLinks struct {
	Organizations *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"organizations,omitempty"`
	Self *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"self,omitempty"`
}

// NewIsolationSegmentLinks instantiates a new IsolationSegmentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsolationSegmentLinks() *IsolationSegmentLinks {
	this := IsolationSegmentLinks{}
	return &this
}

// NewIsolationSegmentLinksWithDefaults instantiates a new IsolationSegmentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsolationSegmentLinksWithDefaults() *IsolationSegmentLinks {
	this := IsolationSegmentLinks{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *IsolationSegmentLinks) GetOrganizations() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Organizations) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsolationSegmentLinks) GetOrganizationsOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *IsolationSegmentLinks) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Organizations field.
func (o *IsolationSegmentLinks) SetOrganizations(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Organizations = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IsolationSegmentLinks) GetSelf() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Self) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsolationSegmentLinks) GetSelfOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IsolationSegmentLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Self field.
func (o *IsolationSegmentLinks) SetSelf(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Self = &v
}

func (o IsolationSegmentLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsolationSegmentLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableIsolationSegmentLinks struct {
	value *IsolationSegmentLinks
	isSet bool
}

func (v NullableIsolationSegmentLinks) Get() *IsolationSegmentLinks {
	return v.value
}

func (v *NullableIsolationSegmentLinks) Set(val *IsolationSegmentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIsolationSegmentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIsolationSegmentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsolationSegmentLinks(val *IsolationSegmentLinks) *NullableIsolationSegmentLinks {
	return &NullableIsolationSegmentLinks{value: val, isSet: true}
}

func (v NullableIsolationSegmentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsolationSegmentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


