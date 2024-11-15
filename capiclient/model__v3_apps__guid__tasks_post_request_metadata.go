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

// checks if the V3AppsGuidTasksPostRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3AppsGuidTasksPostRequestMetadata{}

// V3AppsGuidTasksPostRequestMetadata struct for V3AppsGuidTasksPostRequestMetadata
type V3AppsGuidTasksPostRequestMetadata struct {
	Annotations *map[string]string `json:"annotations,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// NewV3AppsGuidTasksPostRequestMetadata instantiates a new V3AppsGuidTasksPostRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AppsGuidTasksPostRequestMetadata() *V3AppsGuidTasksPostRequestMetadata {
	this := V3AppsGuidTasksPostRequestMetadata{}
	return &this
}

// NewV3AppsGuidTasksPostRequestMetadataWithDefaults instantiates a new V3AppsGuidTasksPostRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AppsGuidTasksPostRequestMetadataWithDefaults() *V3AppsGuidTasksPostRequestMetadata {
	this := V3AppsGuidTasksPostRequestMetadata{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V3AppsGuidTasksPostRequestMetadata) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidTasksPostRequestMetadata) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V3AppsGuidTasksPostRequestMetadata) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *V3AppsGuidTasksPostRequestMetadata) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *V3AppsGuidTasksPostRequestMetadata) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidTasksPostRequestMetadata) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *V3AppsGuidTasksPostRequestMetadata) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *V3AppsGuidTasksPostRequestMetadata) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o V3AppsGuidTasksPostRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3AppsGuidTasksPostRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableV3AppsGuidTasksPostRequestMetadata struct {
	value *V3AppsGuidTasksPostRequestMetadata
	isSet bool
}

func (v NullableV3AppsGuidTasksPostRequestMetadata) Get() *V3AppsGuidTasksPostRequestMetadata {
	return v.value
}

func (v *NullableV3AppsGuidTasksPostRequestMetadata) Set(val *V3AppsGuidTasksPostRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AppsGuidTasksPostRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AppsGuidTasksPostRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AppsGuidTasksPostRequestMetadata(val *V3AppsGuidTasksPostRequestMetadata) *NullableV3AppsGuidTasksPostRequestMetadata {
	return &NullableV3AppsGuidTasksPostRequestMetadata{value: val, isSet: true}
}

func (v NullableV3AppsGuidTasksPostRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AppsGuidTasksPostRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


