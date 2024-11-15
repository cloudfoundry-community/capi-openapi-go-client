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

// checks if the V3AppsGuidPatchRequestLifecycleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3AppsGuidPatchRequestLifecycleData{}

// V3AppsGuidPatchRequestLifecycleData struct for V3AppsGuidPatchRequestLifecycleData
type V3AppsGuidPatchRequestLifecycleData struct {
	Buildpacks []string `json:"buildpacks,omitempty"`
	Stack *string `json:"stack,omitempty"`
}

// NewV3AppsGuidPatchRequestLifecycleData instantiates a new V3AppsGuidPatchRequestLifecycleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AppsGuidPatchRequestLifecycleData() *V3AppsGuidPatchRequestLifecycleData {
	this := V3AppsGuidPatchRequestLifecycleData{}
	return &this
}

// NewV3AppsGuidPatchRequestLifecycleDataWithDefaults instantiates a new V3AppsGuidPatchRequestLifecycleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AppsGuidPatchRequestLifecycleDataWithDefaults() *V3AppsGuidPatchRequestLifecycleData {
	this := V3AppsGuidPatchRequestLifecycleData{}
	return &this
}

// GetBuildpacks returns the Buildpacks field value if set, zero value otherwise.
func (o *V3AppsGuidPatchRequestLifecycleData) GetBuildpacks() []string {
	if o == nil || IsNil(o.Buildpacks) {
		var ret []string
		return ret
	}
	return o.Buildpacks
}

// GetBuildpacksOk returns a tuple with the Buildpacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidPatchRequestLifecycleData) GetBuildpacksOk() ([]string, bool) {
	if o == nil || IsNil(o.Buildpacks) {
		return nil, false
	}
	return o.Buildpacks, true
}

// HasBuildpacks returns a boolean if a field has been set.
func (o *V3AppsGuidPatchRequestLifecycleData) HasBuildpacks() bool {
	if o != nil && !IsNil(o.Buildpacks) {
		return true
	}

	return false
}

// SetBuildpacks gets a reference to the given []string and assigns it to the Buildpacks field.
func (o *V3AppsGuidPatchRequestLifecycleData) SetBuildpacks(v []string) {
	o.Buildpacks = v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *V3AppsGuidPatchRequestLifecycleData) GetStack() string {
	if o == nil || IsNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidPatchRequestLifecycleData) GetStackOk() (*string, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *V3AppsGuidPatchRequestLifecycleData) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *V3AppsGuidPatchRequestLifecycleData) SetStack(v string) {
	o.Stack = &v
}

func (o V3AppsGuidPatchRequestLifecycleData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3AppsGuidPatchRequestLifecycleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buildpacks) {
		toSerialize["buildpacks"] = o.Buildpacks
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	return toSerialize, nil
}

type NullableV3AppsGuidPatchRequestLifecycleData struct {
	value *V3AppsGuidPatchRequestLifecycleData
	isSet bool
}

func (v NullableV3AppsGuidPatchRequestLifecycleData) Get() *V3AppsGuidPatchRequestLifecycleData {
	return v.value
}

func (v *NullableV3AppsGuidPatchRequestLifecycleData) Set(val *V3AppsGuidPatchRequestLifecycleData) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AppsGuidPatchRequestLifecycleData) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AppsGuidPatchRequestLifecycleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AppsGuidPatchRequestLifecycleData(val *V3AppsGuidPatchRequestLifecycleData) *NullableV3AppsGuidPatchRequestLifecycleData {
	return &NullableV3AppsGuidPatchRequestLifecycleData{value: val, isSet: true}
}

func (v NullableV3AppsGuidPatchRequestLifecycleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AppsGuidPatchRequestLifecycleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


