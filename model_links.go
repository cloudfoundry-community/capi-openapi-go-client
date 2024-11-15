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

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Links{}

// Links struct for Links
type Links struct {
	Parameters *Link `json:"parameters,omitempty"`
	Self *Link `json:"self,omitempty"`
	ServiceCredentialBindings *Link `json:"service_credential_bindings,omitempty"`
	ServicePlan *Link `json:"service_plan,omitempty"`
	ServiceRouteBindings *Link `json:"service_route_bindings,omitempty"`
	SharedSpaces *Link `json:"shared_spaces,omitempty"`
	Space *Link `json:"space,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Links) GetParameters() Link {
	if o == nil || IsNil(o.Parameters) {
		var ret Link
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetParametersOk() (*Link, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Links) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given Link and assigns it to the Parameters field.
func (o *Links) SetParameters(v Link) {
	o.Parameters = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Links) GetSelf() Link {
	if o == nil || IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetSelfOk() (*Link, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Links) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *Links) SetSelf(v Link) {
	o.Self = &v
}

// GetServiceCredentialBindings returns the ServiceCredentialBindings field value if set, zero value otherwise.
func (o *Links) GetServiceCredentialBindings() Link {
	if o == nil || IsNil(o.ServiceCredentialBindings) {
		var ret Link
		return ret
	}
	return *o.ServiceCredentialBindings
}

// GetServiceCredentialBindingsOk returns a tuple with the ServiceCredentialBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetServiceCredentialBindingsOk() (*Link, bool) {
	if o == nil || IsNil(o.ServiceCredentialBindings) {
		return nil, false
	}
	return o.ServiceCredentialBindings, true
}

// HasServiceCredentialBindings returns a boolean if a field has been set.
func (o *Links) HasServiceCredentialBindings() bool {
	if o != nil && !IsNil(o.ServiceCredentialBindings) {
		return true
	}

	return false
}

// SetServiceCredentialBindings gets a reference to the given Link and assigns it to the ServiceCredentialBindings field.
func (o *Links) SetServiceCredentialBindings(v Link) {
	o.ServiceCredentialBindings = &v
}

// GetServicePlan returns the ServicePlan field value if set, zero value otherwise.
func (o *Links) GetServicePlan() Link {
	if o == nil || IsNil(o.ServicePlan) {
		var ret Link
		return ret
	}
	return *o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetServicePlanOk() (*Link, bool) {
	if o == nil || IsNil(o.ServicePlan) {
		return nil, false
	}
	return o.ServicePlan, true
}

// HasServicePlan returns a boolean if a field has been set.
func (o *Links) HasServicePlan() bool {
	if o != nil && !IsNil(o.ServicePlan) {
		return true
	}

	return false
}

// SetServicePlan gets a reference to the given Link and assigns it to the ServicePlan field.
func (o *Links) SetServicePlan(v Link) {
	o.ServicePlan = &v
}

// GetServiceRouteBindings returns the ServiceRouteBindings field value if set, zero value otherwise.
func (o *Links) GetServiceRouteBindings() Link {
	if o == nil || IsNil(o.ServiceRouteBindings) {
		var ret Link
		return ret
	}
	return *o.ServiceRouteBindings
}

// GetServiceRouteBindingsOk returns a tuple with the ServiceRouteBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetServiceRouteBindingsOk() (*Link, bool) {
	if o == nil || IsNil(o.ServiceRouteBindings) {
		return nil, false
	}
	return o.ServiceRouteBindings, true
}

// HasServiceRouteBindings returns a boolean if a field has been set.
func (o *Links) HasServiceRouteBindings() bool {
	if o != nil && !IsNil(o.ServiceRouteBindings) {
		return true
	}

	return false
}

// SetServiceRouteBindings gets a reference to the given Link and assigns it to the ServiceRouteBindings field.
func (o *Links) SetServiceRouteBindings(v Link) {
	o.ServiceRouteBindings = &v
}

// GetSharedSpaces returns the SharedSpaces field value if set, zero value otherwise.
func (o *Links) GetSharedSpaces() Link {
	if o == nil || IsNil(o.SharedSpaces) {
		var ret Link
		return ret
	}
	return *o.SharedSpaces
}

// GetSharedSpacesOk returns a tuple with the SharedSpaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetSharedSpacesOk() (*Link, bool) {
	if o == nil || IsNil(o.SharedSpaces) {
		return nil, false
	}
	return o.SharedSpaces, true
}

// HasSharedSpaces returns a boolean if a field has been set.
func (o *Links) HasSharedSpaces() bool {
	if o != nil && !IsNil(o.SharedSpaces) {
		return true
	}

	return false
}

// SetSharedSpaces gets a reference to the given Link and assigns it to the SharedSpaces field.
func (o *Links) SetSharedSpaces(v Link) {
	o.SharedSpaces = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *Links) GetSpace() Link {
	if o == nil || IsNil(o.Space) {
		var ret Link
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetSpaceOk() (*Link, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *Links) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given Link and assigns it to the Space field.
func (o *Links) SetSpace(v Link) {
	o.Space = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.ServiceCredentialBindings) {
		toSerialize["service_credential_bindings"] = o.ServiceCredentialBindings
	}
	if !IsNil(o.ServicePlan) {
		toSerialize["service_plan"] = o.ServicePlan
	}
	if !IsNil(o.ServiceRouteBindings) {
		toSerialize["service_route_bindings"] = o.ServiceRouteBindings
	}
	if !IsNil(o.SharedSpaces) {
		toSerialize["shared_spaces"] = o.SharedSpaces
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	return toSerialize, nil
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


