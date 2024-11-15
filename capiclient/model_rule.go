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

// checks if the Rule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rule{}

// Rule struct for Rule
type Rule struct {
	Code NullableInt32 `json:"code,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Log NullableBool `json:"log,omitempty"`
	Ports NullableString `json:"ports,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Type NullableInt32 `json:"type,omitempty"`
}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule() *Rule {
	this := Rule{}
	return &this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := Rule{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rule) GetCode() int32 {
	if o == nil || IsNil(o.Code.Get()) {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rule) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *Rule) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *Rule) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *Rule) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *Rule) UnsetCode() {
	o.Code.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rule) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Rule) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Rule) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Rule) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Rule) UnsetDescription() {
	o.Description.Unset()
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Rule) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Rule) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *Rule) SetDestination(v string) {
	o.Destination = &v
}

// GetLog returns the Log field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rule) GetLog() bool {
	if o == nil || IsNil(o.Log.Get()) {
		var ret bool
		return ret
	}
	return *o.Log.Get()
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rule) GetLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Log.Get(), o.Log.IsSet()
}

// HasLog returns a boolean if a field has been set.
func (o *Rule) HasLog() bool {
	if o != nil && o.Log.IsSet() {
		return true
	}

	return false
}

// SetLog gets a reference to the given NullableBool and assigns it to the Log field.
func (o *Rule) SetLog(v bool) {
	o.Log.Set(&v)
}
// SetLogNil sets the value for Log to be an explicit nil
func (o *Rule) SetLogNil() {
	o.Log.Set(nil)
}

// UnsetLog ensures that no value is present for Log, not even an explicit nil
func (o *Rule) UnsetLog() {
	o.Log.Unset()
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rule) GetPorts() string {
	if o == nil || IsNil(o.Ports.Get()) {
		var ret string
		return ret
	}
	return *o.Ports.Get()
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rule) GetPortsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports.Get(), o.Ports.IsSet()
}

// HasPorts returns a boolean if a field has been set.
func (o *Rule) HasPorts() bool {
	if o != nil && o.Ports.IsSet() {
		return true
	}

	return false
}

// SetPorts gets a reference to the given NullableString and assigns it to the Ports field.
func (o *Rule) SetPorts(v string) {
	o.Ports.Set(&v)
}
// SetPortsNil sets the value for Ports to be an explicit nil
func (o *Rule) SetPortsNil() {
	o.Ports.Set(nil)
}

// UnsetPorts ensures that no value is present for Ports, not even an explicit nil
func (o *Rule) UnsetPorts() {
	o.Ports.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Rule) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Rule) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Rule) SetProtocol(v string) {
	o.Protocol = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rule) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rule) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Rule) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *Rule) SetType(v int32) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Rule) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Rule) UnsetType() {
	o.Type.Unset()
}

func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if o.Log.IsSet() {
		toSerialize["log"] = o.Log.Get()
	}
	if o.Ports.IsSet() {
		toSerialize["ports"] = o.Ports.Get()
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


