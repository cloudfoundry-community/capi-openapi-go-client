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

// checks if the V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks{}

// V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks struct for V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks
type V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks struct {
	Self *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"self,omitempty"`
}

// NewV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks instantiates a new V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks() *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks {
	this := V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks{}
	return &this
}

// NewV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinksWithDefaults instantiates a new V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinksWithDefaults() *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks {
	this := V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) GetSelf() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Self) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) GetSelfOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Self field.
func (o *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) SetSelf(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Self = &v
}

func (o V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks struct {
	value *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks
	isSet bool
}

func (v NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) Get() *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks {
	return v.value
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) Set(val *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks(val *V3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) *NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks {
	return &NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks{value: val, isSet: true}
}

func (v NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsSpacesGet200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

