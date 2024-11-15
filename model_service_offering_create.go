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

// checks if the ServiceOfferingCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOfferingCreate{}

// ServiceOfferingCreate struct for ServiceOfferingCreate
type ServiceOfferingCreate struct {
	Available *bool `json:"available,omitempty"`
	BrokerCatalog *BrokerCatalog `json:"broker_catalog,omitempty"`
	Description *string `json:"description,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationships *ServiceOfferingRelationships `json:"relationships,omitempty"`
	Requires []string `json:"requires,omitempty"`
	Shareable *bool `json:"shareable,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewServiceOfferingCreate instantiates a new ServiceOfferingCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOfferingCreate() *ServiceOfferingCreate {
	this := ServiceOfferingCreate{}
	return &this
}

// NewServiceOfferingCreateWithDefaults instantiates a new ServiceOfferingCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOfferingCreateWithDefaults() *ServiceOfferingCreate {
	this := ServiceOfferingCreate{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *ServiceOfferingCreate) SetAvailable(v bool) {
	o.Available = &v
}

// GetBrokerCatalog returns the BrokerCatalog field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetBrokerCatalog() BrokerCatalog {
	if o == nil || IsNil(o.BrokerCatalog) {
		var ret BrokerCatalog
		return ret
	}
	return *o.BrokerCatalog
}

// GetBrokerCatalogOk returns a tuple with the BrokerCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetBrokerCatalogOk() (*BrokerCatalog, bool) {
	if o == nil || IsNil(o.BrokerCatalog) {
		return nil, false
	}
	return o.BrokerCatalog, true
}

// HasBrokerCatalog returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasBrokerCatalog() bool {
	if o != nil && !IsNil(o.BrokerCatalog) {
		return true
	}

	return false
}

// SetBrokerCatalog gets a reference to the given BrokerCatalog and assigns it to the BrokerCatalog field.
func (o *ServiceOfferingCreate) SetBrokerCatalog(v BrokerCatalog) {
	o.BrokerCatalog = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceOfferingCreate) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *ServiceOfferingCreate) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceOfferingCreate) SetName(v string) {
	o.Name = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetRelationships() ServiceOfferingRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ServiceOfferingRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetRelationshipsOk() (*ServiceOfferingRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceOfferingRelationships and assigns it to the Relationships field.
func (o *ServiceOfferingCreate) SetRelationships(v ServiceOfferingRelationships) {
	o.Relationships = &v
}

// GetRequires returns the Requires field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetRequires() []string {
	if o == nil || IsNil(o.Requires) {
		var ret []string
		return ret
	}
	return o.Requires
}

// GetRequiresOk returns a tuple with the Requires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetRequiresOk() ([]string, bool) {
	if o == nil || IsNil(o.Requires) {
		return nil, false
	}
	return o.Requires, true
}

// HasRequires returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasRequires() bool {
	if o != nil && !IsNil(o.Requires) {
		return true
	}

	return false
}

// SetRequires gets a reference to the given []string and assigns it to the Requires field.
func (o *ServiceOfferingCreate) SetRequires(v []string) {
	o.Requires = v
}

// GetShareable returns the Shareable field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetShareable() bool {
	if o == nil || IsNil(o.Shareable) {
		var ret bool
		return ret
	}
	return *o.Shareable
}

// GetShareableOk returns a tuple with the Shareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetShareableOk() (*bool, bool) {
	if o == nil || IsNil(o.Shareable) {
		return nil, false
	}
	return o.Shareable, true
}

// HasShareable returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasShareable() bool {
	if o != nil && !IsNil(o.Shareable) {
		return true
	}

	return false
}

// SetShareable gets a reference to the given bool and assigns it to the Shareable field.
func (o *ServiceOfferingCreate) SetShareable(v bool) {
	o.Shareable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceOfferingCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOfferingCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceOfferingCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceOfferingCreate) SetTags(v []string) {
	o.Tags = v
}

func (o ServiceOfferingCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOfferingCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.BrokerCatalog) {
		toSerialize["broker_catalog"] = o.BrokerCatalog
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Requires) {
		toSerialize["requires"] = o.Requires
	}
	if !IsNil(o.Shareable) {
		toSerialize["shareable"] = o.Shareable
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableServiceOfferingCreate struct {
	value *ServiceOfferingCreate
	isSet bool
}

func (v NullableServiceOfferingCreate) Get() *ServiceOfferingCreate {
	return v.value
}

func (v *NullableServiceOfferingCreate) Set(val *ServiceOfferingCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOfferingCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOfferingCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOfferingCreate(val *ServiceOfferingCreate) *NullableServiceOfferingCreate {
	return &NullableServiceOfferingCreate{value: val, isSet: true}
}

func (v NullableServiceOfferingCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOfferingCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

