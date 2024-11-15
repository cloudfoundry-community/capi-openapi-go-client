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

// checks if the V3PackagesGuidDelete202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3PackagesGuidDelete202Response{}

// V3PackagesGuidDelete202Response struct for V3PackagesGuidDelete202Response
type V3PackagesGuidDelete202Response struct {
	Guid *string `json:"guid,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewV3PackagesGuidDelete202Response instantiates a new V3PackagesGuidDelete202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3PackagesGuidDelete202Response() *V3PackagesGuidDelete202Response {
	this := V3PackagesGuidDelete202Response{}
	return &this
}

// NewV3PackagesGuidDelete202ResponseWithDefaults instantiates a new V3PackagesGuidDelete202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3PackagesGuidDelete202ResponseWithDefaults() *V3PackagesGuidDelete202Response {
	this := V3PackagesGuidDelete202Response{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *V3PackagesGuidDelete202Response) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3PackagesGuidDelete202Response) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *V3PackagesGuidDelete202Response) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *V3PackagesGuidDelete202Response) SetGuid(v string) {
	o.Guid = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V3PackagesGuidDelete202Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3PackagesGuidDelete202Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V3PackagesGuidDelete202Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V3PackagesGuidDelete202Response) SetUrl(v string) {
	o.Url = &v
}

func (o V3PackagesGuidDelete202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3PackagesGuidDelete202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableV3PackagesGuidDelete202Response struct {
	value *V3PackagesGuidDelete202Response
	isSet bool
}

func (v NullableV3PackagesGuidDelete202Response) Get() *V3PackagesGuidDelete202Response {
	return v.value
}

func (v *NullableV3PackagesGuidDelete202Response) Set(val *V3PackagesGuidDelete202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3PackagesGuidDelete202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3PackagesGuidDelete202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3PackagesGuidDelete202Response(val *V3PackagesGuidDelete202Response) *NullableV3PackagesGuidDelete202Response {
	return &NullableV3PackagesGuidDelete202Response{value: val, isSet: true}
}

func (v NullableV3PackagesGuidDelete202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3PackagesGuidDelete202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

