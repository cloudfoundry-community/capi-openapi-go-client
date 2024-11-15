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

// checks if the ServiceCredentialBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceCredentialBinding{}

// ServiceCredentialBinding struct for ServiceCredentialBinding
type ServiceCredentialBinding struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Guid *string `json:"guid,omitempty"`
	LastOperation *ServiceCredentialBindingLastOperation `json:"last_operation,omitempty"`
	Links *ServiceCredentialBindingLinks `json:"links,omitempty"`
	Metadata *V3AppsGuidDropletsCurrentGet200ResponseMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationships *ServiceCredentialBindingRelationships `json:"relationships,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewServiceCredentialBinding instantiates a new ServiceCredentialBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCredentialBinding() *ServiceCredentialBinding {
	this := ServiceCredentialBinding{}
	return &this
}

// NewServiceCredentialBindingWithDefaults instantiates a new ServiceCredentialBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCredentialBindingWithDefaults() *ServiceCredentialBinding {
	this := ServiceCredentialBinding{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceCredentialBinding) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ServiceCredentialBinding) SetGuid(v string) {
	o.Guid = &v
}

// GetLastOperation returns the LastOperation field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetLastOperation() ServiceCredentialBindingLastOperation {
	if o == nil || IsNil(o.LastOperation) {
		var ret ServiceCredentialBindingLastOperation
		return ret
	}
	return *o.LastOperation
}

// GetLastOperationOk returns a tuple with the LastOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetLastOperationOk() (*ServiceCredentialBindingLastOperation, bool) {
	if o == nil || IsNil(o.LastOperation) {
		return nil, false
	}
	return o.LastOperation, true
}

// HasLastOperation returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasLastOperation() bool {
	if o != nil && !IsNil(o.LastOperation) {
		return true
	}

	return false
}

// SetLastOperation gets a reference to the given ServiceCredentialBindingLastOperation and assigns it to the LastOperation field.
func (o *ServiceCredentialBinding) SetLastOperation(v ServiceCredentialBindingLastOperation) {
	o.LastOperation = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetLinks() ServiceCredentialBindingLinks {
	if o == nil || IsNil(o.Links) {
		var ret ServiceCredentialBindingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetLinksOk() (*ServiceCredentialBindingLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ServiceCredentialBindingLinks and assigns it to the Links field.
func (o *ServiceCredentialBinding) SetLinks(v ServiceCredentialBindingLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseMetadata and assigns it to the Metadata field.
func (o *ServiceCredentialBinding) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceCredentialBinding) SetName(v string) {
	o.Name = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetRelationships() ServiceCredentialBindingRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ServiceCredentialBindingRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetRelationshipsOk() (*ServiceCredentialBindingRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceCredentialBindingRelationships and assigns it to the Relationships field.
func (o *ServiceCredentialBinding) SetRelationships(v ServiceCredentialBindingRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceCredentialBinding) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceCredentialBinding) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBinding) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceCredentialBinding) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ServiceCredentialBinding) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ServiceCredentialBinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceCredentialBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.LastOperation) {
		toSerialize["last_operation"] = o.LastOperation
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableServiceCredentialBinding struct {
	value *ServiceCredentialBinding
	isSet bool
}

func (v NullableServiceCredentialBinding) Get() *ServiceCredentialBinding {
	return v.value
}

func (v *NullableServiceCredentialBinding) Set(val *ServiceCredentialBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCredentialBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCredentialBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCredentialBinding(val *ServiceCredentialBinding) *NullableServiceCredentialBinding {
	return &NullableServiceCredentialBinding{value: val, isSet: true}
}

func (v NullableServiceCredentialBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCredentialBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


