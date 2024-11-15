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

// checks if the V3TasksGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3TasksGet200Response{}

// V3TasksGet200Response struct for V3TasksGet200Response
type V3TasksGet200Response struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Resources []Task `json:"resources,omitempty"`
}

// NewV3TasksGet200Response instantiates a new V3TasksGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3TasksGet200Response() *V3TasksGet200Response {
	this := V3TasksGet200Response{}
	return &this
}

// NewV3TasksGet200ResponseWithDefaults instantiates a new V3TasksGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3TasksGet200ResponseWithDefaults() *V3TasksGet200Response {
	this := V3TasksGet200Response{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *V3TasksGet200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3TasksGet200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *V3TasksGet200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *V3TasksGet200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V3TasksGet200Response) GetResources() []Task {
	if o == nil || IsNil(o.Resources) {
		var ret []Task
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3TasksGet200Response) GetResourcesOk() ([]Task, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V3TasksGet200Response) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Task and assigns it to the Resources field.
func (o *V3TasksGet200Response) SetResources(v []Task) {
	o.Resources = v
}

func (o V3TasksGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3TasksGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableV3TasksGet200Response struct {
	value *V3TasksGet200Response
	isSet bool
}

func (v NullableV3TasksGet200Response) Get() *V3TasksGet200Response {
	return v.value
}

func (v *NullableV3TasksGet200Response) Set(val *V3TasksGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3TasksGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3TasksGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3TasksGet200Response(val *V3TasksGet200Response) *NullableV3TasksGet200Response {
	return &NullableV3TasksGet200Response{value: val, isSet: true}
}

func (v NullableV3TasksGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3TasksGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


