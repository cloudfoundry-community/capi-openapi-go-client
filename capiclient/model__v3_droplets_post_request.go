/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V3DropletsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3DropletsPostRequest{}

// V3DropletsPostRequest struct for V3DropletsPostRequest
type V3DropletsPostRequest struct {
	Relationships V3DropletsPostRequestRelationships `json:"relationships"`
}

type _V3DropletsPostRequest V3DropletsPostRequest

// NewV3DropletsPostRequest instantiates a new V3DropletsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3DropletsPostRequest(relationships V3DropletsPostRequestRelationships) *V3DropletsPostRequest {
	this := V3DropletsPostRequest{}
	this.Relationships = relationships
	return &this
}

// NewV3DropletsPostRequestWithDefaults instantiates a new V3DropletsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3DropletsPostRequestWithDefaults() *V3DropletsPostRequest {
	this := V3DropletsPostRequest{}
	return &this
}

// GetRelationships returns the Relationships field value
func (o *V3DropletsPostRequest) GetRelationships() V3DropletsPostRequestRelationships {
	if o == nil {
		var ret V3DropletsPostRequestRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *V3DropletsPostRequest) GetRelationshipsOk() (*V3DropletsPostRequestRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *V3DropletsPostRequest) SetRelationships(v V3DropletsPostRequestRelationships) {
	o.Relationships = v
}

func (o V3DropletsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3DropletsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *V3DropletsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"relationships",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV3DropletsPostRequest := _V3DropletsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV3DropletsPostRequest)

	if err != nil {
		return err
	}

	*o = V3DropletsPostRequest(varV3DropletsPostRequest)

	return err
}

type NullableV3DropletsPostRequest struct {
	value *V3DropletsPostRequest
	isSet bool
}

func (v NullableV3DropletsPostRequest) Get() *V3DropletsPostRequest {
	return v.value
}

func (v *NullableV3DropletsPostRequest) Set(val *V3DropletsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3DropletsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3DropletsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3DropletsPostRequest(val *V3DropletsPostRequest) *NullableV3DropletsPostRequest {
	return &NullableV3DropletsPostRequest{value: val, isSet: true}
}

func (v NullableV3DropletsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3DropletsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


