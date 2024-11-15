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

// checks if the SecurityGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupUpdate{}

// SecurityGroupUpdate struct for SecurityGroupUpdate
type SecurityGroupUpdate struct {
	GloballyEnabled *SecurityGroupUpdateGloballyEnabled `json:"globally_enabled,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Rules []Rule `json:"rules,omitempty"`
}

// NewSecurityGroupUpdate instantiates a new SecurityGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupUpdate() *SecurityGroupUpdate {
	this := SecurityGroupUpdate{}
	return &this
}

// NewSecurityGroupUpdateWithDefaults instantiates a new SecurityGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupUpdateWithDefaults() *SecurityGroupUpdate {
	this := SecurityGroupUpdate{}
	return &this
}

// GetGloballyEnabled returns the GloballyEnabled field value if set, zero value otherwise.
func (o *SecurityGroupUpdate) GetGloballyEnabled() SecurityGroupUpdateGloballyEnabled {
	if o == nil || IsNil(o.GloballyEnabled) {
		var ret SecurityGroupUpdateGloballyEnabled
		return ret
	}
	return *o.GloballyEnabled
}

// GetGloballyEnabledOk returns a tuple with the GloballyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupUpdate) GetGloballyEnabledOk() (*SecurityGroupUpdateGloballyEnabled, bool) {
	if o == nil || IsNil(o.GloballyEnabled) {
		return nil, false
	}
	return o.GloballyEnabled, true
}

// HasGloballyEnabled returns a boolean if a field has been set.
func (o *SecurityGroupUpdate) HasGloballyEnabled() bool {
	if o != nil && !IsNil(o.GloballyEnabled) {
		return true
	}

	return false
}

// SetGloballyEnabled gets a reference to the given SecurityGroupUpdateGloballyEnabled and assigns it to the GloballyEnabled field.
func (o *SecurityGroupUpdate) SetGloballyEnabled(v SecurityGroupUpdateGloballyEnabled) {
	o.GloballyEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityGroupUpdate) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityGroupUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SecurityGroupUpdate) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SecurityGroupUpdate) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SecurityGroupUpdate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SecurityGroupUpdate) UnsetName() {
	o.Name.Unset()
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityGroupUpdate) GetRules() []Rule {
	if o == nil {
		var ret []Rule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityGroupUpdate) GetRulesOk() ([]Rule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SecurityGroupUpdate) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []Rule and assigns it to the Rules field.
func (o *SecurityGroupUpdate) SetRules(v []Rule) {
	o.Rules = v
}

func (o SecurityGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GloballyEnabled) {
		toSerialize["globally_enabled"] = o.GloballyEnabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableSecurityGroupUpdate struct {
	value *SecurityGroupUpdate
	isSet bool
}

func (v NullableSecurityGroupUpdate) Get() *SecurityGroupUpdate {
	return v.value
}

func (v *NullableSecurityGroupUpdate) Set(val *SecurityGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupUpdate(val *SecurityGroupUpdate) *NullableSecurityGroupUpdate {
	return &NullableSecurityGroupUpdate{value: val, isSet: true}
}

func (v NullableSecurityGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


