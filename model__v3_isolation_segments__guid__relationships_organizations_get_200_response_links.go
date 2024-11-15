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

// checks if the V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks{}

// V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks struct for V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks
type V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks struct {
	Related *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"related,omitempty"`
	Self *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"self,omitempty"`
}

// NewV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks instantiates a new V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks() *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks {
	this := V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks{}
	return &this
}

// NewV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksWithDefaults instantiates a new V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksWithDefaults() *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks {
	this := V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks{}
	return &this
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) GetRelated() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Related) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) GetRelatedOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Related field.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) SetRelated(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Related = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) GetSelf() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Self) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) GetSelfOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Self field.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) SetSelf(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Self = &v
}

func (o V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks struct {
	value *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks
	isSet bool
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) Get() *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks {
	return v.value
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) Set(val *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks(val *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks {
	return &NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks{value: val, isSet: true}
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

