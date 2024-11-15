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

// checks if the V3StacksGuidAppsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3StacksGuidAppsGet200Response{}

// V3StacksGuidAppsGet200Response struct for V3StacksGuidAppsGet200Response
type V3StacksGuidAppsGet200Response struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Resources []App `json:"resources,omitempty"`
}

// NewV3StacksGuidAppsGet200Response instantiates a new V3StacksGuidAppsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3StacksGuidAppsGet200Response() *V3StacksGuidAppsGet200Response {
	this := V3StacksGuidAppsGet200Response{}
	return &this
}

// NewV3StacksGuidAppsGet200ResponseWithDefaults instantiates a new V3StacksGuidAppsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3StacksGuidAppsGet200ResponseWithDefaults() *V3StacksGuidAppsGet200Response {
	this := V3StacksGuidAppsGet200Response{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *V3StacksGuidAppsGet200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3StacksGuidAppsGet200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *V3StacksGuidAppsGet200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *V3StacksGuidAppsGet200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V3StacksGuidAppsGet200Response) GetResources() []App {
	if o == nil || IsNil(o.Resources) {
		var ret []App
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3StacksGuidAppsGet200Response) GetResourcesOk() ([]App, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V3StacksGuidAppsGet200Response) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []App and assigns it to the Resources field.
func (o *V3StacksGuidAppsGet200Response) SetResources(v []App) {
	o.Resources = v
}

func (o V3StacksGuidAppsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3StacksGuidAppsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableV3StacksGuidAppsGet200Response struct {
	value *V3StacksGuidAppsGet200Response
	isSet bool
}

func (v NullableV3StacksGuidAppsGet200Response) Get() *V3StacksGuidAppsGet200Response {
	return v.value
}

func (v *NullableV3StacksGuidAppsGet200Response) Set(val *V3StacksGuidAppsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3StacksGuidAppsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3StacksGuidAppsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3StacksGuidAppsGet200Response(val *V3StacksGuidAppsGet200Response) *NullableV3StacksGuidAppsGet200Response {
	return &NullableV3StacksGuidAppsGet200Response{value: val, isSet: true}
}

func (v NullableV3StacksGuidAppsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3StacksGuidAppsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


