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

// checks if the PackageLinksUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageLinksUpload{}

// PackageLinksUpload struct for PackageLinksUpload
type PackageLinksUpload struct {
	// URL to upload the package
	Href *string `json:"href,omitempty"`
	// HTTP method for the upload URL
	Method *string `json:"method,omitempty"`
}

// NewPackageLinksUpload instantiates a new PackageLinksUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLinksUpload() *PackageLinksUpload {
	this := PackageLinksUpload{}
	return &this
}

// NewPackageLinksUploadWithDefaults instantiates a new PackageLinksUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLinksUploadWithDefaults() *PackageLinksUpload {
	this := PackageLinksUpload{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PackageLinksUpload) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLinksUpload) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PackageLinksUpload) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PackageLinksUpload) SetHref(v string) {
	o.Href = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PackageLinksUpload) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLinksUpload) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PackageLinksUpload) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PackageLinksUpload) SetMethod(v string) {
	o.Method = &v
}

func (o PackageLinksUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageLinksUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullablePackageLinksUpload struct {
	value *PackageLinksUpload
	isSet bool
}

func (v NullablePackageLinksUpload) Get() *PackageLinksUpload {
	return v.value
}

func (v *NullablePackageLinksUpload) Set(val *PackageLinksUpload) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLinksUpload) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLinksUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLinksUpload(val *PackageLinksUpload) *NullablePackageLinksUpload {
	return &NullablePackageLinksUpload{value: val, isSet: true}
}

func (v NullablePackageLinksUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLinksUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


