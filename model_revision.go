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

// checks if the Revision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Revision{}

// Revision struct for Revision
type Revision struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Deployable *bool `json:"deployable,omitempty"`
	Description *string `json:"description,omitempty"`
	Droplet *V3AppsPostRequestRelationshipsSpaceData `json:"droplet,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Links *RevisionLinks `json:"links,omitempty"`
	Metadata *V3AppsGuidTasksPostRequestMetadata `json:"metadata,omitempty"`
	Processes *map[string]RevisionProcessesValue `json:"processes,omitempty"`
	Relationships *V3AppsGuidDropletsCurrentGet200ResponseRelationships `json:"relationships,omitempty"`
	Sidecars []Sidecar `json:"sidecars,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewRevision instantiates a new Revision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevision() *Revision {
	this := Revision{}
	return &this
}

// NewRevisionWithDefaults instantiates a new Revision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionWithDefaults() *Revision {
	this := Revision{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Revision) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Revision) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Revision) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeployable returns the Deployable field value if set, zero value otherwise.
func (o *Revision) GetDeployable() bool {
	if o == nil || IsNil(o.Deployable) {
		var ret bool
		return ret
	}
	return *o.Deployable
}

// GetDeployableOk returns a tuple with the Deployable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetDeployableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deployable) {
		return nil, false
	}
	return o.Deployable, true
}

// HasDeployable returns a boolean if a field has been set.
func (o *Revision) HasDeployable() bool {
	if o != nil && !IsNil(o.Deployable) {
		return true
	}

	return false
}

// SetDeployable gets a reference to the given bool and assigns it to the Deployable field.
func (o *Revision) SetDeployable(v bool) {
	o.Deployable = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Revision) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Revision) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Revision) SetDescription(v string) {
	o.Description = &v
}

// GetDroplet returns the Droplet field value if set, zero value otherwise.
func (o *Revision) GetDroplet() V3AppsPostRequestRelationshipsSpaceData {
	if o == nil || IsNil(o.Droplet) {
		var ret V3AppsPostRequestRelationshipsSpaceData
		return ret
	}
	return *o.Droplet
}

// GetDropletOk returns a tuple with the Droplet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetDropletOk() (*V3AppsPostRequestRelationshipsSpaceData, bool) {
	if o == nil || IsNil(o.Droplet) {
		return nil, false
	}
	return o.Droplet, true
}

// HasDroplet returns a boolean if a field has been set.
func (o *Revision) HasDroplet() bool {
	if o != nil && !IsNil(o.Droplet) {
		return true
	}

	return false
}

// SetDroplet gets a reference to the given V3AppsPostRequestRelationshipsSpaceData and assigns it to the Droplet field.
func (o *Revision) SetDroplet(v V3AppsPostRequestRelationshipsSpaceData) {
	o.Droplet = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Revision) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Revision) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Revision) SetGuid(v string) {
	o.Guid = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Revision) GetLinks() RevisionLinks {
	if o == nil || IsNil(o.Links) {
		var ret RevisionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetLinksOk() (*RevisionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Revision) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RevisionLinks and assigns it to the Links field.
func (o *Revision) SetLinks(v RevisionLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Revision) GetMetadata() V3AppsGuidTasksPostRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidTasksPostRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Revision) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidTasksPostRequestMetadata and assigns it to the Metadata field.
func (o *Revision) SetMetadata(v V3AppsGuidTasksPostRequestMetadata) {
	o.Metadata = &v
}

// GetProcesses returns the Processes field value if set, zero value otherwise.
func (o *Revision) GetProcesses() map[string]RevisionProcessesValue {
	if o == nil || IsNil(o.Processes) {
		var ret map[string]RevisionProcessesValue
		return ret
	}
	return *o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetProcessesOk() (*map[string]RevisionProcessesValue, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *Revision) HasProcesses() bool {
	if o != nil && !IsNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given map[string]RevisionProcessesValue and assigns it to the Processes field.
func (o *Revision) SetProcesses(v map[string]RevisionProcessesValue) {
	o.Processes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Revision) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Revision) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseRelationships and assigns it to the Relationships field.
func (o *Revision) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships) {
	o.Relationships = &v
}

// GetSidecars returns the Sidecars field value if set, zero value otherwise.
func (o *Revision) GetSidecars() []Sidecar {
	if o == nil || IsNil(o.Sidecars) {
		var ret []Sidecar
		return ret
	}
	return o.Sidecars
}

// GetSidecarsOk returns a tuple with the Sidecars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetSidecarsOk() ([]Sidecar, bool) {
	if o == nil || IsNil(o.Sidecars) {
		return nil, false
	}
	return o.Sidecars, true
}

// HasSidecars returns a boolean if a field has been set.
func (o *Revision) HasSidecars() bool {
	if o != nil && !IsNil(o.Sidecars) {
		return true
	}

	return false
}

// SetSidecars gets a reference to the given []Sidecar and assigns it to the Sidecars field.
func (o *Revision) SetSidecars(v []Sidecar) {
	o.Sidecars = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Revision) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Revision) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Revision) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Revision) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Revision) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Revision) SetVersion(v int32) {
	o.Version = &v
}

func (o Revision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Revision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Deployable) {
		toSerialize["deployable"] = o.Deployable
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Droplet) {
		toSerialize["droplet"] = o.Droplet
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Processes) {
		toSerialize["processes"] = o.Processes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Sidecars) {
		toSerialize["sidecars"] = o.Sidecars
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRevision struct {
	value *Revision
	isSet bool
}

func (v NullableRevision) Get() *Revision {
	return v.value
}

func (v *NullableRevision) Set(val *Revision) {
	v.value = val
	v.isSet = true
}

func (v NullableRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevision(val *Revision) *NullableRevision {
	return &NullableRevision{value: val, isSet: true}
}

func (v NullableRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

