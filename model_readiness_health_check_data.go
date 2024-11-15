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

// checks if the ReadinessHealthCheckData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadinessHealthCheckData{}

// ReadinessHealthCheckData struct for ReadinessHealthCheckData
type ReadinessHealthCheckData struct {
	Endpoint NullableString `json:"endpoint,omitempty"`
	Interval NullableInt32 `json:"interval,omitempty"`
	InvocationTimeout NullableInt32 `json:"invocation_timeout,omitempty"`
}

// NewReadinessHealthCheckData instantiates a new ReadinessHealthCheckData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadinessHealthCheckData() *ReadinessHealthCheckData {
	this := ReadinessHealthCheckData{}
	return &this
}

// NewReadinessHealthCheckDataWithDefaults instantiates a new ReadinessHealthCheckData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadinessHealthCheckDataWithDefaults() *ReadinessHealthCheckData {
	this := ReadinessHealthCheckData{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReadinessHealthCheckData) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReadinessHealthCheckData) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ReadinessHealthCheckData) HasEndpoint() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableString and assigns it to the Endpoint field.
func (o *ReadinessHealthCheckData) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *ReadinessHealthCheckData) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *ReadinessHealthCheckData) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReadinessHealthCheckData) GetInterval() int32 {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReadinessHealthCheckData) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *ReadinessHealthCheckData) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *ReadinessHealthCheckData) SetInterval(v int32) {
	o.Interval.Set(&v)
}
// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *ReadinessHealthCheckData) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *ReadinessHealthCheckData) UnsetInterval() {
	o.Interval.Unset()
}

// GetInvocationTimeout returns the InvocationTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReadinessHealthCheckData) GetInvocationTimeout() int32 {
	if o == nil || IsNil(o.InvocationTimeout.Get()) {
		var ret int32
		return ret
	}
	return *o.InvocationTimeout.Get()
}

// GetInvocationTimeoutOk returns a tuple with the InvocationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReadinessHealthCheckData) GetInvocationTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvocationTimeout.Get(), o.InvocationTimeout.IsSet()
}

// HasInvocationTimeout returns a boolean if a field has been set.
func (o *ReadinessHealthCheckData) HasInvocationTimeout() bool {
	if o != nil && o.InvocationTimeout.IsSet() {
		return true
	}

	return false
}

// SetInvocationTimeout gets a reference to the given NullableInt32 and assigns it to the InvocationTimeout field.
func (o *ReadinessHealthCheckData) SetInvocationTimeout(v int32) {
	o.InvocationTimeout.Set(&v)
}
// SetInvocationTimeoutNil sets the value for InvocationTimeout to be an explicit nil
func (o *ReadinessHealthCheckData) SetInvocationTimeoutNil() {
	o.InvocationTimeout.Set(nil)
}

// UnsetInvocationTimeout ensures that no value is present for InvocationTimeout, not even an explicit nil
func (o *ReadinessHealthCheckData) UnsetInvocationTimeout() {
	o.InvocationTimeout.Unset()
}

func (o ReadinessHealthCheckData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadinessHealthCheckData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.InvocationTimeout.IsSet() {
		toSerialize["invocation_timeout"] = o.InvocationTimeout.Get()
	}
	return toSerialize, nil
}

type NullableReadinessHealthCheckData struct {
	value *ReadinessHealthCheckData
	isSet bool
}

func (v NullableReadinessHealthCheckData) Get() *ReadinessHealthCheckData {
	return v.value
}

func (v *NullableReadinessHealthCheckData) Set(val *ReadinessHealthCheckData) {
	v.value = val
	v.isSet = true
}

func (v NullableReadinessHealthCheckData) IsSet() bool {
	return v.isSet
}

func (v *NullableReadinessHealthCheckData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadinessHealthCheckData(val *ReadinessHealthCheckData) *NullableReadinessHealthCheckData {
	return &NullableReadinessHealthCheckData{value: val, isSet: true}
}

func (v NullableReadinessHealthCheckData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadinessHealthCheckData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


