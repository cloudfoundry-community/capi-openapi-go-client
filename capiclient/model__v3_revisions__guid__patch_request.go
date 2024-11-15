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

// checks if the V3RevisionsGuidPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3RevisionsGuidPatchRequest{}

// V3RevisionsGuidPatchRequest struct for V3RevisionsGuidPatchRequest
type V3RevisionsGuidPatchRequest struct {
	Metadata *V3AppsGuidTasksPostRequestMetadata `json:"metadata,omitempty"`
}

// NewV3RevisionsGuidPatchRequest instantiates a new V3RevisionsGuidPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3RevisionsGuidPatchRequest() *V3RevisionsGuidPatchRequest {
	this := V3RevisionsGuidPatchRequest{}
	return &this
}

// NewV3RevisionsGuidPatchRequestWithDefaults instantiates a new V3RevisionsGuidPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3RevisionsGuidPatchRequestWithDefaults() *V3RevisionsGuidPatchRequest {
	this := V3RevisionsGuidPatchRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V3RevisionsGuidPatchRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidTasksPostRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3RevisionsGuidPatchRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V3RevisionsGuidPatchRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidTasksPostRequestMetadata and assigns it to the Metadata field.
func (o *V3RevisionsGuidPatchRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata) {
	o.Metadata = &v
}

func (o V3RevisionsGuidPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3RevisionsGuidPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableV3RevisionsGuidPatchRequest struct {
	value *V3RevisionsGuidPatchRequest
	isSet bool
}

func (v NullableV3RevisionsGuidPatchRequest) Get() *V3RevisionsGuidPatchRequest {
	return v.value
}

func (v *NullableV3RevisionsGuidPatchRequest) Set(val *V3RevisionsGuidPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3RevisionsGuidPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3RevisionsGuidPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3RevisionsGuidPatchRequest(val *V3RevisionsGuidPatchRequest) *NullableV3RevisionsGuidPatchRequest {
	return &NullableV3RevisionsGuidPatchRequest{value: val, isSet: true}
}

func (v NullableV3RevisionsGuidPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3RevisionsGuidPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

