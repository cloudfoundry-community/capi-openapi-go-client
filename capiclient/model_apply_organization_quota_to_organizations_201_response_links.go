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

// checks if the ApplyOrganizationQuotaToOrganizations201ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyOrganizationQuotaToOrganizations201ResponseLinks{}

// ApplyOrganizationQuotaToOrganizations201ResponseLinks struct for ApplyOrganizationQuotaToOrganizations201ResponseLinks
type ApplyOrganizationQuotaToOrganizations201ResponseLinks struct {
	// URL of the applied quota relationship
	Self *string `json:"self,omitempty"`
}

// NewApplyOrganizationQuotaToOrganizations201ResponseLinks instantiates a new ApplyOrganizationQuotaToOrganizations201ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyOrganizationQuotaToOrganizations201ResponseLinks() *ApplyOrganizationQuotaToOrganizations201ResponseLinks {
	this := ApplyOrganizationQuotaToOrganizations201ResponseLinks{}
	return &this
}

// NewApplyOrganizationQuotaToOrganizations201ResponseLinksWithDefaults instantiates a new ApplyOrganizationQuotaToOrganizations201ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOrganizationQuotaToOrganizations201ResponseLinksWithDefaults() *ApplyOrganizationQuotaToOrganizations201ResponseLinks {
	this := ApplyOrganizationQuotaToOrganizations201ResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApplyOrganizationQuotaToOrganizations201ResponseLinks) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrganizationQuotaToOrganizations201ResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApplyOrganizationQuotaToOrganizations201ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ApplyOrganizationQuotaToOrganizations201ResponseLinks) SetSelf(v string) {
	o.Self = &v
}

func (o ApplyOrganizationQuotaToOrganizations201ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOrganizationQuotaToOrganizations201ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableApplyOrganizationQuotaToOrganizations201ResponseLinks struct {
	value *ApplyOrganizationQuotaToOrganizations201ResponseLinks
	isSet bool
}

func (v NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) Get() *ApplyOrganizationQuotaToOrganizations201ResponseLinks {
	return v.value
}

func (v *NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) Set(val *ApplyOrganizationQuotaToOrganizations201ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOrganizationQuotaToOrganizations201ResponseLinks(val *ApplyOrganizationQuotaToOrganizations201ResponseLinks) *NullableApplyOrganizationQuotaToOrganizations201ResponseLinks {
	return &NullableApplyOrganizationQuotaToOrganizations201ResponseLinks{value: val, isSet: true}
}

func (v NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOrganizationQuotaToOrganizations201ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


