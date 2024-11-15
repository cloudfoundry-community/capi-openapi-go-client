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

// checks if the ListOrganizationQuotas200ResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationQuotas200ResponsePagination{}

// ListOrganizationQuotas200ResponsePagination struct for ListOrganizationQuotas200ResponsePagination
type ListOrganizationQuotas200ResponsePagination struct {
	First *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"first,omitempty"`
	Last *V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated `json:"last,omitempty"`
	Next NullableListOrganizationQuotas200ResponsePaginationNext `json:"next,omitempty"`
	Previous NullableListOrganizationQuotas200ResponsePaginationNext `json:"previous,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	TotalResults *int32 `json:"total_results,omitempty"`
}

// NewListOrganizationQuotas200ResponsePagination instantiates a new ListOrganizationQuotas200ResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationQuotas200ResponsePagination() *ListOrganizationQuotas200ResponsePagination {
	this := ListOrganizationQuotas200ResponsePagination{}
	return &this
}

// NewListOrganizationQuotas200ResponsePaginationWithDefaults instantiates a new ListOrganizationQuotas200ResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationQuotas200ResponsePaginationWithDefaults() *ListOrganizationQuotas200ResponsePagination {
	this := ListOrganizationQuotas200ResponsePagination{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *ListOrganizationQuotas200ResponsePagination) GetFirst() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.First) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationQuotas200ResponsePagination) GetFirstOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the First field.
func (o *ListOrganizationQuotas200ResponsePagination) SetFirst(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *ListOrganizationQuotas200ResponsePagination) GetLast() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated {
	if o == nil || IsNil(o.Last) {
		var ret V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationQuotas200ResponsePagination) GetLastOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated and assigns it to the Last field.
func (o *ListOrganizationQuotas200ResponsePagination) SetLast(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListOrganizationQuotas200ResponsePagination) GetNext() ListOrganizationQuotas200ResponsePaginationNext {
	if o == nil || IsNil(o.Next.Get()) {
		var ret ListOrganizationQuotas200ResponsePaginationNext
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizationQuotas200ResponsePagination) GetNextOk() (*ListOrganizationQuotas200ResponsePaginationNext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableListOrganizationQuotas200ResponsePaginationNext and assigns it to the Next field.
func (o *ListOrganizationQuotas200ResponsePagination) SetNext(v ListOrganizationQuotas200ResponsePaginationNext) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *ListOrganizationQuotas200ResponsePagination) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *ListOrganizationQuotas200ResponsePagination) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListOrganizationQuotas200ResponsePagination) GetPrevious() ListOrganizationQuotas200ResponsePaginationNext {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret ListOrganizationQuotas200ResponsePaginationNext
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizationQuotas200ResponsePagination) GetPreviousOk() (*ListOrganizationQuotas200ResponsePaginationNext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableListOrganizationQuotas200ResponsePaginationNext and assigns it to the Previous field.
func (o *ListOrganizationQuotas200ResponsePagination) SetPrevious(v ListOrganizationQuotas200ResponsePaginationNext) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *ListOrganizationQuotas200ResponsePagination) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *ListOrganizationQuotas200ResponsePagination) UnsetPrevious() {
	o.Previous.Unset()
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *ListOrganizationQuotas200ResponsePagination) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationQuotas200ResponsePagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *ListOrganizationQuotas200ResponsePagination) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ListOrganizationQuotas200ResponsePagination) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationQuotas200ResponsePagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ListOrganizationQuotas200ResponsePagination) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ListOrganizationQuotas200ResponsePagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o ListOrganizationQuotas200ResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationQuotas200ResponsePagination) ToMap() (map[string]interface{}, error) {
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

type NullableListOrganizationQuotas200ResponsePagination struct {
	value *ListOrganizationQuotas200ResponsePagination
	isSet bool
}

func (v NullableListOrganizationQuotas200ResponsePagination) Get() *ListOrganizationQuotas200ResponsePagination {
	return v.value
}

func (v *NullableListOrganizationQuotas200ResponsePagination) Set(val *ListOrganizationQuotas200ResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationQuotas200ResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationQuotas200ResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationQuotas200ResponsePagination(val *ListOrganizationQuotas200ResponsePagination) *NullableListOrganizationQuotas200ResponsePagination {
	return &NullableListOrganizationQuotas200ResponsePagination{value: val, isSet: true}
}

func (v NullableListOrganizationQuotas200ResponsePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationQuotas200ResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

