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

// checks if the ProcessScale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessScale{}

// ProcessScale struct for ProcessScale
type ProcessScale struct {
	DiskInMb *int32 `json:"disk_in_mb,omitempty"`
	Instances *int32 `json:"instances,omitempty"`
	LogRateLimitInBytesPerSecond *int32 `json:"log_rate_limit_in_bytes_per_second,omitempty"`
	MemoryInMb *int32 `json:"memory_in_mb,omitempty"`
}

// NewProcessScale instantiates a new ProcessScale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessScale() *ProcessScale {
	this := ProcessScale{}
	return &this
}

// NewProcessScaleWithDefaults instantiates a new ProcessScale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessScaleWithDefaults() *ProcessScale {
	this := ProcessScale{}
	return &this
}

// GetDiskInMb returns the DiskInMb field value if set, zero value otherwise.
func (o *ProcessScale) GetDiskInMb() int32 {
	if o == nil || IsNil(o.DiskInMb) {
		var ret int32
		return ret
	}
	return *o.DiskInMb
}

// GetDiskInMbOk returns a tuple with the DiskInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessScale) GetDiskInMbOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskInMb) {
		return nil, false
	}
	return o.DiskInMb, true
}

// HasDiskInMb returns a boolean if a field has been set.
func (o *ProcessScale) HasDiskInMb() bool {
	if o != nil && !IsNil(o.DiskInMb) {
		return true
	}

	return false
}

// SetDiskInMb gets a reference to the given int32 and assigns it to the DiskInMb field.
func (o *ProcessScale) SetDiskInMb(v int32) {
	o.DiskInMb = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ProcessScale) GetInstances() int32 {
	if o == nil || IsNil(o.Instances) {
		var ret int32
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessScale) GetInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ProcessScale) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given int32 and assigns it to the Instances field.
func (o *ProcessScale) SetInstances(v int32) {
	o.Instances = &v
}

// GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field value if set, zero value otherwise.
func (o *ProcessScale) GetLogRateLimitInBytesPerSecond() int32 {
	if o == nil || IsNil(o.LogRateLimitInBytesPerSecond) {
		var ret int32
		return ret
	}
	return *o.LogRateLimitInBytesPerSecond
}

// GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessScale) GetLogRateLimitInBytesPerSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.LogRateLimitInBytesPerSecond) {
		return nil, false
	}
	return o.LogRateLimitInBytesPerSecond, true
}

// HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.
func (o *ProcessScale) HasLogRateLimitInBytesPerSecond() bool {
	if o != nil && !IsNil(o.LogRateLimitInBytesPerSecond) {
		return true
	}

	return false
}

// SetLogRateLimitInBytesPerSecond gets a reference to the given int32 and assigns it to the LogRateLimitInBytesPerSecond field.
func (o *ProcessScale) SetLogRateLimitInBytesPerSecond(v int32) {
	o.LogRateLimitInBytesPerSecond = &v
}

// GetMemoryInMb returns the MemoryInMb field value if set, zero value otherwise.
func (o *ProcessScale) GetMemoryInMb() int32 {
	if o == nil || IsNil(o.MemoryInMb) {
		var ret int32
		return ret
	}
	return *o.MemoryInMb
}

// GetMemoryInMbOk returns a tuple with the MemoryInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessScale) GetMemoryInMbOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryInMb) {
		return nil, false
	}
	return o.MemoryInMb, true
}

// HasMemoryInMb returns a boolean if a field has been set.
func (o *ProcessScale) HasMemoryInMb() bool {
	if o != nil && !IsNil(o.MemoryInMb) {
		return true
	}

	return false
}

// SetMemoryInMb gets a reference to the given int32 and assigns it to the MemoryInMb field.
func (o *ProcessScale) SetMemoryInMb(v int32) {
	o.MemoryInMb = &v
}

func (o ProcessScale) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessScale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiskInMb) {
		toSerialize["disk_in_mb"] = o.DiskInMb
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.LogRateLimitInBytesPerSecond) {
		toSerialize["log_rate_limit_in_bytes_per_second"] = o.LogRateLimitInBytesPerSecond
	}
	if !IsNil(o.MemoryInMb) {
		toSerialize["memory_in_mb"] = o.MemoryInMb
	}
	return toSerialize, nil
}

type NullableProcessScale struct {
	value *ProcessScale
	isSet bool
}

func (v NullableProcessScale) Get() *ProcessScale {
	return v.value
}

func (v *NullableProcessScale) Set(val *ProcessScale) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessScale) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessScale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessScale(val *ProcessScale) *NullableProcessScale {
	return &NullableProcessScale{value: val, isSet: true}
}

func (v NullableProcessScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessScale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


