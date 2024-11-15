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

// checks if the V3SpaceQuotasPostRequestApps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3SpaceQuotasPostRequestApps{}

// V3SpaceQuotasPostRequestApps struct for V3SpaceQuotasPostRequestApps
type V3SpaceQuotasPostRequestApps struct {
	LogRateLimitInBytesPerSecond NullableInt32 `json:"log_rate_limit_in_bytes_per_second,omitempty"`
	PerAppTasks NullableInt32 `json:"per_app_tasks,omitempty"`
	PerProcessMemoryInMb NullableInt32 `json:"per_process_memory_in_mb,omitempty"`
	TotalInstances NullableInt32 `json:"total_instances,omitempty"`
	TotalMemoryInMb NullableInt32 `json:"total_memory_in_mb,omitempty"`
}

// NewV3SpaceQuotasPostRequestApps instantiates a new V3SpaceQuotasPostRequestApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3SpaceQuotasPostRequestApps() *V3SpaceQuotasPostRequestApps {
	this := V3SpaceQuotasPostRequestApps{}
	return &this
}

// NewV3SpaceQuotasPostRequestAppsWithDefaults instantiates a new V3SpaceQuotasPostRequestApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3SpaceQuotasPostRequestAppsWithDefaults() *V3SpaceQuotasPostRequestApps {
	this := V3SpaceQuotasPostRequestApps{}
	return &this
}

// GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3SpaceQuotasPostRequestApps) GetLogRateLimitInBytesPerSecond() int32 {
	if o == nil || IsNil(o.LogRateLimitInBytesPerSecond.Get()) {
		var ret int32
		return ret
	}
	return *o.LogRateLimitInBytesPerSecond.Get()
}

// GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3SpaceQuotasPostRequestApps) GetLogRateLimitInBytesPerSecondOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogRateLimitInBytesPerSecond.Get(), o.LogRateLimitInBytesPerSecond.IsSet()
}

// HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestApps) HasLogRateLimitInBytesPerSecond() bool {
	if o != nil && o.LogRateLimitInBytesPerSecond.IsSet() {
		return true
	}

	return false
}

// SetLogRateLimitInBytesPerSecond gets a reference to the given NullableInt32 and assigns it to the LogRateLimitInBytesPerSecond field.
func (o *V3SpaceQuotasPostRequestApps) SetLogRateLimitInBytesPerSecond(v int32) {
	o.LogRateLimitInBytesPerSecond.Set(&v)
}
// SetLogRateLimitInBytesPerSecondNil sets the value for LogRateLimitInBytesPerSecond to be an explicit nil
func (o *V3SpaceQuotasPostRequestApps) SetLogRateLimitInBytesPerSecondNil() {
	o.LogRateLimitInBytesPerSecond.Set(nil)
}

// UnsetLogRateLimitInBytesPerSecond ensures that no value is present for LogRateLimitInBytesPerSecond, not even an explicit nil
func (o *V3SpaceQuotasPostRequestApps) UnsetLogRateLimitInBytesPerSecond() {
	o.LogRateLimitInBytesPerSecond.Unset()
}

// GetPerAppTasks returns the PerAppTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3SpaceQuotasPostRequestApps) GetPerAppTasks() int32 {
	if o == nil || IsNil(o.PerAppTasks.Get()) {
		var ret int32
		return ret
	}
	return *o.PerAppTasks.Get()
}

// GetPerAppTasksOk returns a tuple with the PerAppTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3SpaceQuotasPostRequestApps) GetPerAppTasksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerAppTasks.Get(), o.PerAppTasks.IsSet()
}

// HasPerAppTasks returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestApps) HasPerAppTasks() bool {
	if o != nil && o.PerAppTasks.IsSet() {
		return true
	}

	return false
}

// SetPerAppTasks gets a reference to the given NullableInt32 and assigns it to the PerAppTasks field.
func (o *V3SpaceQuotasPostRequestApps) SetPerAppTasks(v int32) {
	o.PerAppTasks.Set(&v)
}
// SetPerAppTasksNil sets the value for PerAppTasks to be an explicit nil
func (o *V3SpaceQuotasPostRequestApps) SetPerAppTasksNil() {
	o.PerAppTasks.Set(nil)
}

// UnsetPerAppTasks ensures that no value is present for PerAppTasks, not even an explicit nil
func (o *V3SpaceQuotasPostRequestApps) UnsetPerAppTasks() {
	o.PerAppTasks.Unset()
}

// GetPerProcessMemoryInMb returns the PerProcessMemoryInMb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3SpaceQuotasPostRequestApps) GetPerProcessMemoryInMb() int32 {
	if o == nil || IsNil(o.PerProcessMemoryInMb.Get()) {
		var ret int32
		return ret
	}
	return *o.PerProcessMemoryInMb.Get()
}

// GetPerProcessMemoryInMbOk returns a tuple with the PerProcessMemoryInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3SpaceQuotasPostRequestApps) GetPerProcessMemoryInMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerProcessMemoryInMb.Get(), o.PerProcessMemoryInMb.IsSet()
}

// HasPerProcessMemoryInMb returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestApps) HasPerProcessMemoryInMb() bool {
	if o != nil && o.PerProcessMemoryInMb.IsSet() {
		return true
	}

	return false
}

