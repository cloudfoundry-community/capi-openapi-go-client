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

// checks if the RevisionsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionsList{}

// RevisionsList struct for RevisionsList
type RevisionsList struct {
	Pagination *RevisionsListPagination `json:"pagination,omitempty"`
	Resources []Revision `json:"resources,omitempty"`
}

// NewRevisionsList instantiates a new RevisionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionsList() *RevisionsList {
	this := RevisionsList{}
	return &this
}

// NewRevisionsListWithDefaults instantiates a new RevisionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionsListWithDefaults() *RevisionsList {
	this := RevisionsList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *RevisionsList) GetPagination() RevisionsListPagination {
	if o == nil || IsNil(o.Pagination) {
		var ret RevisionsListPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsList) GetPaginationOk() (*RevisionsListPagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *RevisionsList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given RevisionsListPagination and assigns it to the Pagination field.
func (o *RevisionsList) SetPagination(v RevisionsListPagination) {
	o.Pagination = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RevisionsList) GetResources() []Revision {
	if o == nil || IsNil(o.Resources) {
		var ret []Revision
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsList) GetResourcesOk() ([]Revision, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RevisionsList) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Revision and assigns it to the Resources field.
func (o *RevisionsList) SetResources(v []Revision) {
	o.Resources = v
}

func (o RevisionsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableRevisionsList struct {
	value *RevisionsList
	isSet bool
}

func (v NullableRevisionsList) Get() *RevisionsList {
	return v.value
}

func (v *NullableRevisionsList) Set(val *RevisionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionsList(val *RevisionsList) *NullableRevisionsList {
	return &NullableRevisionsList{value: val, isSet: true}
}

func (v NullableRevisionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


