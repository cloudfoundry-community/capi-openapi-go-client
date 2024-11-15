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

// checks if the RevisionsListPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionsListPagination{}

// RevisionsListPagination struct for RevisionsListPagination
type RevisionsListPagination struct {
	First *Get200ResponseLinksLogCache `json:"first,omitempty"`
	Last *Get200ResponseLinksLogCache `json:"last,omitempty"`
	Next NullableV3Get200ResponseLinksServiceInstances `json:"next,omitempty"`
	Previous NullableV3Get200ResponseLinksServiceInstances `json:"previous,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	TotalResults *int32 `json:"total_results,omitempty"`
}

// NewRevisionsListPagination instantiates a new RevisionsListPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionsListPagination() *RevisionsListPagination {
	this := RevisionsListPagination{}
	return &this
}

// NewRevisionsListPaginationWithDefaults instantiates a new RevisionsListPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionsListPaginationWithDefaults() *RevisionsListPagination {
	this := RevisionsListPagination{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *RevisionsListPagination) GetFirst() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.First) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsListPagination) GetFirstOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given Get200ResponseLinksLogCache and assigns it to the First field.
func (o *RevisionsListPagination) SetFirst(v Get200ResponseLinksLogCache) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *RevisionsListPagination) GetLast() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Last) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsListPagination) GetLastOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Last field.
func (o *RevisionsListPagination) SetLast(v Get200ResponseLinksLogCache) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionsListPagination) GetNext() V3Get200ResponseLinksServiceInstances {
	if o == nil || IsNil(o.Next.Get()) {
		var ret V3Get200ResponseLinksServiceInstances
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionsListPagination) GetNextOk() (*V3Get200ResponseLinksServiceInstances, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableV3Get200ResponseLinksServiceInstances and assigns it to the Next field.
func (o *RevisionsListPagination) SetNext(v V3Get200ResponseLinksServiceInstances) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *RevisionsListPagination) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *RevisionsListPagination) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionsListPagination) GetPrevious() V3Get200ResponseLinksServiceInstances {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret V3Get200ResponseLinksServiceInstances
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionsListPagination) GetPreviousOk() (*V3Get200ResponseLinksServiceInstances, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableV3Get200ResponseLinksServiceInstances and assigns it to the Previous field.
func (o *RevisionsListPagination) SetPrevious(v V3Get200ResponseLinksServiceInstances) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *RevisionsListPagination) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *RevisionsListPagination) UnsetPrevious() {
	o.Previous.Unset()
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *RevisionsListPagination) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsListPagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *RevisionsListPagination) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *RevisionsListPagination) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionsListPagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *RevisionsListPagination) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *RevisionsListPagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o RevisionsListPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionsListPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
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

type NullableRevisionsListPagination struct {
	value *RevisionsListPagination
	isSet bool
}

func (v NullableRevisionsListPagination) Get() *RevisionsListPagination {
	return v.value
}

func (v *NullableRevisionsListPagination) Set(val *RevisionsListPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionsListPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionsListPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionsListPagination(val *RevisionsListPagination) *NullableRevisionsListPagination {
	return &NullableRevisionsListPagination{value: val, isSet: true}
}

func (v NullableRevisionsListPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionsListPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

