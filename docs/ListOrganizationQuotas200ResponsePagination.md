# ListOrganizationQuotas200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated.md) |  | [optional] 
**Last** | Pointer to [**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated.md) |  | [optional] 
**Next** | Pointer to [**NullableListOrganizationQuotas200ResponsePaginationNext**](ListOrganizationQuotas200ResponsePaginationNext.md) |  | [optional] 
**Previous** | Pointer to [**NullableListOrganizationQuotas200ResponsePaginationNext**](ListOrganizationQuotas200ResponsePaginationNext.md) |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewListOrganizationQuotas200ResponsePagination

`func NewListOrganizationQuotas200ResponsePagination() *ListOrganizationQuotas200ResponsePagination`

NewListOrganizationQuotas200ResponsePagination instantiates a new ListOrganizationQuotas200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationQuotas200ResponsePaginationWithDefaults

`func NewListOrganizationQuotas200ResponsePaginationWithDefaults() *ListOrganizationQuotas200ResponsePagination`

NewListOrganizationQuotas200ResponsePaginationWithDefaults instantiates a new ListOrganizationQuotas200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *ListOrganizationQuotas200ResponsePagination) GetFirst() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetFirstOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ListOrganizationQuotas200ResponsePagination) SetFirst(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated)`

SetFirst sets First field to given value.

### HasFirst

`func (o *ListOrganizationQuotas200ResponsePagination) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *ListOrganizationQuotas200ResponsePagination) GetLast() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetLastOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ListOrganizationQuotas200ResponsePagination) SetLast(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated)`

SetLast sets Last field to given value.

### HasLast

`func (o *ListOrganizationQuotas200ResponsePagination) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *ListOrganizationQuotas200ResponsePagination) GetNext() ListOrganizationQuotas200ResponsePaginationNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetNextOk() (*ListOrganizationQuotas200ResponsePaginationNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListOrganizationQuotas200ResponsePagination) SetNext(v ListOrganizationQuotas200ResponsePaginationNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListOrganizationQuotas200ResponsePagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *ListOrganizationQuotas200ResponsePagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ListOrganizationQuotas200ResponsePagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *ListOrganizationQuotas200ResponsePagination) GetPrevious() ListOrganizationQuotas200ResponsePaginationNext`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetPreviousOk() (*ListOrganizationQuotas200ResponsePaginationNext, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ListOrganizationQuotas200ResponsePagination) SetPrevious(v ListOrganizationQuotas200ResponsePaginationNext)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *ListOrganizationQuotas200ResponsePagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *ListOrganizationQuotas200ResponsePagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *ListOrganizationQuotas200ResponsePagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *ListOrganizationQuotas200ResponsePagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListOrganizationQuotas200ResponsePagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListOrganizationQuotas200ResponsePagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *ListOrganizationQuotas200ResponsePagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ListOrganizationQuotas200ResponsePagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ListOrganizationQuotas200ResponsePagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ListOrganizationQuotas200ResponsePagination) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


