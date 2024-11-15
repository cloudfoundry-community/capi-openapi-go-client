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

// checks if the AppUsageEventListPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUsageEventListPagination{}

// AppUsageEventListPagination struct for AppUsageEventListPagination
type AppUsageEventListPagination struct {
	First *AppUsageEventListPaginationFirst `json:"first,omitempty"`
	Last *AppUsageEventListPaginationLast `json:"last,omitempty"`
	Next *AppUsageEventListPaginationNext `json:"next,omitempty"`
	// Link to the previous page, if applicable
	Previous map[string]interface{} `json:"previous,omitempty"`
	// Total number of pages
	TotalPages *int32 `json:"total_pages,omitempty"`
	// Total number of results
	TotalResults *int32 `json:"total_results,omitempty"`
}

// NewAppUsageEventListPagination instantiates a new AppUsageEventListPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUsageEventListPagination() *AppUsageEventListPagination {
	this := AppUsageEventListPagination{}
	return &this
}

// NewAppUsageEventListPaginationWithDefaults instantiates a new AppUsageEventListPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUsageEventListPaginationWithDefaults() *AppUsageEventListPagination {
	this := AppUsageEventListPagination{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *AppUsageEventListPagination) GetFirst() AppUsageEventListPaginationFirst {
	if o == nil || IsNil(o.First) {
		var ret AppUsageEventListPaginationFirst
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPagination) GetFirstOk() (*AppUsageEventListPaginationFirst, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given AppUsageEventListPaginationFirst and assigns it to the First field.
func (o *AppUsageEventListPagination) SetFirst(v AppUsageEventListPaginationFirst) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *AppUsageEventListPagination) GetLast() AppUsageEventListPaginationLast {
	if o == nil || IsNil(o.Last) {
		var ret AppUsageEventListPaginationLast
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPagination) GetLastOk() (*AppUsageEventListPaginationLast, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given AppUsageEventListPaginationLast and assigns it to the Last field.
func (o *AppUsageEventListPagination) SetLast(v AppUsageEventListPaginationLast) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *AppUsageEventListPagination) GetNext() AppUsageEventListPaginationNext {
	if o == nil || IsNil(o.Next) {
		var ret AppUsageEventListPaginationNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPagination) GetNextOk() (*AppUsageEventListPaginationNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given AppUsageEventListPaginationNext and assigns it to the Next field.
func (o *AppUsageEventListPagination) SetNext(v AppUsageEventListPaginationNext) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUsageEventListPagination) GetPrevious() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUsageEventListPagination) GetPreviousOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Previous) {
		return map[string]interface{}{}, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given map[string]interface{} and assigns it to the Previous field.
func (o *AppUsageEventListPagination) SetPrevious(v map[string]interface{}) {
	o.Previous = v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *AppUsageEventListPagination) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *AppUsageEventListPagination) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *AppUsageEventListPagination) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageEventListPagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *AppUsageEventListPagination) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *AppUsageEventListPagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o AppUsageEventListPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUsageEventListPagination) ToMap() (map[string]interface{}, error) {
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
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["total_results"] = o.TotalResults
	}
	return toSerialize, nil
}

type NullableAppUsageEventListPagination struct {
	value *AppUsageEventListPagination
	isSet bool
}

func (v NullableAppUsageEventListPagination) Get() *AppUsageEventListPagination {
	return v.value
}

func (v *NullableAppUsageEventListPagination) Set(val *AppUsageEventListPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUsageEventListPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUsageEventListPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUsageEventListPagination(val *AppUsageEventListPagination) *NullableAppUsageEventListPagination {
	return &NullableAppUsageEventListPagination{value: val, isSet: true}
}

func (v NullableAppUsageEventListPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUsageEventListPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


