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

// checks if the Authentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authentication{}

// Authentication struct for Authentication
type Authentication struct {
	Credentials *AuthenticationCredentials `json:"credentials,omitempty"`
	// Type of the authentication mechanisms. Valid value is basic.
	Type *string `json:"type,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Authentication) GetCredentials() AuthenticationCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret AuthenticationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetCredentialsOk() (*AuthenticationCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Authentication) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AuthenticationCredentials and assigns it to the Credentials field.
func (o *Authentication) SetCredentials(v AuthenticationCredentials) {
	o.Credentials = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Authentication) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Authentication) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Authentication) SetType(v string) {
	o.Type = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

