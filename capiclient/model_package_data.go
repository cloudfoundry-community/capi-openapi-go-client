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

// checks if the PackageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageData{}

// PackageData struct for PackageData
type PackageData struct {
	Checksum *PackageDataChecksum `json:"checksum,omitempty"`
	// If an error occurs this field will contain the error message
	Error NullableString `json:"error,omitempty"`
	// The registry address of the image (for Docker packages)
	Image *string `json:"image,omitempty"`
	// The password for the image's registry (for Docker packages)
	Password *string `json:"password,omitempty"`
	// The username for the image's registry (for Docker packages)
	Username *string `json:"username,omitempty"`
}

// NewPackageData instantiates a new PackageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageData() *PackageData {
	this := PackageData{}
	return &this
}

// NewPackageDataWithDefaults instantiates a new PackageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDataWithDefaults() *PackageData {
	this := PackageData{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *PackageData) GetChecksum() PackageDataChecksum {
	if o == nil || IsNil(o.Checksum) {
		var ret PackageDataChecksum
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageData) GetChecksumOk() (*PackageDataChecksum, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *PackageData) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given PackageDataChecksum and assigns it to the Checksum field.
func (o *PackageData) SetChecksum(v PackageDataChecksum) {
	o.Checksum = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageData) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageData) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *PackageData) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *PackageData) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *PackageData) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *PackageData) UnsetError() {
	o.Error.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PackageData) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageData) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PackageData) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *PackageData) SetImage(v string) {
	o.Image = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PackageData) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageData) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PackageData) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PackageData) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PackageData) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageData) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PackageData) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PackageData) SetUsername(v string) {
	o.Username = &v
}

func (o PackageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullablePackageData struct {
	value *PackageData
	isSet bool
}

func (v NullablePackageData) Get() *PackageData {
	return v.value
}

func (v *NullablePackageData) Set(val *PackageData) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageData) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageData(val *PackageData) *NullablePackageData {
	return &NullablePackageData{value: val, isSet: true}
}

func (v NullablePackageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


