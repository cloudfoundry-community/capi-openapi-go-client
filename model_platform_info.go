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

// checks if the PlatformInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlatformInfo{}

// PlatformInfo struct for PlatformInfo
type PlatformInfo struct {
	Build *string `json:"build,omitempty"`
	CliVersion *PlatformInfoCliVersion `json:"cli_version,omitempty"`
	Custom *map[string]string `json:"custom,omitempty"`
	Description *string `json:"description,omitempty"`
	Links *PlatformInfoLinks `json:"links,omitempty"`
	Name *string `json:"name,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewPlatformInfo instantiates a new PlatformInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformInfo() *PlatformInfo {
	this := PlatformInfo{}
	return &this
}

// NewPlatformInfoWithDefaults instantiates a new PlatformInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformInfoWithDefaults() *PlatformInfo {
	this := PlatformInfo{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *PlatformInfo) GetBuild() string {
	if o == nil || IsNil(o.Build) {
		var ret string
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetBuildOk() (*string, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *PlatformInfo) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given string and assigns it to the Build field.
func (o *PlatformInfo) SetBuild(v string) {
	o.Build = &v
}

// GetCliVersion returns the CliVersion field value if set, zero value otherwise.
func (o *PlatformInfo) GetCliVersion() PlatformInfoCliVersion {
	if o == nil || IsNil(o.CliVersion) {
		var ret PlatformInfoCliVersion
		return ret
	}
	return *o.CliVersion
}

// GetCliVersionOk returns a tuple with the CliVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetCliVersionOk() (*PlatformInfoCliVersion, bool) {
	if o == nil || IsNil(o.CliVersion) {
		return nil, false
	}
	return o.CliVersion, true
}

// HasCliVersion returns a boolean if a field has been set.
func (o *PlatformInfo) HasCliVersion() bool {
	if o != nil && !IsNil(o.CliVersion) {
		return true
	}

	return false
}

// SetCliVersion gets a reference to the given PlatformInfoCliVersion and assigns it to the CliVersion field.
func (o *PlatformInfo) SetCliVersion(v PlatformInfoCliVersion) {
	o.CliVersion = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *PlatformInfo) GetCustom() map[string]string {
	if o == nil || IsNil(o.Custom) {
		var ret map[string]string
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetCustomOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *PlatformInfo) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]string and assigns it to the Custom field.
func (o *PlatformInfo) SetCustom(v map[string]string) {
	o.Custom = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlatformInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlatformInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlatformInfo) SetDescription(v string) {
	o.Description = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PlatformInfo) GetLinks() PlatformInfoLinks {
	if o == nil || IsNil(o.Links) {
		var ret PlatformInfoLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetLinksOk() (*PlatformInfoLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PlatformInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PlatformInfoLinks and assigns it to the Links field.
func (o *PlatformInfo) SetLinks(v PlatformInfoLinks) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlatformInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlatformInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlatformInfo) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PlatformInfo) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInfo) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PlatformInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *PlatformInfo) SetVersion(v int32) {
	o.Version = &v
}

func (o PlatformInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !IsNil(o.CliVersion) {
		toSerialize["cli_version"] = o.CliVersion
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePlatformInfo struct {
	value *PlatformInfo
	isSet bool
}

func (v NullablePlatformInfo) Get() *PlatformInfo {
	return v.value
}

func (v *NullablePlatformInfo) Set(val *PlatformInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformInfo(val *PlatformInfo) *NullablePlatformInfo {
	return &NullablePlatformInfo{value: val, isSet: true}
}

func (v NullablePlatformInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


