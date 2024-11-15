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

// checks if the V3SidecarsGuidPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3SidecarsGuidPatchRequest{}

// V3SidecarsGuidPatchRequest struct for V3SidecarsGuidPatchRequest
type V3SidecarsGuidPatchRequest struct {
	Command *string `json:"command,omitempty"`
	MemoryInMb *int32 `json:"memory_in_mb,omitempty"`
	Name *string `json:"name,omitempty"`
	ProcessTypes []string `json:"process_types,omitempty"`
}

// NewV3SidecarsGuidPatchRequest instantiates a new V3SidecarsGuidPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3SidecarsGuidPatchRequest() *V3SidecarsGuidPatchRequest {
	this := V3SidecarsGuidPatchRequest{}
	return &this
}

// NewV3SidecarsGuidPatchRequestWithDefaults instantiates a new V3SidecarsGuidPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3SidecarsGuidPatchRequestWithDefaults() *V3SidecarsGuidPatchRequest {
	this := V3SidecarsGuidPatchRequest{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V3SidecarsGuidPatchRequest) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3SidecarsGuidPatchRequest) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V3SidecarsGuidPatchRequest) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *V3SidecarsGuidPatchRequest) SetCommand(v string) {
	o.Command = &v
}

// GetMemoryInMb returns the MemoryInMb field value if set, zero value otherwise.
func (o *V3SidecarsGuidPatchRequest) GetMemoryInMb() int32 {
	if o == nil || IsNil(o.MemoryInMb) {
		var ret int32
		return ret
	}
	return *o.MemoryInMb
}

// GetMemoryInMbOk returns a tuple with the MemoryInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3SidecarsGuidPatchRequest) GetMemoryInMbOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryInMb) {
		return nil, false
	}
	return o.MemoryInMb, true
}

// HasMemoryInMb returns a boolean if a field has been set.
func (o *V3SidecarsGuidPatchRequest) HasMemoryInMb() bool {
	if o != nil && !IsNil(o.MemoryInMb) {
		return true
	}

	return false
}

// SetMemoryInMb gets a reference to the given int32 and assigns it to the MemoryInMb field.
func (o *V3SidecarsGuidPatchRequest) SetMemoryInMb(v int32) {
	o.MemoryInMb = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V3SidecarsGuidPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3SidecarsGuidPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V3SidecarsGuidPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V3SidecarsGuidPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetProcessTypes returns the ProcessTypes field value if set, zero value otherwise.
func (o *V3SidecarsGuidPatchRequest) GetProcessTypes() []string {
	if o == nil || IsNil(o.ProcessTypes) {
		var ret []string
		return ret
	}
	return o.ProcessTypes
}

// GetProcessTypesOk returns a tuple with the ProcessTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3SidecarsGuidPatchRequest) GetProcessTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessTypes) {
		return nil, false
	}
	return o.ProcessTypes, true
}

// HasProcessTypes returns a boolean if a field has been set.
func (o *V3SidecarsGuidPatchRequest) HasProcessTypes() bool {
	if o != nil && !IsNil(o.ProcessTypes) {
		return true
	}

	return false
}

// SetProcessTypes gets a reference to the given []string and assigns it to the ProcessTypes field.
func (o *V3SidecarsGuidPatchRequest) SetProcessTypes(v []string) {
	o.ProcessTypes = v
}

func (o V3SidecarsGuidPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3SidecarsGuidPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.MemoryInMb) {
		toSerialize["memory_in_mb"] = o.MemoryInMb
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProcessTypes) {
		toSerialize["process_types"] = o.ProcessTypes
	}
	return toSerialize, nil
}

type NullableV3SidecarsGuidPatchRequest struct {
	value *V3SidecarsGuidPatchRequest
	isSet bool
}

func (v NullableV3SidecarsGuidPatchRequest) Get() *V3SidecarsGuidPatchRequest {
	return v.value
}

func (v *NullableV3SidecarsGuidPatchRequest) Set(val *V3SidecarsGuidPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3SidecarsGuidPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3SidecarsGuidPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3SidecarsGuidPatchRequest(val *V3SidecarsGuidPatchRequest) *NullableV3SidecarsGuidPatchRequest {
	return &NullableV3SidecarsGuidPatchRequest{value: val, isSet: true}
}

func (v NullableV3SidecarsGuidPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3SidecarsGuidPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


