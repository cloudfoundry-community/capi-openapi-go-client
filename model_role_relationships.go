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

// checks if the RoleRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleRelationships{}

// RoleRelationships struct for RoleRelationships
type RoleRelationships struct {
	Organization *V3AppsPostRequestRelationshipsSpace `json:"organization,omitempty"`
	Space *RoleRelationshipsSpace `json:"space,omitempty"`
	User *V3RolesPostRequestRelationshipsUser `json:"user,omitempty"`
}

// NewRoleRelationships instantiates a new RoleRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRelationships() *RoleRelationships {
	this := RoleRelationships{}
	return &this
}

// NewRoleRelationshipsWithDefaults instantiates a new RoleRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleRelationshipsWithDefaults() *RoleRelationships {
	this := RoleRelationships{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RoleRelationships) GetOrganization() V3AppsPostRequestRelationshipsSpace {
	if o == nil || IsNil(o.Organization) {
		var ret V3AppsPostRequestRelationshipsSpace
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRelationships) GetOrganizationOk() (*V3AppsPostRequestRelationshipsSpace, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RoleRelationships) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given V3AppsPostRequestRelationshipsSpace and assigns it to the Organization field.
func (o *RoleRelationships) SetOrganization(v V3AppsPostRequestRelationshipsSpace) {
	o.Organization = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *RoleRelationships) GetSpace() RoleRelationshipsSpace {
	if o == nil || IsNil(o.Space) {
		var ret RoleRelationshipsSpace
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRelationships) GetSpaceOk() (*RoleRelationshipsSpace, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *RoleRelationships) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given RoleRelationshipsSpace and assigns it to the Space field.
func (o *RoleRelationships) SetSpace(v RoleRelationshipsSpace) {
	o.Space = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RoleRelationships) GetUser() V3RolesPostRequestRelationshipsUser {
	if o == nil || IsNil(o.User) {
		var ret V3RolesPostRequestRelationshipsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRelationships) GetUserOk() (*V3RolesPostRequestRelationshipsUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RoleRelationships) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given V3RolesPostRequestRelationshipsUser and assigns it to the User field.
func (o *RoleRelationships) SetUser(v V3RolesPostRequestRelationshipsUser) {
	o.User = &v
}

func (o RoleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableRoleRelationships struct {
	value *RoleRelationships
	isSet bool
}

func (v NullableRoleRelationships) Get() *RoleRelationships {
	return v.value
}

func (v *NullableRoleRelationships) Set(val *RoleRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRelationships(val *RoleRelationships) *NullableRoleRelationships {
	return &NullableRoleRelationships{value: val, isSet: true}
}

func (v NullableRoleRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


