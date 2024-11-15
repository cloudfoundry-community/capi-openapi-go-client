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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Internal *bool `json:"internal,omitempty"`
	Links *DomainLinks `json:"links,omitempty"`
	Metadata *V3AppsGuidDropletsCurrentGet200ResponseMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationships *DomainRelationships `json:"relationships,omitempty"`
	RouterGroup *V3DropletsPostRequestRelationshipsAppData `json:"router_group,omitempty"`
	SupportedProtocols []string `json:"supported_protocols,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Domain) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Domain) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Domain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Domain) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Domain) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Domain) SetGuid(v string) {
	o.Guid = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *Domain) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *Domain) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *Domain) SetInternal(v bool) {
	o.Internal = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Domain) GetLinks() DomainLinks {
	if o == nil || IsNil(o.Links) {
		var ret DomainLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetLinksOk() (*DomainLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Domain) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DomainLinks and assigns it to the Links field.
func (o *Domain) SetLinks(v DomainLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Domain) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Domain) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseMetadata and assigns it to the Metadata field.
func (o *Domain) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Domain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Domain) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Domain) SetName(v string) {
	o.Name = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Domain) GetRelationships() DomainRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret DomainRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRelationshipsOk() (*DomainRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Domain) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DomainRelationships and assigns it to the Relationships field.
func (o *Domain) SetRelationships(v DomainRelationships) {
	o.Relationships = &v
}

// GetRouterGroup returns the RouterGroup field value if set, zero value otherwise.
func (o *Domain) GetRouterGroup() V3DropletsPostRequestRelationshipsAppData {
	if o == nil || IsNil(o.RouterGroup) {
		var ret V3DropletsPostRequestRelationshipsAppData
		return ret
	}
	return *o.RouterGroup
}

// GetRouterGroupOk returns a tuple with the RouterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRouterGroupOk() (*V3DropletsPostRequestRelationshipsAppData, bool) {
	if o == nil || IsNil(o.RouterGroup) {
		return nil, false
	}
	return o.RouterGroup, true
}

// HasRouterGroup returns a boolean if a field has been set.
func (o *Domain) HasRouterGroup() bool {
	if o != nil && !IsNil(o.RouterGroup) {
		return true
	}

	return false
}

// SetRouterGroup gets a reference to the given V3DropletsPostRequestRelationshipsAppData and assigns it to the RouterGroup field.
func (o *Domain) SetRouterGroup(v V3DropletsPostRequestRelationshipsAppData) {
	o.RouterGroup = &v
}

// GetSupportedProtocols returns the SupportedProtocols field value if set, zero value otherwise.
func (o *Domain) GetSupportedProtocols() []string {
	if o == nil || IsNil(o.SupportedProtocols) {
		var ret []string
		return ret
	}
	return o.SupportedProtocols
}

// GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSupportedProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedProtocols) {
		return nil, false
	}
	return o.SupportedProtocols, true
}

// HasSupportedProtocols returns a boolean if a field has been set.
func (o *Domain) HasSupportedProtocols() bool {
	if o != nil && !IsNil(o.SupportedProtocols) {
		return true
	}

	return false
}

// SetSupportedProtocols gets a reference to the given []string and assigns it to the SupportedProtocols field.
func (o *Domain) SetSupportedProtocols(v []string) {
	o.SupportedProtocols = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Domain) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Domain) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Domain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
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
	if !IsNil(o.RouterGroup) {
		toSerialize["router_group"] = o.RouterGroup
	}
	if !IsNil(o.SupportedProtocols) {
		toSerialize["supported_protocols"] = o.SupportedProtocols
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


