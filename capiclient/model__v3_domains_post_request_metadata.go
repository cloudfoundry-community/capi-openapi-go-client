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

// checks if the V3DomainsPostRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3DomainsPostRequestMetadata{}

// V3DomainsPostRequestMetadata struct for V3DomainsPostRequestMetadata
type V3DomainsPostRequestMetadata struct {
	// Annotations applied to the domain
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Labels applied to the domain
	Labels map[string]interface{} `json:"labels,omitempty"`
}

// NewV3DomainsPostRequestMetadata instantiates a new V3DomainsPostRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3DomainsPostRequestMetadata() *V3DomainsPostRequestMetadata {
	this := V3DomainsPostRequestMetadata{}
	return &this
}

// NewV3DomainsPostRequestMetadataWithDefaults instantiates a new V3DomainsPostRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3DomainsPostRequestMetadataWithDefaults() *V3DomainsPostRequestMetadata {
	this := V3DomainsPostRequestMetadata{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V3DomainsPostRequestMetadata) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsPostRequestMetadata) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V3DomainsPostRequestMetadata) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *V3DomainsPostRequestMetadata) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *V3DomainsPostRequestMetadata) GetLabels() map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsPostRequestMetadata) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *V3DomainsPostRequestMetadata) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *V3DomainsPostRequestMetadata) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

func (o V3DomainsPostRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3DomainsPostRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableV3DomainsPostRequestMetadata struct {
	value *V3DomainsPostRequestMetadata
	isSet bool
}

func (v NullableV3DomainsPostRequestMetadata) Get() *V3DomainsPostRequestMetadata {
	return v.value
}

func (v *NullableV3DomainsPostRequestMetadata) Set(val *V3DomainsPostRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableV3DomainsPostRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableV3DomainsPostRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3DomainsPostRequestMetadata(val *V3DomainsPostRequestMetadata) *NullableV3DomainsPostRequestMetadata {
	return &NullableV3DomainsPostRequestMetadata{value: val, isSet: true}
}

func (v NullableV3DomainsPostRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3DomainsPostRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

