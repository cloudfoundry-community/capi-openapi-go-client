# AppUsageEventListPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**AppUsageEventListPaginationFirst**](AppUsageEventListPaginationFirst.md) |  | [optional] 
**Last** | Pointer to [**AppUsageEventListPaginationLast**](AppUsageEventListPaginationLast.md) |  | [optional] 
**Next** | Pointer to [**AppUsageEventListPaginationNext**](AppUsageEventListPaginationNext.md) |  | [optional] 
**Previous** | Pointer to **map[string]interface{}** | Link to the previous page, if applicable | [optional] 
**TotalPages** | Pointer to **int32** | Total number of pages | [optional] 
**TotalResults** | Pointer to **int32** | Total number of results | [optional] 

## Methods

### NewAppUsageEventListPagination

`func NewAppUsageEventListPagination() *AppUsageEventListPagination`

NewAppUsageEventListPagination instantiates a new AppUsageEventListPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageEventListPaginationWithDefaults

`func NewAppUsageEventListPaginationWithDefaults() *AppUsageEventListPagination`

NewAppUsageEventListPaginationWithDefaults instantiates a new AppUsageEventListPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *AppUsageEventListPagination) GetFirst() AppUsageEventListPaginationFirst`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *AppUsageEventListPagination) GetFirstOk() (*AppUsageEventListPaginationFirst, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *AppUsageEventListPagination) SetFirst(v AppUsageEventListPaginationFirst)`

SetFirst sets First field to given value.

### HasFirst

`func (o *AppUsageEventListPagination) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *AppUsageEventListPagination) GetLast() AppUsageEventListPaginationLast`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *AppUsageEventListPagination) GetLastOk() (*AppUsageEventListPaginationLast, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *AppUsageEventListPagination) SetLast(v AppUsageEventListPaginationLast)`

SetLast sets Last field to given value.

### HasLast

`func (o *AppUsageEventListPagination) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *AppUsageEventListPagination) GetNext() AppUsageEventListPaginationNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AppUsageEventListPagination) GetNextOk() (*AppUsageEventListPaginationNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AppUsageEventListPagination) SetNext(v AppUsageEventListPaginationNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *AppUsageEventListPagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *AppUsageEventListPagination) GetPrevious() map[string]interface{}`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AppUsageEventListPagination) GetPreviousOk() (*map[string]interface{}, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AppUsageEventListPagination) SetPrevious(v map[string]interface{})`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AppUsageEventListPagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *AppUsageEventListPagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AppUsageEventListPagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *AppUsageEventListPagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *AppUsageEventListPagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *AppUsageEventListPagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *AppUsageEventListPagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *AppUsageEventListPagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *AppUsageEventListPagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *AppUsageEventListPagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *AppUsageEventListPagination) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


