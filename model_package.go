/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Package type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Package{}

// Package struct for Package
type Package struct {
	// The time with zone when the object was created
	CreatedAt time.Time `json:"created_at"`
	Data *PackageData `json:"data,omitempty"`
	// Unique identifier for the package
	Guid string `json:"guid"`
	Links *PackageLinks `json:"links,omitempty"`
	Metadata *V3PackagesPostRequestMetadata `json:"metadata,omitempty"`
	Relationships *V3PackagesPostRequestRelationships `json:"relationships,omitempty"`
	// State of the package; valid states are AWAITING_UPLOAD, PROCESSING_UPLOAD, READY, FAILED, COPYING, EXPIRED
	State string `json:"state"`
	// Package type; valid values are bits, docker
	Type string `json:"type"`
	// The time with zone when the object was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type _Package Package

// NewPackage instantiates a new Package object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackage(createdAt time.Time, guid string, state string, type_ string, updatedAt time.Time) *Package {
	this := Package{}
	this.CreatedAt = createdAt
	this.Guid = guid
	this.State = state
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewPackageWithDefaults instantiates a new Package object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWithDefaults() *Package {
	this := Package{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Package) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Package) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Package) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Package) GetData() PackageData {
	if o == nil || IsNil(o.Data) {
		var ret PackageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetDataOk() (*PackageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Package) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PackageData and assigns it to the Data field.
func (o *Package) SetData(v PackageData) {
	o.Data = &v
}

// GetGuid returns the Guid field value
func (o *Package) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *Package) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *Package) SetGuid(v string) {
	o.Guid = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Package) GetLinks() PackageLinks {
	if o == nil || IsNil(o.Links) {
		var ret PackageLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetLinksOk() (*PackageLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Package) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PackageLinks and assigns it to the Links field.
func (o *Package) SetLinks(v PackageLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Package) GetMetadata() V3PackagesPostRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3PackagesPostRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetMetadataOk() (*V3PackagesPostRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Package) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3PackagesPostRequestMetadata and assigns it to the Metadata field.
func (o *Package) SetMetadata(v V3PackagesPostRequestMetadata) {
	o.Metadata = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Package) GetRelationships() V3PackagesPostRequestRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret V3PackagesPostRequestRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetRelationshipsOk() (*V3PackagesPostRequestRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Package) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given V3PackagesPostRequestRelationships and assigns it to the Relationships field.
func (o *Package) SetRelationships(v V3PackagesPostRequestRelationships) {
	o.Relationships = &v
}

// GetState returns the State field value
func (o *Package) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Package) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Package) SetState(v string) {
	o.State = v
}

// GetType returns the Type field value
func (o *Package) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Package) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Package) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Package) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Package) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Package) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Package) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Package) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["guid"] = o.Guid
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["state"] = o.State
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *Package) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"guid",
		"state",
		"type",
		"updated_at",
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

	varPackage := _Package{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPackage)

	if err != nil {
		return err
	}

	*o = Package(varPackage)

	return err
}

type NullablePackage struct {
	value *Package
	isSet bool
}

func (v NullablePackage) Get() *Package {
	return v.value
}

func (v *NullablePackage) Set(val *Package) {
	v.value = val
	v.isSet = true
}

func (v NullablePackage) IsSet() bool {
	return v.isSet
}

func (v *NullablePackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackage(val *Package) *NullablePackage {
	return &NullablePackage{value: val, isSet: true}
}

func (v NullablePackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

