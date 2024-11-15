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

// checks if the ServiceBrokerCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceBrokerCreate{}

// ServiceBrokerCreate struct for ServiceBrokerCreate
type ServiceBrokerCreate struct {
	Authentication *Authentication `json:"authentication,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationships *ServiceBrokerRelationships `json:"relationships,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewServiceBrokerCreate instantiates a new ServiceBrokerCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBrokerCreate() *ServiceBrokerCreate {
	this := ServiceBrokerCreate{}
	return &this
}

// NewServiceBrokerCreateWithDefaults instantiates a new ServiceBrokerCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBrokerCreateWithDefaults() *ServiceBrokerCreate {
	this := ServiceBrokerCreate{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *ServiceBrokerCreate) GetAuthentication() Authentication {
	if o == nil || IsNil(o.Authentication) {
		var ret Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBrokerCreate) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *ServiceBrokerCreate) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given Authentication and assigns it to the Authentication field.
func (o *ServiceBrokerCreate) SetAuthentication(v Authentication) {
	o.Authentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceBrokerCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBrokerCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceBrokerCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceBrokerCreate) SetName(v string) {
	o.Name = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceBrokerCreate) GetRelationships() ServiceBrokerRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ServiceBrokerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBrokerCreate) GetRelationshipsOk() (*ServiceBrokerRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceBrokerCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceBrokerRelationships and assigns it to the Relationships field.
func (o *ServiceBrokerCreate) SetRelationships(v ServiceBrokerRelationships) {
	o.Relationships = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceBrokerCreate) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBrokerCreate) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceBrokerCreate) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceBrokerCreate) SetUrl(v string) {
	o.Url = &v
}

func (o ServiceBrokerCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBrokerCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableServiceBrokerCreate struct {
	value *ServiceBrokerCreate
	isSet bool
}

func (v NullableServiceBrokerCreate) Get() *ServiceBrokerCreate {
	return v.value
}

func (v *NullableServiceBrokerCreate) Set(val *ServiceBrokerCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBrokerCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBrokerCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBrokerCreate(val *ServiceBrokerCreate) *NullableServiceBrokerCreate {
	return &NullableServiceBrokerCreate{value: val, isSet: true}
}

func (v NullableServiceBrokerCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBrokerCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


