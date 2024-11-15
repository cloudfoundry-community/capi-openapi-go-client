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

// checks if the ServicesQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesQuota{}

// ServicesQuota struct for ServicesQuota
type ServicesQuota struct {
	// Specifies whether instances of paid service plans can be created
	PaidServicesAllowed *bool `json:"paid_services_allowed,omitempty"`
	// Total number of service instances allowed in an organization
	TotalServiceInstances NullableInt32 `json:"total_service_instances,omitempty"`
	// Total number of service keys allowed in an organization
	TotalServiceKeys NullableInt32 `json:"total_service_keys,omitempty"`
}

// NewServicesQuota instantiates a new ServicesQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesQuota() *ServicesQuota {
	this := ServicesQuota{}
	return &this
}

// NewServicesQuotaWithDefaults instantiates a new ServicesQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesQuotaWithDefaults() *ServicesQuota {
	this := ServicesQuota{}
	return &this
}

// GetPaidServicesAllowed returns the PaidServicesAllowed field value if set, zero value otherwise.
func (o *ServicesQuota) GetPaidServicesAllowed() bool {
	if o == nil || IsNil(o.PaidServicesAllowed) {
		var ret bool
		return ret
	}
	return *o.PaidServicesAllowed
}

// GetPaidServicesAllowedOk returns a tuple with the PaidServicesAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesQuota) GetPaidServicesAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.PaidServicesAllowed) {
		return nil, false
	}
	return o.PaidServicesAllowed, true
}

// HasPaidServicesAllowed returns a boolean if a field has been set.
func (o *ServicesQuota) HasPaidServicesAllowed() bool {
	if o != nil && !IsNil(o.PaidServicesAllowed) {
		return true
	}

	return false
}

// SetPaidServicesAllowed gets a reference to the given bool and assigns it to the PaidServicesAllowed field.
func (o *ServicesQuota) SetPaidServicesAllowed(v bool) {
	o.PaidServicesAllowed = &v
}

// GetTotalServiceInstances returns the TotalServiceInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicesQuota) GetTotalServiceInstances() int32 {
	if o == nil || IsNil(o.TotalServiceInstances.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalServiceInstances.Get()
}

// GetTotalServiceInstancesOk returns a tuple with the TotalServiceInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicesQuota) GetTotalServiceInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalServiceInstances.Get(), o.TotalServiceInstances.IsSet()
}

// HasTotalServiceInstances returns a boolean if a field has been set.
func (o *ServicesQuota) HasTotalServiceInstances() bool {
	if o != nil && o.TotalServiceInstances.IsSet() {
		return true
	}

	return false
}

// SetTotalServiceInstances gets a reference to the given NullableInt32 and assigns it to the TotalServiceInstances field.
func (o *ServicesQuota) SetTotalServiceInstances(v int32) {
	o.TotalServiceInstances.Set(&v)
}
// SetTotalServiceInstancesNil sets the value for TotalServiceInstances to be an explicit nil
func (o *ServicesQuota) SetTotalServiceInstancesNil() {
	o.TotalServiceInstances.Set(nil)
}

// UnsetTotalServiceInstances ensures that no value is present for TotalServiceInstances, not even an explicit nil
func (o *ServicesQuota) UnsetTotalServiceInstances() {
	o.TotalServiceInstances.Unset()
}

// GetTotalServiceKeys returns the TotalServiceKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicesQuota) GetTotalServiceKeys() int32 {
	if o == nil || IsNil(o.TotalServiceKeys.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalServiceKeys.Get()
}

// GetTotalServiceKeysOk returns a tuple with the TotalServiceKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicesQuota) GetTotalServiceKeysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalServiceKeys.Get(), o.TotalServiceKeys.IsSet()
}

// HasTotalServiceKeys returns a boolean if a field has been set.
func (o *ServicesQuota) HasTotalServiceKeys() bool {
	if o != nil && o.TotalServiceKeys.IsSet() {
		return true
	}

	return false
}

// SetTotalServiceKeys gets a reference to the given NullableInt32 and assigns it to the TotalServiceKeys field.
func (o *ServicesQuota) SetTotalServiceKeys(v int32) {
	o.TotalServiceKeys.Set(&v)
}
// SetTotalServiceKeysNil sets the value for TotalServiceKeys to be an explicit nil
func (o *ServicesQuota) SetTotalServiceKeysNil() {
	o.TotalServiceKeys.Set(nil)
}

// UnsetTotalServiceKeys ensures that no value is present for TotalServiceKeys, not even an explicit nil
func (o *ServicesQuota) UnsetTotalServiceKeys() {
	o.TotalServiceKeys.Unset()
}

func (o ServicesQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaidServicesAllowed) {
		toSerialize["paid_services_allowed"] = o.PaidServicesAllowed
	}
	if o.TotalServiceInstances.IsSet() {
		toSerialize["total_service_instances"] = o.TotalServiceInstances.Get()
	}
	if o.TotalServiceKeys.IsSet() {
		toSerialize["total_service_keys"] = o.TotalServiceKeys.Get()
	}
	return toSerialize, nil
}

type NullableServicesQuota struct {
	value *ServicesQuota
	isSet bool
}

func (v NullableServicesQuota) Get() *ServicesQuota {
	return v.value
}

func (v *NullableServicesQuota) Set(val *ServicesQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesQuota(val *ServicesQuota) *NullableServicesQuota {
	return &NullableServicesQuota{value: val, isSet: true}
}

func (v NullableServicesQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

