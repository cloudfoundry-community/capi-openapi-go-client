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

// checks if the DomainRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRelationships{}

// DomainRelationships struct for DomainRelationships
type DomainRelationships struct {
	Organization *V3DropletsPostRequestRelationshipsApp `json:"organization,omitempty"`
	SharedOrganizations *DomainRelationshipsSharedOrganizations `json:"shared_organizations,omitempty"`
}

// NewDomainRelationships instantiates a new DomainRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRelationships() *DomainRelationships {
	this := DomainRelationships{}
	return &this
}

// NewDomainRelationshipsWithDefaults instantiates a new DomainRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRelationshipsWithDefaults() *DomainRelationships {
	this := DomainRelationships{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DomainRelationships) GetOrganization() V3DropletsPostRequestRelationshipsApp {
	if o == nil || IsNil(o.Organization) {
		var ret V3DropletsPostRequestRelationshipsApp
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRelationships) GetOrganizationOk() (*V3DropletsPostRequestRelationshipsApp, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DomainRelationships) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given V3DropletsPostRequestRelationshipsApp and assigns it to the Organization field.
func (o *DomainRelationships) SetOrganization(v V3DropletsPostRequestRelationshipsApp) {
	o.Organization = &v
}

// GetSharedOrganizations returns the SharedOrganizations field value if set, zero value otherwise.
func (o *DomainRelationships) GetSharedOrganizations() DomainRelationshipsSharedOrganizations {
	if o == nil || IsNil(o.SharedOrganizations) {
		var ret DomainRelationshipsSharedOrganizations
		return ret
	}
	return *o.SharedOrganizations
}

// GetSharedOrganizationsOk returns a tuple with the SharedOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRelationships) GetSharedOrganizationsOk() (*DomainRelationshipsSharedOrganizations, bool) {
	if o == nil || IsNil(o.SharedOrganizations) {
		return nil, false
	}
	return o.SharedOrganizations, true
}

// HasSharedOrganizations returns a boolean if a field has been set.
func (o *DomainRelationships) HasSharedOrganizations() bool {
	if o != nil && !IsNil(o.SharedOrganizations) {
		return true
	}

	return false
}

// SetSharedOrganizations gets a reference to the given DomainRelationshipsSharedOrganizations and assigns it to the SharedOrganizations field.
func (o *DomainRelationships) SetSharedOrganizations(v DomainRelationshipsSharedOrganizations) {
	o.SharedOrganizations = &v
}

func (o DomainRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.SharedOrganizations) {
		toSerialize["shared_organizations"] = o.SharedOrganizations
	}
	return toSerialize, nil
}

type NullableDomainRelationships struct {
	value *DomainRelationships
	isSet bool
}

func (v NullableDomainRelationships) Get() *DomainRelationships {
	return v.value
}

func (v *NullableDomainRelationships) Set(val *DomainRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRelationships(val *DomainRelationships) *NullableDomainRelationships {
	return &NullableDomainRelationships{value: val, isSet: true}
}

func (v NullableDomainRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

