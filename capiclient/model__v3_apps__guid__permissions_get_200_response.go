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

// checks if the V3AppsGuidPermissionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3AppsGuidPermissionsGet200Response{}

// V3AppsGuidPermissionsGet200Response struct for V3AppsGuidPermissionsGet200Response
type V3AppsGuidPermissionsGet200Response struct {
	ReadBasicData *bool `json:"read_basic_data,omitempty"`
	ReadSensitiveData *bool `json:"read_sensitive_data,omitempty"`
}

// NewV3AppsGuidPermissionsGet200Response instantiates a new V3AppsGuidPermissionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AppsGuidPermissionsGet200Response() *V3AppsGuidPermissionsGet200Response {
	this := V3AppsGuidPermissionsGet200Response{}
	return &this
}

// NewV3AppsGuidPermissionsGet200ResponseWithDefaults instantiates a new V3AppsGuidPermissionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AppsGuidPermissionsGet200ResponseWithDefaults() *V3AppsGuidPermissionsGet200Response {
	this := V3AppsGuidPermissionsGet200Response{}
	return &this
}

// GetReadBasicData returns the ReadBasicData field value if set, zero value otherwise.
func (o *V3AppsGuidPermissionsGet200Response) GetReadBasicData() bool {
	if o == nil || IsNil(o.ReadBasicData) {
		var ret bool
		return ret
	}
	return *o.ReadBasicData
}

// GetReadBasicDataOk returns a tuple with the ReadBasicData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidPermissionsGet200Response) GetReadBasicDataOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadBasicData) {
		return nil, false
	}
	return o.ReadBasicData, true
}

// HasReadBasicData returns a boolean if a field has been set.
func (o *V3AppsGuidPermissionsGet200Response) HasReadBasicData() bool {
	if o != nil && !IsNil(o.ReadBasicData) {
		return true
	}

	return false
}

// SetReadBasicData gets a reference to the given bool and assigns it to the ReadBasicData field.
func (o *V3AppsGuidPermissionsGet200Response) SetReadBasicData(v bool) {
	o.ReadBasicData = &v
}

// GetReadSensitiveData returns the ReadSensitiveData field value if set, zero value otherwise.
func (o *V3AppsGuidPermissionsGet200Response) GetReadSensitiveData() bool {
	if o == nil || IsNil(o.ReadSensitiveData) {
		var ret bool
		return ret
	}
	return *o.ReadSensitiveData
}

// GetReadSensitiveDataOk returns a tuple with the ReadSensitiveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidPermissionsGet200Response) GetReadSensitiveDataOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadSensitiveData) {
		return nil, false
	}
	return o.ReadSensitiveData, true
}

// HasReadSensitiveData returns a boolean if a field has been set.
func (o *V3AppsGuidPermissionsGet200Response) HasReadSensitiveData() bool {
	if o != nil && !IsNil(o.ReadSensitiveData) {
		return true
	}

	return false
}

// SetReadSensitiveData gets a reference to the given bool and assigns it to the ReadSensitiveData field.
func (o *V3AppsGuidPermissionsGet200Response) SetReadSensitiveData(v bool) {
	o.ReadSensitiveData = &v
}

func (o V3AppsGuidPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3AppsGuidPermissionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReadBasicData) {
		toSerialize["read_basic_data"] = o.ReadBasicData
	}
	if !IsNil(o.ReadSensitiveData) {
		toSerialize["read_sensitive_data"] = o.ReadSensitiveData
	}
	return toSerialize, nil
}

type NullableV3AppsGuidPermissionsGet200Response struct {
	value *V3AppsGuidPermissionsGet200Response
	isSet bool
}

func (v NullableV3AppsGuidPermissionsGet200Response) Get() *V3AppsGuidPermissionsGet200Response {
	return v.value
}

func (v *NullableV3AppsGuidPermissionsGet200Response) Set(val *V3AppsGuidPermissionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AppsGuidPermissionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AppsGuidPermissionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AppsGuidPermissionsGet200Response(val *V3AppsGuidPermissionsGet200Response) *NullableV3AppsGuidPermissionsGet200Response {
	return &NullableV3AppsGuidPermissionsGet200Response{value: val, isSet: true}
}

func (v NullableV3AppsGuidPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AppsGuidPermissionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


