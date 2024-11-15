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

// checks if the Pagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pagination{}

// Pagination struct for Pagination
type Pagination struct {
	First Link `json:"first,omitempty"`
	Last Link `json:"last,omitempty"`
	Next Link `json:"next,omitempty"`
	Previous NullableListOrganizationQuotas200ResponsePaginationNext `json:"previous,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	TotalResults *int32 `json:"total_results,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *Pagination) GetFirst() Link {
	if o == nil || IsNil(o.First) {
		var ret Link
		return ret
	}
	return o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetFirstOk() (Link, bool) {
	if o == nil || IsNil(o.First) {
		return Link{}, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *Pagination) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given Link and assigns it to the First field.
func (o *Pagination) SetFirst(v Link) {
	o.First = v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *Pagination) GetLast() Link {
	if o == nil || IsNil(o.Last) {
		var ret Link
		return ret
	}
	return o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetLastOk() (Link, bool) {
	if o == nil || IsNil(o.Last) {
		return Link{}, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *Pagination) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given Link and assigns it to the Last field.
func (o *Pagination) SetLast(v Link) {
	o.Last = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Pagination) GetNext() Link {
	if o == nil || IsNil(o.Next) {
		var ret Link
		return ret
	}
	return o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetNextOk() (Link, bool) {
	if o == nil || IsNil(o.Next) {
		return Link{}, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Pagination) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *Pagination) SetNext(v Link) {
	o.Next = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pagination) GetPrevious() ListOrganizationQuotas200ResponsePaginationNext {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret ListOrganizationQuotas200ResponsePaginationNext
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pagination) GetPreviousOk() (*ListOrganizationQuotas200ResponsePaginationNext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *Pagination) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableListOrganizationQuotas200ResponsePaginationNext and assigns it to the Previous field.
func (o *Pagination) SetPrevious(v ListOrganizationQuotas200ResponsePaginationNext) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *Pagination) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *Pagination) UnsetPrevious() {
	o.Previous.Unset()
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *Pagination) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *Pagination) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *Pagination) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *Pagination) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *Pagination) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *Pagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["total_results"] = o.TotalResults
	}
	return toSerialize, nil
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

