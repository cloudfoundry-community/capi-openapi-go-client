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

// checks if the MaintenanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceInfo{}

// MaintenanceInfo struct for MaintenanceInfo
type MaintenanceInfo struct {
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewMaintenanceInfo instantiates a new MaintenanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceInfo() *MaintenanceInfo {
	this := MaintenanceInfo{}
	return &this
}

// NewMaintenanceInfoWithDefaults instantiates a new MaintenanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceInfoWithDefaults() *MaintenanceInfo {
	this := MaintenanceInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MaintenanceInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MaintenanceInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MaintenanceInfo) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MaintenanceInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MaintenanceInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MaintenanceInfo) SetVersion(v string) {
	o.Version = &v
}

func (o MaintenanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableMaintenanceInfo struct {
	value *MaintenanceInfo
	isSet bool
}

func (v NullableMaintenanceInfo) Get() *MaintenanceInfo {
	return v.value
}

func (v *NullableMaintenanceInfo) Set(val *MaintenanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceInfo(val *MaintenanceInfo) *NullableMaintenanceInfo {
	return &NullableMaintenanceInfo{value: val, isSet: true}
}

func (v NullableMaintenanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

