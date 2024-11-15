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

// checks if the ApplyOrganizationQuotaToOrganizations201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyOrganizationQuotaToOrganizations201Response{}

// ApplyOrganizationQuotaToOrganizations201Response struct for ApplyOrganizationQuotaToOrganizations201Response
type ApplyOrganizationQuotaToOrganizations201Response struct {
	Data []V3DropletsPostRequestRelationshipsAppData `json:"data,omitempty"`
	Links *ApplyOrganizationQuotaToOrganizations201ResponseLinks `json:"links,omitempty"`
}

// NewApplyOrganizationQuotaToOrganizations201Response instantiates a new ApplyOrganizationQuotaToOrganizations201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyOrganizationQuotaToOrganizations201Response() *ApplyOrganizationQuotaToOrganizations201Response {
	this := ApplyOrganizationQuotaToOrganizations201Response{}
	return &this
}

// NewApplyOrganizationQuotaToOrganizations201ResponseWithDefaults instantiates a new ApplyOrganizationQuotaToOrganizations201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOrganizationQuotaToOrganizations201ResponseWithDefaults() *ApplyOrganizationQuotaToOrganizations201Response {
	this := ApplyOrganizationQuotaToOrganizations201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApplyOrganizationQuotaToOrganizations201Response) GetData() []V3DropletsPostRequestRelationshipsAppData {
	if o == nil || IsNil(o.Data) {
		var ret []V3DropletsPostRequestRelationshipsAppData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrganizationQuotaToOrganizations201Response) GetDataOk() ([]V3DropletsPostRequestRelationshipsAppData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApplyOrganizationQuotaToOrganizations201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []V3DropletsPostRequestRelationshipsAppData and assigns it to the Data field.
func (o *ApplyOrganizationQuotaToOrganizations201Response) SetData(v []V3DropletsPostRequestRelationshipsAppData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplyOrganizationQuotaToOrganizations201Response) GetLinks() ApplyOrganizationQuotaToOrganizations201ResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret ApplyOrganizationQuotaToOrganizations201ResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrganizationQuotaToOrganizations201Response) GetLinksOk() (*ApplyOrganizationQuotaToOrganizations201ResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplyOrganizationQuotaToOrganizations201Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApplyOrganizationQuotaToOrganizations201ResponseLinks and assigns it to the Links field.
func (o *ApplyOrganizationQuotaToOrganizations201Response) SetLinks(v ApplyOrganizationQuotaToOrganizations201ResponseLinks) {
	o.Links = &v
}

func (o ApplyOrganizationQuotaToOrganizations201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOrganizationQuotaToOrganizations201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableApplyOrganizationQuotaToOrganizations201Response struct {
	value *ApplyOrganizationQuotaToOrganizations201Response
	isSet bool
}

func (v NullableApplyOrganizationQuotaToOrganizations201Response) Get() *ApplyOrganizationQuotaToOrganizations201Response {
	return v.value
}

func (v *NullableApplyOrganizationQuotaToOrganizations201Response) Set(val *ApplyOrganizationQuotaToOrganizations201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyOrganizationQuotaToOrganizations201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOrganizationQuotaToOrganizations201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOrganizationQuotaToOrganizations201Response(val *ApplyOrganizationQuotaToOrganizations201Response) *NullableApplyOrganizationQuotaToOrganizations201Response {
	return &NullableApplyOrganizationQuotaToOrganizations201Response{value: val, isSet: true}
}

func (v NullableApplyOrganizationQuotaToOrganizations201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOrganizationQuotaToOrganizations201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

