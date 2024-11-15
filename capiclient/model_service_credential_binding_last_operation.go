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

// checks if the ServiceCredentialBindingLastOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceCredentialBindingLastOperation{}

// ServiceCredentialBindingLastOperation struct for ServiceCredentialBindingLastOperation
type ServiceCredentialBindingLastOperation struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewServiceCredentialBindingLastOperation instantiates a new ServiceCredentialBindingLastOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCredentialBindingLastOperation() *ServiceCredentialBindingLastOperation {
	this := ServiceCredentialBindingLastOperation{}
	return &this
}

// NewServiceCredentialBindingLastOperationWithDefaults instantiates a new ServiceCredentialBindingLastOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCredentialBindingLastOperationWithDefaults() *ServiceCredentialBindingLastOperation {
	this := ServiceCredentialBindingLastOperation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceCredentialBindingLastOperation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBindingLastOperation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceCredentialBindingLastOperation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceCredentialBindingLastOperation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceCredentialBindingLastOperation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBindingLastOperation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceCredentialBindingLastOperation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceCredentialBindingLastOperation) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ServiceCredentialBindingLastOperation) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBindingLastOperation) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ServiceCredentialBindingLastOperation) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ServiceCredentialBindingLastOperation) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceCredentialBindingLastOperation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBindingLastOperation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceCredentialBindingLastOperation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceCredentialBindingLastOperation) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceCredentialBindingLastOperation) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCredentialBindingLastOperation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceCredentialBindingLastOperation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ServiceCredentialBindingLastOperation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ServiceCredentialBindingLastOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceCredentialBindingLastOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableServiceCredentialBindingLastOperation struct {
	value *ServiceCredentialBindingLastOperation
	isSet bool
}

func (v NullableServiceCredentialBindingLastOperation) Get() *ServiceCredentialBindingLastOperation {
	return v.value
}

func (v *NullableServiceCredentialBindingLastOperation) Set(val *ServiceCredentialBindingLastOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCredentialBindingLastOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCredentialBindingLastOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCredentialBindingLastOperation(val *ServiceCredentialBindingLastOperation) *NullableServiceCredentialBindingLastOperation {
	return &NullableServiceCredentialBindingLastOperation{value: val, isSet: true}
}

func (v NullableServiceCredentialBindingLastOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCredentialBindingLastOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

