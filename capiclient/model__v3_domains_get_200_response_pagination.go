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

// checks if the V3DomainsGet200ResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3DomainsGet200ResponsePagination{}

// V3DomainsGet200ResponsePagination struct for V3DomainsGet200ResponsePagination
type V3DomainsGet200ResponsePagination struct {
	First *Get200ResponseLinksLogCache `json:"first,omitempty"`
	Last *Get200ResponseLinksLogCache `json:"last,omitempty"`
	Next *Get200ResponseLinksLogCache `json:"next,omitempty"`
	Previous *Get200ResponseLinksLogCache `json:"previous,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	TotalResults *int32 `json:"total_results,omitempty"`
}

// NewV3DomainsGet200ResponsePagination instantiates a new V3DomainsGet200ResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3DomainsGet200ResponsePagination() *V3DomainsGet200ResponsePagination {
	this := V3DomainsGet200ResponsePagination{}
	return &this
}

// NewV3DomainsGet200ResponsePaginationWithDefaults instantiates a new V3DomainsGet200ResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3DomainsGet200ResponsePaginationWithDefaults() *V3DomainsGet200ResponsePagination {
	this := V3DomainsGet200ResponsePagination{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetFirst() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.First) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetFirstOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given Get200ResponseLinksLogCache and assigns it to the First field.
func (o *V3DomainsGet200ResponsePagination) SetFirst(v Get200ResponseLinksLogCache) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetLast() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Last) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetLastOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Last field.
func (o *V3DomainsGet200ResponsePagination) SetLast(v Get200ResponseLinksLogCache) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetNext() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Next) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetNextOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Next field.
func (o *V3DomainsGet200ResponsePagination) SetNext(v Get200ResponseLinksLogCache) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetPrevious() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Previous) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetPreviousOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Previous field.
func (o *V3DomainsGet200ResponsePagination) SetPrevious(v Get200ResponseLinksLogCache) {
	o.Previous = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *V3DomainsGet200ResponsePagination) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *V3DomainsGet200ResponsePagination) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3DomainsGet200ResponsePagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *V3DomainsGet200ResponsePagination) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *V3DomainsGet200ResponsePagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o V3DomainsGet200ResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3DomainsGet200ResponsePagination) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Previous) {
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

type NullableV3DomainsGet200ResponsePagination struct {
	value *V3DomainsGet200ResponsePagination
	isSet bool
}

func (v NullableV3DomainsGet200ResponsePagination) Get() *V3DomainsGet200ResponsePagination {
	return v.value
}

func (v *NullableV3DomainsGet200ResponsePagination) Set(val *V3DomainsGet200ResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableV3DomainsGet200ResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableV3DomainsGet200ResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3DomainsGet200ResponsePagination(val *V3DomainsGet200ResponsePagination) *NullableV3DomainsGet200ResponsePagination {
	return &NullableV3DomainsGet200ResponsePagination{value: val, isSet: true}
}

func (v NullableV3DomainsGet200ResponsePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3DomainsGet200ResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


