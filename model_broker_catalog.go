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

// checks if the BrokerCatalog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrokerCatalog{}

// BrokerCatalog struct for BrokerCatalog
type BrokerCatalog struct {
	Features *BrokerCatalogFeatures `json:"features,omitempty"`
	Id *string `json:"id,omitempty"`
	MaximumPollingDuration NullableInt32 `json:"maximum_polling_duration,omitempty"`
	Metadata *BrokerCatalogMetadata `json:"metadata,omitempty"`
}

// NewBrokerCatalog instantiates a new BrokerCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerCatalog() *BrokerCatalog {
	this := BrokerCatalog{}
	return &this
}

// NewBrokerCatalogWithDefaults instantiates a new BrokerCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerCatalogWithDefaults() *BrokerCatalog {
	this := BrokerCatalog{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BrokerCatalog) GetFeatures() BrokerCatalogFeatures {
	if o == nil || IsNil(o.Features) {
		var ret BrokerCatalogFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerCatalog) GetFeaturesOk() (*BrokerCatalogFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BrokerCatalog) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given BrokerCatalogFeatures and assigns it to the Features field.
func (o *BrokerCatalog) SetFeatures(v BrokerCatalogFeatures) {
	o.Features = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrokerCatalog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerCatalog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrokerCatalog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrokerCatalog) SetId(v string) {
	o.Id = &v
}

// GetMaximumPollingDuration returns the MaximumPollingDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrokerCatalog) GetMaximumPollingDuration() int32 {
	if o == nil || IsNil(o.MaximumPollingDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumPollingDuration.Get()
}

// GetMaximumPollingDurationOk returns a tuple with the MaximumPollingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrokerCatalog) GetMaximumPollingDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumPollingDuration.Get(), o.MaximumPollingDuration.IsSet()
}

// HasMaximumPollingDuration returns a boolean if a field has been set.
func (o *BrokerCatalog) HasMaximumPollingDuration() bool {
	if o != nil && o.MaximumPollingDuration.IsSet() {
		return true
	}

	return false
}

// SetMaximumPollingDuration gets a reference to the given NullableInt32 and assigns it to the MaximumPollingDuration field.
func (o *BrokerCatalog) SetMaximumPollingDuration(v int32) {
	o.MaximumPollingDuration.Set(&v)
}
// SetMaximumPollingDurationNil sets the value for MaximumPollingDuration to be an explicit nil
func (o *BrokerCatalog) SetMaximumPollingDurationNil() {
	o.MaximumPollingDuration.Set(nil)
}

// UnsetMaximumPollingDuration ensures that no value is present for MaximumPollingDuration, not even an explicit nil
func (o *BrokerCatalog) UnsetMaximumPollingDuration() {
	o.MaximumPollingDuration.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BrokerCatalog) GetMetadata() BrokerCatalogMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret BrokerCatalogMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerCatalog) GetMetadataOk() (*BrokerCatalogMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BrokerCatalog) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given BrokerCatalogMetadata and assigns it to the Metadata field.
func (o *BrokerCatalog) SetMetadata(v BrokerCatalogMetadata) {
	o.Metadata = &v
}

func (o BrokerCatalog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrokerCatalog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.MaximumPollingDuration.IsSet() {
		toSerialize["maximum_polling_duration"] = o.MaximumPollingDuration.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableBrokerCatalog struct {
	value *BrokerCatalog
	isSet bool
}

func (v NullableBrokerCatalog) Get() *BrokerCatalog {
	return v.value
}

func (v *NullableBrokerCatalog) Set(val *BrokerCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerCatalog(val *BrokerCatalog) *NullableBrokerCatalog {
	return &NullableBrokerCatalog{value: val, isSet: true}
}

func (v NullableBrokerCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