// SetPerProcessMemoryInMb gets a reference to the given NullableInt32 and assigns it to the PerProcessMemoryInMb field.
func (o *V3SpaceQuotasPostRequestApps) SetPerProcessMemoryInMb(v int32) {
	o.PerProcessMemoryInMb.Set(&v)
}
// SetPerProcessMemoryInMbNil sets the value for PerProcessMemoryInMb to be an explicit nil
func (o *V3SpaceQuotasPostRequestApps) SetPerProcessMemoryInMbNil() {
	o.PerProcessMemoryInMb.Set(nil)
}

// UnsetPerProcessMemoryInMb ensures that no value is present for PerProcessMemoryInMb, not even an explicit nil
func (o *V3SpaceQuotasPostRequestApps) UnsetPerProcessMemoryInMb() {
	o.PerProcessMemoryInMb.Unset()
}

// GetTotalInstances returns the TotalInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3SpaceQuotasPostRequestApps) GetTotalInstances() int32 {
	if o == nil || IsNil(o.TotalInstances.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalInstances.Get()
}

// GetTotalInstancesOk returns a tuple with the TotalInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3SpaceQuotasPostRequestApps) GetTotalInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalInstances.Get(), o.TotalInstances.IsSet()
}

// HasTotalInstances returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestApps) HasTotalInstances() bool {
	if o != nil && o.TotalInstances.IsSet() {
		return true
	}

	return false
}

// SetTotalInstances gets a reference to the given NullableInt32 and assigns it to the TotalInstances field.
func (o *V3SpaceQuotasPostRequestApps) SetTotalInstances(v int32) {
	o.TotalInstances.Set(&v)
}
// SetTotalInstancesNil sets the value for TotalInstances to be an explicit nil
func (o *V3SpaceQuotasPostRequestApps) SetTotalInstancesNil() {
	o.TotalInstances.Set(nil)
}

// UnsetTotalInstances ensures that no value is present for TotalInstances, not even an explicit nil
func (o *V3SpaceQuotasPostRequestApps) UnsetTotalInstances() {
	o.TotalInstances.Unset()
}

// GetTotalMemoryInMb returns the TotalMemoryInMb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3SpaceQuotasPostRequestApps) GetTotalMemoryInMb() int32 {
	if o == nil || IsNil(o.TotalMemoryInMb.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalMemoryInMb.Get()
}

// GetTotalMemoryInMbOk returns a tuple with the TotalMemoryInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3SpaceQuotasPostRequestApps) GetTotalMemoryInMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalMemoryInMb.Get(), o.TotalMemoryInMb.IsSet()
}

// HasTotalMemoryInMb returns a boolean if a field has been set.
func (o *V3SpaceQuotasPostRequestApps) HasTotalMemoryInMb() bool {
	if o != nil && o.TotalMemoryInMb.IsSet() {
		return true
	}

	return false
}

// SetTotalMemoryInMb gets a reference to the given NullableInt32 and assigns it to the TotalMemoryInMb field.
func (o *V3SpaceQuotasPostRequestApps) SetTotalMemoryInMb(v int32) {
	o.TotalMemoryInMb.Set(&v)
}
// SetTotalMemoryInMbNil sets the value for TotalMemoryInMb to be an explicit nil
func (o *V3SpaceQuotasPostRequestApps) SetTotalMemoryInMbNil() {
	o.TotalMemoryInMb.Set(nil)
}

// UnsetTotalMemoryInMb ensures that no value is present for TotalMemoryInMb, not even an explicit nil
func (o *V3SpaceQuotasPostRequestApps) UnsetTotalMemoryInMb() {
	o.TotalMemoryInMb.Unset()
}

func (o V3SpaceQuotasPostRequestApps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3SpaceQuotasPostRequestApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LogRateLimitInBytesPerSecond.IsSet() {
		toSerialize["log_rate_limit_in_bytes_per_second"] = o.LogRateLimitInBytesPerSecond.Get()
	}
	if o.PerAppTasks.IsSet() {
		toSerialize["per_app_tasks"] = o.PerAppTasks.Get()
	}
	if o.PerProcessMemoryInMb.IsSet() {
		toSerialize["per_process_memory_in_mb"] = o.PerProcessMemoryInMb.Get()
	}
	if o.TotalInstances.IsSet() {
		toSerialize["total_instances"] = o.TotalInstances.Get()
	}
	if o.TotalMemoryInMb.IsSet() {
		toSerialize["total_memory_in_mb"] = o.TotalMemoryInMb.Get()
	}
	return toSerialize, nil
}

type NullableV3SpaceQuotasPostRequestApps struct {
	value *V3SpaceQuotasPostRequestApps
	isSet bool
}

func (v NullableV3SpaceQuotasPostRequestApps) Get() *V3SpaceQuotasPostRequestApps {
	return v.value
}

func (v *NullableV3SpaceQuotasPostRequestApps) Set(val *V3SpaceQuotasPostRequestApps) {
	v.value = val
	v.isSet = true
}

func (v NullableV3SpaceQuotasPostRequestApps) IsSet() bool {
	return v.isSet
}

func (v *NullableV3SpaceQuotasPostRequestApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3SpaceQuotasPostRequestApps(val *V3SpaceQuotasPostRequestApps) *NullableV3SpaceQuotasPostRequestApps {
	return &NullableV3SpaceQuotasPostRequestApps{value: val, isSet: true}
}

func (v NullableV3SpaceQuotasPostRequestApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3SpaceQuotasPostRequestApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


