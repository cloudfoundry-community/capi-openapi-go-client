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

// checks if the DestinationApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationApp{}

// DestinationApp struct for DestinationApp
type DestinationApp struct {
	Guid *string `json:"guid,omitempty"`
	Process *DestinationAppProcess `json:"process,omitempty"`
}

// NewDestinationApp instantiates a new DestinationApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationApp() *DestinationApp {
	this := DestinationApp{}
	return &this
}

// NewDestinationAppWithDefaults instantiates a new DestinationApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationAppWithDefaults() *DestinationApp {
	this := DestinationApp{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *DestinationApp) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApp) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *DestinationApp) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *DestinationApp) SetGuid(v string) {
	o.Guid = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *DestinationApp) GetProcess() DestinationAppProcess {
	if o == nil || IsNil(o.Process) {
		var ret DestinationAppProcess
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApp) GetProcessOk() (*DestinationAppProcess, bool) {
	if o == nil || IsNil(o.Process) {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *DestinationApp) HasProcess() bool {
	if o != nil && !IsNil(o.Process) {
		return true
	}

	return false
}

// SetProcess gets a reference to the given DestinationAppProcess and assigns it to the Process field.
func (o *DestinationApp) SetProcess(v DestinationAppProcess) {
	o.Process = &v
}

func (o DestinationApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Process) {
		toSerialize["process"] = o.Process
	}
	return toSerialize, nil
}

type NullableDestinationApp struct {
	value *DestinationApp
	isSet bool
}

func (v NullableDestinationApp) Get() *DestinationApp {
	return v.value
}

func (v *NullableDestinationApp) Set(val *DestinationApp) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationApp) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationApp(val *DestinationApp) *NullableDestinationApp {
	return &NullableDestinationApp{value: val, isSet: true}
}

func (v NullableDestinationApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


