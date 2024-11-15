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

// checks if the UserProvidedServiceInstanceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProvidedServiceInstanceUpdate{}

// UserProvidedServiceInstanceUpdate struct for UserProvidedServiceInstanceUpdate
type UserProvidedServiceInstanceUpdate struct {
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	Metadata *V3AppsGuidDropletsCurrentGet200ResponseMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	RouteServiceUrl *string `json:"route_service_url,omitempty"`
	SyslogDrainUrl *string `json:"syslog_drain_url,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewUserProvidedServiceInstanceUpdate instantiates a new UserProvidedServiceInstanceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProvidedServiceInstanceUpdate() *UserProvidedServiceInstanceUpdate {
	this := UserProvidedServiceInstanceUpdate{}
	return &this
}

// NewUserProvidedServiceInstanceUpdateWithDefaults instantiates a new UserProvidedServiceInstanceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProvidedServiceInstanceUpdateWithDefaults() *UserProvidedServiceInstanceUpdate {
	this := UserProvidedServiceInstanceUpdate{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetCredentials() map[string]interface{} {
	if o == nil || IsNil(o.Credentials) {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *UserProvidedServiceInstanceUpdate) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidDropletsCurrentGet200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidDropletsCurrentGet200ResponseMetadata and assigns it to the Metadata field.
func (o *UserProvidedServiceInstanceUpdate) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserProvidedServiceInstanceUpdate) SetName(v string) {
	o.Name = &v
}

// GetRouteServiceUrl returns the RouteServiceUrl field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetRouteServiceUrl() string {
	if o == nil || IsNil(o.RouteServiceUrl) {
		var ret string
		return ret
	}
	return *o.RouteServiceUrl
}

// GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetRouteServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RouteServiceUrl) {
		return nil, false
	}
	return o.RouteServiceUrl, true
}

// HasRouteServiceUrl returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasRouteServiceUrl() bool {
	if o != nil && !IsNil(o.RouteServiceUrl) {
		return true
	}

	return false
}

// SetRouteServiceUrl gets a reference to the given string and assigns it to the RouteServiceUrl field.
func (o *UserProvidedServiceInstanceUpdate) SetRouteServiceUrl(v string) {
	o.RouteServiceUrl = &v
}

// GetSyslogDrainUrl returns the SyslogDrainUrl field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetSyslogDrainUrl() string {
	if o == nil || IsNil(o.SyslogDrainUrl) {
		var ret string
		return ret
	}
	return *o.SyslogDrainUrl
}

// GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetSyslogDrainUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogDrainUrl) {
		return nil, false
	}
	return o.SyslogDrainUrl, true
}

// HasSyslogDrainUrl returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasSyslogDrainUrl() bool {
	if o != nil && !IsNil(o.SyslogDrainUrl) {
		return true
	}

	return false
}

// SetSyslogDrainUrl gets a reference to the given string and assigns it to the SyslogDrainUrl field.
func (o *UserProvidedServiceInstanceUpdate) SetSyslogDrainUrl(v string) {
	o.SyslogDrainUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UserProvidedServiceInstanceUpdate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvidedServiceInstanceUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UserProvidedServiceInstanceUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UserProvidedServiceInstanceUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o UserProvidedServiceInstanceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProvidedServiceInstanceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RouteServiceUrl) {
		toSerialize["route_service_url"] = o.RouteServiceUrl
	}
	if !IsNil(o.SyslogDrainUrl) {
		toSerialize["syslog_drain_url"] = o.SyslogDrainUrl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableUserProvidedServiceInstanceUpdate struct {
	value *UserProvidedServiceInstanceUpdate
	isSet bool
}

func (v NullableUserProvidedServiceInstanceUpdate) Get() *UserProvidedServiceInstanceUpdate {
	return v.value
}

func (v *NullableUserProvidedServiceInstanceUpdate) Set(val *UserProvidedServiceInstanceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProvidedServiceInstanceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProvidedServiceInstanceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProvidedServiceInstanceUpdate(val *UserProvidedServiceInstanceUpdate) *NullableUserProvidedServiceInstanceUpdate {
	return &NullableUserProvidedServiceInstanceUpdate{value: val, isSet: true}
}

func (v NullableUserProvidedServiceInstanceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProvidedServiceInstanceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

