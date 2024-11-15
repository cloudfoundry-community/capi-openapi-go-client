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

// checks if the ProcessStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessStats{}

// ProcessStats struct for ProcessStats
type ProcessStats struct {
	Details NullableString `json:"details,omitempty"`
	DiskQuota NullableInt32 `json:"disk_quota,omitempty"`
	FdsQuota *int32 `json:"fds_quota,omitempty"`
	Host *string `json:"host,omitempty"`
	Index *int32 `json:"index,omitempty"`
	InstanceInternalIp *string `json:"instance_internal_ip,omitempty"`
	InstancePorts []ProcessStatsInstancePortsInner `json:"instance_ports,omitempty"`
	IsolationSegment NullableString `json:"isolation_segment,omitempty"`
	LogRateLimit NullableInt32 `json:"log_rate_limit,omitempty"`
	MemQuota NullableInt32 `json:"mem_quota,omitempty"`
	State *string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
	Uptime *int32 `json:"uptime,omitempty"`
	Usage *ProcessStatsUsage `json:"usage,omitempty"`
}

// NewProcessStats instantiates a new ProcessStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessStats() *ProcessStats {
	this := ProcessStats{}
	return &this
}

// NewProcessStatsWithDefaults instantiates a new ProcessStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessStatsWithDefaults() *ProcessStats {
	this := ProcessStats{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessStats) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessStats) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *ProcessStats) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *ProcessStats) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *ProcessStats) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *ProcessStats) UnsetDetails() {
	o.Details.Unset()
}

// GetDiskQuota returns the DiskQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessStats) GetDiskQuota() int32 {
	if o == nil || IsNil(o.DiskQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.DiskQuota.Get()
}

// GetDiskQuotaOk returns a tuple with the DiskQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessStats) GetDiskQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskQuota.Get(), o.DiskQuota.IsSet()
}

// HasDiskQuota returns a boolean if a field has been set.
func (o *ProcessStats) HasDiskQuota() bool {
	if o != nil && o.DiskQuota.IsSet() {
		return true
	}

	return false
}

// SetDiskQuota gets a reference to the given NullableInt32 and assigns it to the DiskQuota field.
func (o *ProcessStats) SetDiskQuota(v int32) {
	o.DiskQuota.Set(&v)
}
// SetDiskQuotaNil sets the value for DiskQuota to be an explicit nil
func (o *ProcessStats) SetDiskQuotaNil() {
	o.DiskQuota.Set(nil)
}

// UnsetDiskQuota ensures that no value is present for DiskQuota, not even an explicit nil
func (o *ProcessStats) UnsetDiskQuota() {
	o.DiskQuota.Unset()
}

// GetFdsQuota returns the FdsQuota field value if set, zero value otherwise.
func (o *ProcessStats) GetFdsQuota() int32 {
	if o == nil || IsNil(o.FdsQuota) {
		var ret int32
		return ret
	}
	return *o.FdsQuota
}

// GetFdsQuotaOk returns a tuple with the FdsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetFdsQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.FdsQuota) {
		return nil, false
	}
	return o.FdsQuota, true
}

// HasFdsQuota returns a boolean if a field has been set.
func (o *ProcessStats) HasFdsQuota() bool {
	if o != nil && !IsNil(o.FdsQuota) {
		return true
	}

	return false
}

// SetFdsQuota gets a reference to the given int32 and assigns it to the FdsQuota field.
func (o *ProcessStats) SetFdsQuota(v int32) {
	o.FdsQuota = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ProcessStats) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ProcessStats) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ProcessStats) SetHost(v string) {
	o.Host = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ProcessStats) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ProcessStats) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *ProcessStats) SetIndex(v int32) {
	o.Index = &v
}

// GetInstanceInternalIp returns the InstanceInternalIp field value if set, zero value otherwise.
func (o *ProcessStats) GetInstanceInternalIp() string {
	if o == nil || IsNil(o.InstanceInternalIp) {
		var ret string
		return ret
	}
	return *o.InstanceInternalIp
}

// GetInstanceInternalIpOk returns a tuple with the InstanceInternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetInstanceInternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceInternalIp) {
		return nil, false
	}
	return o.InstanceInternalIp, true
}

// HasInstanceInternalIp returns a boolean if a field has been set.
func (o *ProcessStats) HasInstanceInternalIp() bool {
	if o != nil && !IsNil(o.InstanceInternalIp) {
		return true
	}

	return false
}

// SetInstanceInternalIp gets a reference to the given string and assigns it to the InstanceInternalIp field.
func (o *ProcessStats) SetInstanceInternalIp(v string) {
	o.InstanceInternalIp = &v
}

// GetInstancePorts returns the InstancePorts field value if set, zero value otherwise.
func (o *ProcessStats) GetInstancePorts() []ProcessStatsInstancePortsInner {
	if o == nil || IsNil(o.InstancePorts) {
		var ret []ProcessStatsInstancePortsInner
		return ret
	}
	return o.InstancePorts
}

// GetInstancePortsOk returns a tuple with the InstancePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetInstancePortsOk() ([]ProcessStatsInstancePortsInner, bool) {
	if o == nil || IsNil(o.InstancePorts) {
		return nil, false
	}
	return o.InstancePorts, true
}

// HasInstancePorts returns a boolean if a field has been set.
func (o *ProcessStats) HasInstancePorts() bool {
	if o != nil && !IsNil(o.InstancePorts) {
		return true
	}

	return false
}

// SetInstancePorts gets a reference to the given []ProcessStatsInstancePortsInner and assigns it to the InstancePorts field.
func (o *ProcessStats) SetInstancePorts(v []ProcessStatsInstancePortsInner) {
	o.InstancePorts = v
}

