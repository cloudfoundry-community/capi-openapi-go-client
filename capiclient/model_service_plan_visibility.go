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

// checks if the ServicePlanVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePlanVisibility{}

// ServicePlanVisibility struct for ServicePlanVisibility
type ServicePlanVisibility struct {
	Organizations []OrganizationVisibility `json:"organizations,omitempty"`
	Space *ServicePlanVisibilitySpace `json:"space,omitempty"`
	// Denotes the visibility of the plan; can be public, admin, organization, space
	Type *string `json:"type,omitempty"`
}

// NewServicePlanVisibility instantiates a new ServicePlanVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlanVisibility() *ServicePlanVisibility {
	this := ServicePlanVisibility{}
	return &this
}

// NewServicePlanVisibilityWithDefaults instantiates a new ServicePlanVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanVisibilityWithDefaults() *ServicePlanVisibility {
	this := ServicePlanVisibility{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ServicePlanVisibility) GetOrganizations() []OrganizationVisibility {
	if o == nil || IsNil(o.Organizations) {
		var ret []OrganizationVisibility
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanVisibility) GetOrganizationsOk() ([]OrganizationVisibility, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ServicePlanVisibility) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationVisibility and assigns it to the Organizations field.
func (o *ServicePlanVisibility) SetOrganizations(v []OrganizationVisibility) {
	o.Organizations = v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *ServicePlanVisibility) GetSpace() ServicePlanVisibilitySpace {
	if o == nil || IsNil(o.Space) {
		var ret ServicePlanVisibilitySpace
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanVisibility) GetSpaceOk() (*ServicePlanVisibilitySpace, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *ServicePlanVisibility) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given ServicePlanVisibilitySpace and assigns it to the Space field.
func (o *ServicePlanVisibility) SetSpace(v ServicePlanVisibilitySpace) {
	o.Space = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServicePlanVisibility) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanVisibility) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServicePlanVisibility) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServicePlanVisibility) SetType(v string) {
	o.Type = &v
}

func (o ServicePlanVisibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePlanVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableServicePlanVisibility struct {
	value *ServicePlanVisibility
	isSet bool
}

func (v NullableServicePlanVisibility) Get() *ServicePlanVisibility {
	return v.value
}

func (v *NullableServicePlanVisibility) Set(val *ServicePlanVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlanVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlanVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlanVisibility(val *ServicePlanVisibility) *NullableServicePlanVisibility {
	return &NullableServicePlanVisibility{value: val, isSet: true}
}

func (v NullableServicePlanVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlanVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

