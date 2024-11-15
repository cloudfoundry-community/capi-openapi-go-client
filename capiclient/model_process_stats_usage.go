/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
	"time"
)

// checks if the ProcessStatsUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessStatsUsage{}

// ProcessStatsUsage struct for ProcessStatsUsage
type ProcessStatsUsage struct {
	Cpu *float32 `json:"cpu,omitempty"`
	Disk *int32 `json:"disk,omitempty"`
	LogRate *int32 `json:"log_rate,omitempty"`
	Mem *int32 `json:"mem,omitempty"`
	Time *time.Time `json:"time,omitempty"`
}

// NewProcessStatsUsage instantiates a new ProcessStatsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessStatsUsage() *ProcessStatsUsage {
	this := ProcessStatsUsage{}
	return &this
}

// NewProcessStatsUsageWithDefaults instantiates a new ProcessStatsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessStatsUsageWithDefaults() *ProcessStatsUsage {
	this := ProcessStatsUsage{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ProcessStatsUsage) GetCpu() float32 {
	if o == nil || IsNil(o.Cpu) {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStatsUsage) GetCpuOk() (*float32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ProcessStatsUsage) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ProcessStatsUsage) SetCpu(v float32) {
	o.Cpu = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ProcessStatsUsage) GetDisk() int32 {
	if o == nil || IsNil(o.Disk) {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStatsUsage) GetDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ProcessStatsUsage) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *ProcessStatsUsage) SetDisk(v int32) {
	o.Disk = &v
}

// GetLogRate returns the LogRate field value if set, zero value otherwise.
func (o *ProcessStatsUsage) GetLogRate() int32 {
	if o == nil || IsNil(o.LogRate) {
		var ret int32
		return ret
	}
	return *o.LogRate
}

// GetLogRateOk returns a tuple with the LogRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStatsUsage) GetLogRateOk() (*int32, bool) {
	if o == nil || IsNil(o.LogRate) {
		return nil, false
	}
	return o.LogRate, true
}

// HasLogRate returns a boolean if a field has been set.
func (o *ProcessStatsUsage) HasLogRate() bool {
	if o != nil && !IsNil(o.LogRate) {
		return true
	}

	return false
}

// SetLogRate gets a reference to the given int32 and assigns it to the LogRate field.
func (o *ProcessStatsUsage) SetLogRate(v int32) {
	o.LogRate = &v
}

// GetMem returns the Mem field value if set, zero value otherwise.
func (o *ProcessStatsUsage) GetMem() int32 {
	if o == nil || IsNil(o.Mem) {
		var ret int32
		return ret
	}
	return *o.Mem
}

// GetMemOk returns a tuple with the Mem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStatsUsage) GetMemOk() (*int32, bool) {
	if o == nil || IsNil(o.Mem) {
		return nil, false
	}
	return o.Mem, true
}

// HasMem returns a boolean if a field has been set.
func (o *ProcessStatsUsage) HasMem() bool {
	if o != nil && !IsNil(o.Mem) {
		return true
	}

	return false
}

// SetMem gets a reference to the given int32 and assigns it to the Mem field.
func (o *ProcessStatsUsage) SetMem(v int32) {
	o.Mem = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ProcessStatsUsage) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStatsUsage) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ProcessStatsUsage) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ProcessStatsUsage) SetTime(v time.Time) {
	o.Time = &v
}

func (o ProcessStatsUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessStatsUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.LogRate) {
		toSerialize["log_rate"] = o.LogRate
	}
	if !IsNil(o.Mem) {
		toSerialize["mem"] = o.Mem
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableProcessStatsUsage struct {
	value *ProcessStatsUsage
	isSet bool
}

func (v NullableProcessStatsUsage) Get() *ProcessStatsUsage {
	return v.value
}

func (v *NullableProcessStatsUsage) Set(val *ProcessStatsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessStatsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessStatsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessStatsUsage(val *ProcessStatsUsage) *NullableProcessStatsUsage {
	return &NullableProcessStatsUsage{value: val, isSet: true}
}

func (v NullableProcessStatsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessStatsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


