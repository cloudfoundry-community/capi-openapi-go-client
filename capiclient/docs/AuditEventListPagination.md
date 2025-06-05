# AuditEventListPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**AppUsageEventListPaginationFirst**](AppUsageEventListPaginationFirst.md) |  | [optional] 
**Last** | Pointer to [**AppUsageEventListPaginationLast**](AppUsageEventListPaginationLast.md) |  | [optional] 
**Next** | Pointer to **map[string]interface{}** | Link to the next page, if applicable | [optional] 
**Previous** | Pointer to **map[string]interface{}** | Link to the previous page, if applicable | [optional] 
**TotalPages** | Pointer to **int32** | Total number of pages | [optional] 
**TotalResults** | Pointer to **int32** | Total number of results | [optional] 

## Methods

### NewAuditEventListPagination

`func NewAuditEventListPagination() *AuditEventListPagination`

NewAuditEventListPagination instantiates a new AuditEventListPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventListPaginationWithDefaults

`func NewAuditEventListPaginationWithDefaults() *AuditEventListPagination`

NewAuditEventListPaginationWithDefaults instantiates a new AuditEventListPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *AuditEventListPagination) GetFirst() AppUsageEventListPaginationFirst`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *AuditEventListPagination) GetFirstOk() (*AppUsageEventListPaginationFirst, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *AuditEventListPagination) SetFirst(v AppUsageEventListPaginationFirst)`

SetFirst sets First field to given value.

### HasFirst

`func (o *AuditEventListPagination) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *AuditEventListPagination) GetLast() AppUsageEventListPaginationLast`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *AuditEventListPagination) GetLastOk() (*AppUsageEventListPaginationLast, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *AuditEventListPagination) SetLast(v AppUsageEventListPaginationLast)`

SetLast sets Last field to given value.

### HasLast

`func (o *AuditEventListPagination) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *AuditEventListPagination) GetNext() map[string]interface{}`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AuditEventListPagination) GetNextOk() (*map[string]interface{}, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AuditEventListPagination) SetNext(v map[string]interface{})`

SetNext sets Next field to given value.

### HasNext

`func (o *AuditEventListPagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *AuditEventListPagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *AuditEventListPagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *AuditEventListPagination) GetPrevious() map[string]interface{}`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AuditEventListPagination) GetPreviousOk() (*map[string]interface{}, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AuditEventListPagination) SetPrevious(v map[string]interface{})`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AuditEventListPagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *AuditEventListPagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AuditEventListPagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *AuditEventListPagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *AuditEventListPagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *AuditEventListPagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *AuditEventListPagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *AuditEventListPagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *AuditEventListPagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *AuditEventListPagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *AuditEventListPagination) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


