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

// checks if the V3UsersGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3UsersGet200Response{}

// V3UsersGet200Response struct for V3UsersGet200Response
type V3UsersGet200Response struct {
	Pagination *V3RoutesGet200ResponsePagination `json:"pagination,omitempty"`
	Resources []User `json:"resources,omitempty"`
}

// NewV3UsersGet200Response instantiates a new V3UsersGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3UsersGet200Response() *V3UsersGet200Response {
	this := V3UsersGet200Response{}
	return &this
}

// NewV3UsersGet200ResponseWithDefaults instantiates a new V3UsersGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3UsersGet200ResponseWithDefaults() *V3UsersGet200Response {
	this := V3UsersGet200Response{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *V3UsersGet200Response) GetPagination() V3RoutesGet200ResponsePagination {
	if o == nil || IsNil(o.Pagination) {
		var ret V3RoutesGet200ResponsePagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3UsersGet200Response) GetPaginationOk() (*V3RoutesGet200ResponsePagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *V3UsersGet200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given V3RoutesGet200ResponsePagination and assigns it to the Pagination field.
func (o *V3UsersGet200Response) SetPagination(v V3RoutesGet200ResponsePagination) {
	o.Pagination = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V3UsersGet200Response) GetResources() []User {
	if o == nil || IsNil(o.Resources) {
		var ret []User
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3UsersGet200Response) GetResourcesOk() ([]User, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V3UsersGet200Response) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []User and assigns it to the Resources field.
func (o *V3UsersGet200Response) SetResources(v []User) {
	o.Resources = v
}

func (o V3UsersGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3UsersGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableV3UsersGet200Response struct {
	value *V3UsersGet200Response
	isSet bool
}

func (v NullableV3UsersGet200Response) Get() *V3UsersGet200Response {
	return v.value
}

func (v *NullableV3UsersGet200Response) Set(val *V3UsersGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3UsersGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3UsersGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3UsersGet200Response(val *V3UsersGet200Response) *NullableV3UsersGet200Response {
	return &NullableV3UsersGet200Response{value: val, isSet: true}
}

func (v NullableV3UsersGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3UsersGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

