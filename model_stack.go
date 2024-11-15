/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
	"time"
)

// checks if the Stack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stack{}

// Stack struct for Stack
type Stack struct {
	BuildRootfsImage *string `json:"build_rootfs_image,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Default *bool `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Links *ServiceUsageEventLinks `json:"links,omitempty"`
	Metadata *V3AppsGuidTasksPostRequestMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	RunRootfsImage *string `json:"run_rootfs_image,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewStack instantiates a new Stack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStack() *Stack {
	this := Stack{}
	return &this
}

// NewStackWithDefaults instantiates a new Stack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackWithDefaults() *Stack {
	this := Stack{}
	return &this
}

// GetBuildRootfsImage returns the BuildRootfsImage field value if set, zero value otherwise.
func (o *Stack) GetBuildRootfsImage() string {
	if o == nil || IsNil(o.BuildRootfsImage) {
		var ret string
		return ret
	}
	return *o.BuildRootfsImage
}

// GetBuildRootfsImageOk returns a tuple with the BuildRootfsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetBuildRootfsImageOk() (*string, bool) {
	if o == nil || IsNil(o.BuildRootfsImage) {
		return nil, false
	}
	return o.BuildRootfsImage, true
}

// HasBuildRootfsImage returns a boolean if a field has been set.
func (o *Stack) HasBuildRootfsImage() bool {
	if o != nil && !IsNil(o.BuildRootfsImage) {
		return true
	}

	return false
}

// SetBuildRootfsImage gets a reference to the given string and assigns it to the BuildRootfsImage field.
func (o *Stack) SetBuildRootfsImage(v string) {
	o.BuildRootfsImage = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Stack) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Stack) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Stack) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Stack) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Stack) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Stack) SetDefault(v bool) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Stack) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Stack) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Stack) SetDescription(v string) {
	o.Description = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Stack) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Stack) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Stack) SetGuid(v string) {
	o.Guid = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Stack) GetLinks() ServiceUsageEventLinks {
	if o == nil || IsNil(o.Links) {
		var ret ServiceUsageEventLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetLinksOk() (*ServiceUsageEventLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Stack) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ServiceUsageEventLinks and assigns it to the Links field.
func (o *Stack) SetLinks(v ServiceUsageEventLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Stack) GetMetadata() V3AppsGuidTasksPostRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret V3AppsGuidTasksPostRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Stack) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V3AppsGuidTasksPostRequestMetadata and assigns it to the Metadata field.
func (o *Stack) SetMetadata(v V3AppsGuidTasksPostRequestMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Stack) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Stack) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Stack) SetName(v string) {
	o.Name = &v
}

// GetRunRootfsImage returns the RunRootfsImage field value if set, zero value otherwise.
func (o *Stack) GetRunRootfsImage() string {
	if o == nil || IsNil(o.RunRootfsImage) {
		var ret string
		return ret
	}
	return *o.RunRootfsImage
}

// GetRunRootfsImageOk returns a tuple with the RunRootfsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetRunRootfsImageOk() (*string, bool) {
	if o == nil || IsNil(o.RunRootfsImage) {
		return nil, false
	}
	return o.RunRootfsImage, true
}

// HasRunRootfsImage returns a boolean if a field has been set.
func (o *Stack) HasRunRootfsImage() bool {
	if o != nil && !IsNil(o.RunRootfsImage) {
		return true
	}

	return false
}

// SetRunRootfsImage gets a reference to the given string and assigns it to the RunRootfsImage field.
func (o *Stack) SetRunRootfsImage(v string) {
	o.RunRootfsImage = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Stack) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Stack) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Stack) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Stack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildRootfsImage) {
		toSerialize["build_rootfs_image"] = o.BuildRootfsImage
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RunRootfsImage) {
		toSerialize["run_rootfs_image"] = o.RunRootfsImage
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableStack struct {
	value *Stack
	isSet bool
}

func (v NullableStack) Get() *Stack {
	return v.value
}

func (v *NullableStack) Set(val *Stack) {
	v.value = val
	v.isSet = true
}

func (v NullableStack) IsSet() bool {
	return v.isSet
}

func (v *NullableStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStack(val *Stack) *NullableStack {
	return &NullableStack{value: val, isSet: true}
}

func (v NullableStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

