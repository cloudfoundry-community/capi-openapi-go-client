# RevisionsListPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Last** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Next** | Pointer to [**NullableV3Get200ResponseLinksServiceInstances**](V3Get200ResponseLinksServiceInstances.md) |  | [optional] 
**Previous** | Pointer to [**NullableV3Get200ResponseLinksServiceInstances**](V3Get200ResponseLinksServiceInstances.md) |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewRevisionsListPagination

`func NewRevisionsListPagination() *RevisionsListPagination`

NewRevisionsListPagination instantiates a new RevisionsListPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionsListPaginationWithDefaults

`func NewRevisionsListPaginationWithDefaults() *RevisionsListPagination`

NewRevisionsListPaginationWithDefaults instantiates a new RevisionsListPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *RevisionsListPagination) GetFirst() Get200ResponseLinksLogCache`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *RevisionsListPagination) GetFirstOk() (*Get200ResponseLinksLogCache, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *RevisionsListPagination) SetFirst(v Get200ResponseLinksLogCache)`

SetFirst sets First field to given value.

### HasFirst

`func (o *RevisionsListPagination) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *RevisionsListPagination) GetLast() Get200ResponseLinksLogCache`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *RevisionsListPagination) GetLastOk() (*Get200ResponseLinksLogCache, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *RevisionsListPagination) SetLast(v Get200ResponseLinksLogCache)`

SetLast sets Last field to given value.

### HasLast

`func (o *RevisionsListPagination) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *RevisionsListPagination) GetNext() V3Get200ResponseLinksServiceInstances`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *RevisionsListPagination) GetNextOk() (*V3Get200ResponseLinksServiceInstances, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *RevisionsListPagination) SetNext(v V3Get200ResponseLinksServiceInstances)`

SetNext sets Next field to given value.

### HasNext

`func (o *RevisionsListPagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *RevisionsListPagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *RevisionsListPagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *RevisionsListPagination) GetPrevious() V3Get200ResponseLinksServiceInstances`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *RevisionsListPagination) GetPreviousOk() (*V3Get200ResponseLinksServiceInstances, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *RevisionsListPagination) SetPrevious(v V3Get200ResponseLinksServiceInstances)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *RevisionsListPagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *RevisionsListPagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *RevisionsListPagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *RevisionsListPagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *RevisionsListPagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *RevisionsListPagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *RevisionsListPagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *RevisionsListPagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RevisionsListPagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RevisionsListPagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *RevisionsListPagination) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


