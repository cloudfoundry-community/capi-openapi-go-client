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

// checks if the ServiceBroker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceBroker{}

// ServiceBroker struct for ServiceBroker
type ServiceBroker struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Links *ServiceBrokerLinks `json:"links,omitempty"`
	Metadata *V3AppsGuidDropletsCurrentGet200ResponseMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationships *ServiceBrokerRelationships `json:"relationships,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewServiceBroker instantiates a new ServiceBroker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBroker() *ServiceBroker {
	this := ServiceBroker{}
	return &this
}

// NewServiceBrokerWithDefaults instantiates a new ServiceBroker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBrokerWithDefaults() *ServiceBroker {
	this := ServiceBroker{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceBroker) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceBroker) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceBroker) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ServiceBroker) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ServiceBroker) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ServiceBroker) SetGuid(v string) {
	o.Guid = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceBroker) GetLinks() ServiceBrokerLinks {
	if o == nil || IsNil(o.Links) {
		var ret ServiceBrokerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetLinksOk() (*ServiceBrokerLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceBroker) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ServiceBrokerLinks and assigns it to the Links field.
func (o *ServiceBroker) SetLinks(v ServiceBrokerLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceBroker) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceBroker) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseMetadata and assigns it to the Metadata field.
func (o *ServiceBroker) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceBroker) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceBroker) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceBroker) SetName(v string) {
	o.Name = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceBroker) GetRelationships() ServiceBrokerRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ServiceBrokerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetRelationshipsOk() (*ServiceBrokerRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceBroker) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceBrokerRelationships and assigns it to the Relationships field.
func (o *ServiceBroker) SetRelationships(v ServiceBrokerRelationships) {
	o.Relationships = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceBroker) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceBroker) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ServiceBroker) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceBroker) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBroker) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceBroker) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceBroker) SetUrl(v string) {
	o.Url = &v
}

func (o ServiceBroker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBroker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableServiceBroker struct {
	value *ServiceBroker
	isSet bool
}

func (v NullableServiceBroker) Get() *ServiceBroker {
	return v.value
}

func (v *NullableServiceBroker) Set(val *ServiceBroker) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBroker) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBroker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBroker(val *ServiceBroker) *NullableServiceBroker {
	return &NullableServiceBroker{value: val, isSet: true}
}

func (v NullableServiceBroker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBroker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


