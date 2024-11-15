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

// checks if the V3DomainsPostRequestOrganizationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3DomainsPostRequestOrganizationData{}

// V3DomainsPostRequestOrganizationData struct for V3DomainsPostRequestOrganizationData
type V3DomainsPostRequestOrganizationData struct {
	// The organization guid
	Guid *string `json:"guid,omitempty"`
}

// NewV3DomainsPostRequestOrganizationData instantiates a new V3DomainsPostRequestOrganizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3DomainsPostRequestOrganizationData() *V3DomainsPostRequestOrganizationData {
	this := V3DomainsPostRequestOrganizationData{}
	return &this
}

// NewV3DomainsPostRequestOrganizationDataWithDefaults instantiates a new V3DomainsPostRequestOrganizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3DomainsPostRequestOrganizationDataWithDefaults() *V3DomainsPostRequestOrganizationData {
	this := V3DomainsPostRequestOrganizationData{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *V3DomainsPostRequestOrganizationData) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsPostRequestOrganizationData) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *V3DomainsPostRequestOrganizationData) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *V3DomainsPostRequestOrganizationData) SetGuid(v string) {
	o.Guid = &v
}

func (o V3DomainsPostRequestOrganizationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3DomainsPostRequestOrganizationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	return toSerialize, nil
}

type NullableV3DomainsPostRequestOrganizationData struct {
	value *V3DomainsPostRequestOrganizationData
	isSet bool
}

func (v NullableV3DomainsPostRequestOrganizationData) Get() *V3DomainsPostRequestOrganizationData {
	return v.value
}

func (v *NullableV3DomainsPostRequestOrganizationData) Set(val *V3DomainsPostRequestOrganizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableV3DomainsPostRequestOrganizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableV3DomainsPostRequestOrganizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3DomainsPostRequestOrganizationData(val *V3DomainsPostRequestOrganizationData) *NullableV3DomainsPostRequestOrganizationData {
	return &NullableV3DomainsPostRequestOrganizationData{value: val, isSet: true}
}

func (v NullableV3DomainsPostRequestOrganizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3DomainsPostRequestOrganizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


