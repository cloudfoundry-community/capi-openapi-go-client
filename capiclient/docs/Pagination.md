# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | [**PaginationFirst**](PaginationFirst.md) |  | 
**Last** | [**PaginationLast**](PaginationLast.md) |  | 
**Next** | Pointer to [**NullablePaginationNext**](PaginationNext.md) |  | [optional] 
**Previous** | Pointer to [**NullablePaginationPrevious**](PaginationPrevious.md) |  | [optional] 
**TotalPages** | **int32** | Total number of pages | 
**TotalResults** | **int32** | Total number of results across all pages | 

## Methods

### NewPagination

`func NewPagination(first PaginationFirst, last PaginationLast, totalPages int32, totalResults int32, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *Pagination) GetFirst() PaginationFirst`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *Pagination) GetFirstOk() (*PaginationFirst, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *Pagination) SetFirst(v PaginationFirst)`

SetFirst sets First field to given value.


### GetLast

`func (o *Pagination) GetLast() PaginationLast`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Pagination) GetLastOk() (*PaginationLast, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Pagination) SetLast(v PaginationLast)`

SetLast sets Last field to given value.


### GetNext

`func (o *Pagination) GetNext() PaginationNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Pagination) GetNextOk() (*PaginationNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Pagination) SetNext(v PaginationNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *Pagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *Pagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *Pagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *Pagination) GetPrevious() PaginationPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Pagination) GetPreviousOk() (*PaginationPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Pagination) SetPrevious(v PaginationPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Pagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *Pagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *Pagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *Pagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Pagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Pagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetTotalResults

`func (o *Pagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *Pagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *Pagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


