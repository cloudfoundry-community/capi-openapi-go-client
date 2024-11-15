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
)

// checks if the Sidecar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sidecar{}

// Sidecar struct for Sidecar
type Sidecar struct {
	Command *string `json:"command,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Guid *string `json:"guid,omitempty"`
	MemoryInMb *int32 `json:"memory_in_mb,omitempty"`
	Name *string `json:"name,omitempty"`
	Origin *string `json:"origin,omitempty"`
	ProcessTypes []string `json:"process_types,omitempty"`
	Relationships *V3AppsGuidDropletsCurrentGet200ResponseRelationships `json:"relationships,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSidecar instantiates a new Sidecar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSidecar() *Sidecar {
	this := Sidecar{}
	return &this
}

// NewSidecarWithDefaults instantiates a new Sidecar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSidecarWithDefaults() *Sidecar {
	this := Sidecar{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *Sidecar) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *Sidecar) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *Sidecar) SetCommand(v string) {
	o.Command = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Sidecar) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Sidecar) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Sidecar) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Sidecar) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Sidecar) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Sidecar) SetGuid(v string) {
	o.Guid = &v
}

// GetMemoryInMb returns the MemoryInMb field value if set, zero value otherwise.
func (o *Sidecar) GetMemoryInMb() int32 {
	if o == nil || IsNil(o.MemoryInMb) {
		var ret int32
		return ret
	}
	return *o.MemoryInMb
}

// GetMemoryInMbOk returns a tuple with the MemoryInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetMemoryInMbOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryInMb) {
		return nil, false
	}
	return o.MemoryInMb, true
}

// HasMemoryInMb returns a boolean if a field has been set.
func (o *Sidecar) HasMemoryInMb() bool {
	if o != nil && !IsNil(o.MemoryInMb) {
		return true
	}

	return false
}

// SetMemoryInMb gets a reference to the given int32 and assigns it to the MemoryInMb field.
func (o *Sidecar) SetMemoryInMb(v int32) {
	o.MemoryInMb = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Sidecar) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Sidecar) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Sidecar) SetName(v string) {
	o.Name = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *Sidecar) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *Sidecar) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *Sidecar) SetOrigin(v string) {
	o.Origin = &v
}

// GetProcessTypes returns the ProcessTypes field value if set, zero value otherwise.
func (o *Sidecar) GetProcessTypes() []string {
	if o == nil || IsNil(o.ProcessTypes) {
		var ret []string
		return ret
	}
	return o.ProcessTypes
}

// GetProcessTypesOk returns a tuple with the ProcessTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetProcessTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessTypes) {
		return nil, false
	}
	return o.ProcessTypes, true
}

// HasProcessTypes returns a boolean if a field has been set.
func (o *Sidecar) HasProcessTypes() bool {
	if o != nil && !IsNil(o.ProcessTypes) {
		return true
	}

	return false
}

// SetProcessTypes gets a reference to the given []string and assigns it to the ProcessTypes field.
func (o *Sidecar) SetProcessTypes(v []string) {
	o.ProcessTypes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Sidecar) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Sidecar) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseRelationships and assigns it to the Relationships field.
func (o *Sidecar) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships) {
	o.Relationships = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Sidecar) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sidecar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Sidecar) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Sidecar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Sidecar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sidecar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.MemoryInMb) {
		toSerialize["memory_in_mb"] = o.MemoryInMb
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.ProcessTypes) {
		toSerialize["process_types"] = o.ProcessTypes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSidecar struct {
	value *Sidecar
	isSet bool
}

func (v NullableSidecar) Get() *Sidecar {
	return v.value
}

func (v *NullableSidecar) Set(val *Sidecar) {
	v.value = val
	v.isSet = true
}

func (v NullableSidecar) IsSet() bool {
	return v.isSet
}

func (v *NullableSidecar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSidecar(val *Sidecar) *NullableSidecar {
	return &NullableSidecar{value: val, isSet: true}
}

func (v NullableSidecar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSidecar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


