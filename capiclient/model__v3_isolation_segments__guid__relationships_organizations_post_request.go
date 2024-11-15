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

// checks if the V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest{}

// V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest struct for V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest
type V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest struct {
	Data []V3AppsPostRequestRelationshipsSpaceData `json:"data,omitempty"`
}

// NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest instantiates a new V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest() *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest {
	this := V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest{}
	return &this
}

// NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequestWithDefaults instantiates a new V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequestWithDefaults() *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest {
	this := V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) GetData() []V3AppsPostRequestRelationshipsSpaceData {
	if o == nil || IsNil(o.Data) {
		var ret []V3AppsPostRequestRelationshipsSpaceData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) GetDataOk() ([]V3AppsPostRequestRelationshipsSpaceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []V3AppsPostRequestRelationshipsSpaceData and assigns it to the Data field.
func (o *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) SetData(v []V3AppsPostRequestRelationshipsSpaceData) {
	o.Data = v
}

func (o V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest struct {
	value *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest
	isSet bool
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) Get() *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest {
	return v.value
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) Set(val *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest(val *V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest {
	return &NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest{value: val, isSet: true}
}

func (v NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

