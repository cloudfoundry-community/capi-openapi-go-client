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

// checks if the RoleLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleLinks{}

// RoleLinks struct for RoleLinks
type RoleLinks struct {
	Organization *Get200ResponseLinksLogCache `json:"organization,omitempty"`
	Self *Get200ResponseLinksLogCache `json:"self,omitempty"`
	User *Get200ResponseLinksLogCache `json:"user,omitempty"`
}

// NewRoleLinks instantiates a new RoleLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleLinks() *RoleLinks {
	this := RoleLinks{}
	return &this
}

// NewRoleLinksWithDefaults instantiates a new RoleLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleLinksWithDefaults() *RoleLinks {
	this := RoleLinks{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RoleLinks) GetOrganization() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Organization) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleLinks) GetOrganizationOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RoleLinks) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Organization field.
func (o *RoleLinks) SetOrganization(v Get200ResponseLinksLogCache) {
	o.Organization = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *RoleLinks) GetSelf() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Self) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleLinks) GetSelfOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *RoleLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Self field.
func (o *RoleLinks) SetSelf(v Get200ResponseLinksLogCache) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RoleLinks) GetUser() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.User) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleLinks) GetUserOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RoleLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given Get200ResponseLinksLogCache and assigns it to the User field.
func (o *RoleLinks) SetUser(v Get200ResponseLinksLogCache) {
	o.User = &v
}

func (o RoleLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableRoleLinks struct {
	value *RoleLinks
	isSet bool
}

func (v NullableRoleLinks) Get() *RoleLinks {
	return v.value
}

func (v *NullableRoleLinks) Set(val *RoleLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleLinks(val *RoleLinks) *NullableRoleLinks {
	return &NullableRoleLinks{value: val, isSet: true}
}

func (v NullableRoleLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