// GetIsolationSegment returns the IsolationSegment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessStats) GetIsolationSegment() string {
	if o == nil || IsNil(o.IsolationSegment.Get()) {
		var ret string
		return ret
	}
	return *o.IsolationSegment.Get()
}

// GetIsolationSegmentOk returns a tuple with the IsolationSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessStats) GetIsolationSegmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsolationSegment.Get(), o.IsolationSegment.IsSet()
}

// HasIsolationSegment returns a boolean if a field has been set.
func (o *ProcessStats) HasIsolationSegment() bool {
	if o != nil && o.IsolationSegment.IsSet() {
		return true
	}

	return false
}

// SetIsolationSegment gets a reference to the given NullableString and assigns it to the IsolationSegment field.
func (o *ProcessStats) SetIsolationSegment(v string) {
	o.IsolationSegment.Set(&v)
}
// SetIsolationSegmentNil sets the value for IsolationSegment to be an explicit nil
func (o *ProcessStats) SetIsolationSegmentNil() {
	o.IsolationSegment.Set(nil)
}

// UnsetIsolationSegment ensures that no value is present for IsolationSegment, not even an explicit nil
func (o *ProcessStats) UnsetIsolationSegment() {
	o.IsolationSegment.Unset()
}

// GetLogRateLimit returns the LogRateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessStats) GetLogRateLimit() int32 {
	if o == nil || IsNil(o.LogRateLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.LogRateLimit.Get()
}

// GetLogRateLimitOk returns a tuple with the LogRateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessStats) GetLogRateLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogRateLimit.Get(), o.LogRateLimit.IsSet()
}

// HasLogRateLimit returns a boolean if a field has been set.
func (o *ProcessStats) HasLogRateLimit() bool {
	if o != nil && o.LogRateLimit.IsSet() {
		return true
	}

	return false
}

// SetLogRateLimit gets a reference to the given NullableInt32 and assigns it to the LogRateLimit field.
func (o *ProcessStats) SetLogRateLimit(v int32) {
	o.LogRateLimit.Set(&v)
}
// SetLogRateLimitNil sets the value for LogRateLimit to be an explicit nil
func (o *ProcessStats) SetLogRateLimitNil() {
	o.LogRateLimit.Set(nil)
}

// UnsetLogRateLimit ensures that no value is present for LogRateLimit, not even an explicit nil
func (o *ProcessStats) UnsetLogRateLimit() {
	o.LogRateLimit.Unset()
}

// GetMemQuota returns the MemQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessStats) GetMemQuota() int32 {
	if o == nil || IsNil(o.MemQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.MemQuota.Get()
}

// GetMemQuotaOk returns a tuple with the MemQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessStats) GetMemQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemQuota.Get(), o.MemQuota.IsSet()
}

// HasMemQuota returns a boolean if a field has been set.
func (o *ProcessStats) HasMemQuota() bool {
	if o != nil && o.MemQuota.IsSet() {
		return true
	}

	return false
}

// SetMemQuota gets a reference to the given NullableInt32 and assigns it to the MemQuota field.
func (o *ProcessStats) SetMemQuota(v int32) {
	o.MemQuota.Set(&v)
}
// SetMemQuotaNil sets the value for MemQuota to be an explicit nil
func (o *ProcessStats) SetMemQuotaNil() {
	o.MemQuota.Set(nil)
}

// UnsetMemQuota ensures that no value is present for MemQuota, not even an explicit nil
func (o *ProcessStats) UnsetMemQuota() {
	o.MemQuota.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProcessStats) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProcessStats) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProcessStats) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProcessStats) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProcessStats) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProcessStats) SetType(v string) {
	o.Type = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *ProcessStats) GetUptime() int32 {
	if o == nil || IsNil(o.Uptime) {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetUptimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *ProcessStats) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *ProcessStats) SetUptime(v int32) {
	o.Uptime = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ProcessStats) GetUsage() ProcessStatsUsage {
	if o == nil || IsNil(o.Usage) {
		var ret ProcessStatsUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessStats) GetUsageOk() (*ProcessStatsUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ProcessStats) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given ProcessStatsUsage and assigns it to the Usage field.
func (o *ProcessStats) SetUsage(v ProcessStatsUsage) {
	o.Usage = &v
}

func (o ProcessStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.DiskQuota.IsSet() {
		toSerialize["disk_quota"] = o.DiskQuota.Get()
	}
	if !IsNil(o.FdsQuota) {
		toSerialize["fds_quota"] = o.FdsQuota
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.InstanceInternalIp) {
		toSerialize["instance_internal_ip"] = o.InstanceInternalIp
	}
	if !IsNil(o.InstancePorts) {
		toSerialize["instance_ports"] = o.InstancePorts
	}
	if o.IsolationSegment.IsSet() {
		toSerialize["isolation_segment"] = o.IsolationSegment.Get()
	}
	if o.LogRateLimit.IsSet() {
		toSerialize["log_rate_limit"] = o.LogRateLimit.Get()
	}
	if o.MemQuota.IsSet() {
		toSerialize["mem_quota"] = o.MemQuota.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableProcessStats struct {
	value *ProcessStats
	isSet bool
}

func (v NullableProcessStats) Get() *ProcessStats {
	return v.value
}

func (v *NullableProcessStats) Set(val *ProcessStats) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessStats) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessStats(val *ProcessStats) *NullableProcessStats {
	return &NullableProcessStats{value: val, isSet: true}
}

func (v NullableProcessStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

