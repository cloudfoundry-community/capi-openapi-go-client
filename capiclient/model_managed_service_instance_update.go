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

// checks if the ManagedServiceInstanceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedServiceInstanceUpdate{}

// ManagedServiceInstanceUpdate struct for ManagedServiceInstanceUpdate
type ManagedServiceInstanceUpdate struct {
	MaintenanceInfo *Get200ResponseLinksCloudControllerV2Meta `json:"maintenance_info,omitempty"`
	Metadata *V3AppsGuidDropletsCurrentGet200ResponseMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Relationships *ManagedServiceInstanceUpdateRelationships `json:"relationships,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewManagedServiceInstanceUpdate instantiates a new ManagedServiceInstanceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedServiceInstanceUpdate() *ManagedServiceInstanceUpdate {
	this := ManagedServiceInstanceUpdate{}
	return &this
}

// NewManagedServiceInstanceUpdateWithDefaults instantiates a new ManagedServiceInstanceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedServiceInstanceUpdateWithDefaults() *ManagedServiceInstanceUpdate {
	this := ManagedServiceInstanceUpdate{}
	return &this
}

// GetMaintenanceInfo returns the MaintenanceInfo field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetMaintenanceInfo() Get200ResponseLinksCloudControllerV2Meta {
	if o == nil || IsNil(o.MaintenanceInfo) {
		var ret Get200ResponseLinksCloudControllerV2Meta
		return ret
	}
	return *o.MaintenanceInfo
}

// GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetMaintenanceInfoOk() (*Get200ResponseLinksCloudControllerV2Meta, bool) {
	if o == nil || IsNil(o.MaintenanceInfo) {
		return nil, false
	}
	return o.MaintenanceInfo, true
}

// HasMaintenanceInfo returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasMaintenanceInfo() bool {
	if o != nil && !IsNil(o.MaintenanceInfo) {
		return true
	}

	return false
}

// SetMaintenanceInfo gets a reference to the given Get200ResponseLinksCloudControllerV2Meta and assigns it to the MaintenanceInfo field.
func (o *ManagedServiceInstanceUpdate) SetMaintenanceInfo(v Get200ResponseLinksCloudControllerV2Meta) {
	o.MaintenanceInfo = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseMetadata and assigns it to the Metadata field.
func (o *ManagedServiceInstanceUpdate) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagedServiceInstanceUpdate) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *ManagedServiceInstanceUpdate) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetRelationships() ManagedServiceInstanceUpdateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ManagedServiceInstanceUpdateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetRelationshipsOk() (*ManagedServiceInstanceUpdateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManagedServiceInstanceUpdateRelationships and assigns it to the Relationships field.
func (o *ManagedServiceInstanceUpdate) SetRelationships(v ManagedServiceInstanceUpdateRelationships) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ManagedServiceInstanceUpdate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServiceInstanceUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ManagedServiceInstanceUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ManagedServiceInstanceUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o ManagedServiceInstanceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedServiceInstanceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaintenanceInfo) {
		toSerialize["maintenance_info"] = o.MaintenanceInfo
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableManagedServiceInstanceUpdate struct {
	value *ManagedServiceInstanceUpdate
	isSet bool
}

func (v NullableManagedServiceInstanceUpdate) Get() *ManagedServiceInstanceUpdate {
	return v.value
}

func (v *NullableManagedServiceInstanceUpdate) Set(val *ManagedServiceInstanceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedServiceInstanceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedServiceInstanceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedServiceInstanceUpdate(val *ManagedServiceInstanceUpdate) *NullableManagedServiceInstanceUpdate {
	return &NullableManagedServiceInstanceUpdate{value: val, isSet: true}
}

func (v NullableManagedServiceInstanceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedServiceInstanceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


