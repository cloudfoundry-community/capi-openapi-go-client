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

// checks if the V3AppsGuidRelationshipsCurrentDropletGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3AppsGuidRelationshipsCurrentDropletGet200Response{}

// V3AppsGuidRelationshipsCurrentDropletGet200Response struct for V3AppsGuidRelationshipsCurrentDropletGet200Response
type V3AppsGuidRelationshipsCurrentDropletGet200Response struct {
	Data *V3AppsPostRequestRelationshipsSpaceData `json:"data,omitempty"`
	Links *map[string]Get200ResponseLinksLogCache `json:"links,omitempty"`
}

// NewV3AppsGuidRelationshipsCurrentDropletGet200Response instantiates a new V3AppsGuidRelationshipsCurrentDropletGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AppsGuidRelationshipsCurrentDropletGet200Response() *V3AppsGuidRelationshipsCurrentDropletGet200Response {
	this := V3AppsGuidRelationshipsCurrentDropletGet200Response{}
	return &this
}

// NewV3AppsGuidRelationshipsCurrentDropletGet200ResponseWithDefaults instantiates a new V3AppsGuidRelationshipsCurrentDropletGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AppsGuidRelationshipsCurrentDropletGet200ResponseWithDefaults() *V3AppsGuidRelationshipsCurrentDropletGet200Response {
	this := V3AppsGuidRelationshipsCurrentDropletGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) GetData() V3AppsPostRequestRelationshipsSpaceData {
	if o == nil || IsNil(o.Data) {
		var ret V3AppsPostRequestRelationshipsSpaceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) GetDataOk() (*V3AppsPostRequestRelationshipsSpaceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V3AppsPostRequestRelationshipsSpaceData and assigns it to the Data field.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) SetData(v V3AppsPostRequestRelationshipsSpaceData) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) GetLinks() map[string]Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Links) {
		var ret map[string]Get200ResponseLinksLogCache
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) GetLinksOk() (*map[string]Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Get200ResponseLinksLogCache and assigns it to the Links field.
func (o *V3AppsGuidRelationshipsCurrentDropletGet200Response) SetLinks(v map[string]Get200ResponseLinksLogCache) {
	o.Links = &v
}

func (o V3AppsGuidRelationshipsCurrentDropletGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3AppsGuidRelationshipsCurrentDropletGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableV3AppsGuidRelationshipsCurrentDropletGet200Response struct {
	value *V3AppsGuidRelationshipsCurrentDropletGet200Response
	isSet bool
}

func (v NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) Get() *V3AppsGuidRelationshipsCurrentDropletGet200Response {
	return v.value
}

func (v *NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) Set(val *V3AppsGuidRelationshipsCurrentDropletGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AppsGuidRelationshipsCurrentDropletGet200Response(val *V3AppsGuidRelationshipsCurrentDropletGet200Response) *NullableV3AppsGuidRelationshipsCurrentDropletGet200Response {
	return &NullableV3AppsGuidRelationshipsCurrentDropletGet200Response{value: val, isSet: true}
}

func (v NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AppsGuidRelationshipsCurrentDropletGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


