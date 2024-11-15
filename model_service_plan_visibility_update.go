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

// checks if the ServicePlanVisibilityUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePlanVisibilityUpdate{}

// ServicePlanVisibilityUpdate struct for ServicePlanVisibilityUpdate
type ServicePlanVisibilityUpdate struct {
	Organizations []V3AppsPostRequestRelationshipsSpaceData `json:"organizations,omitempty"`
	// Denotes the visibility of the plan; can be public, admin, organization
	Type *string `json:"type,omitempty"`
}

// NewServicePlanVisibilityUpdate instantiates a new ServicePlanVisibilityUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlanVisibilityUpdate() *ServicePlanVisibilityUpdate {
	this := ServicePlanVisibilityUpdate{}
	return &this
}

// NewServicePlanVisibilityUpdateWithDefaults instantiates a new ServicePlanVisibilityUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanVisibilityUpdateWithDefaults() *ServicePlanVisibilityUpdate {
	this := ServicePlanVisibilityUpdate{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ServicePlanVisibilityUpdate) GetOrganizations() []V3AppsPostRequestRelationshipsSpaceData {
	if o == nil || IsNil(o.Organizations) {
		var ret []V3AppsPostRequestRelationshipsSpaceData
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanVisibilityUpdate) GetOrganizationsOk() ([]V3AppsPostRequestRelationshipsSpaceData, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ServicePlanVisibilityUpdate) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []V3AppsPostRequestRelationshipsSpaceData and assigns it to the Organizations field.
func (o *ServicePlanVisibilityUpdate) SetOrganizations(v []V3AppsPostRequestRelationshipsSpaceData) {
	o.Organizations = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServicePlanVisibilityUpdate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanVisibilityUpdate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServicePlanVisibilityUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServicePlanVisibilityUpdate) SetType(v string) {
	o.Type = &v
}

func (o ServicePlanVisibilityUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePlanVisibilityUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableServicePlanVisibilityUpdate struct {
	value *ServicePlanVisibilityUpdate
	isSet bool
}

func (v NullableServicePlanVisibilityUpdate) Get() *ServicePlanVisibilityUpdate {
	return v.value
}

func (v *NullableServicePlanVisibilityUpdate) Set(val *ServicePlanVisibilityUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlanVisibilityUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlanVisibilityUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlanVisibilityUpdate(val *ServicePlanVisibilityUpdate) *NullableServicePlanVisibilityUpdate {
	return &NullableServicePlanVisibilityUpdate{value: val, isSet: true}
}

func (v NullableServicePlanVisibilityUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlanVisibilityUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


