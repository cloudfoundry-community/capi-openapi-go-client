# V3PackagesGet200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated.md) |  | [optional] 
**Last** | Pointer to [**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated.md) |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewV3PackagesGet200ResponsePagination

`func NewV3PackagesGet200ResponsePagination() *V3PackagesGet200ResponsePagination`

NewV3PackagesGet200ResponsePagination instantiates a new V3PackagesGet200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PackagesGet200ResponsePaginationWithDefaults

`func NewV3PackagesGet200ResponsePaginationWithDefaults() *V3PackagesGet200ResponsePagination`

NewV3PackagesGet200ResponsePaginationWithDefaults instantiates a new V3PackagesGet200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *V3PackagesGet200ResponsePagination) GetFirst() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *V3PackagesGet200ResponsePagination) GetFirstOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *V3PackagesGet200ResponsePagination) SetFirst(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated)`

SetFirst sets First field to given value.

### HasFirst

`func (o *V3PackagesGet200ResponsePagination) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *V3PackagesGet200ResponsePagination) GetLast() V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *V3PackagesGet200ResponsePagination) GetLastOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *V3PackagesGet200ResponsePagination) SetLast(v V3IsolationSegmentsGuidRelationshipsOrganizationsGet200ResponseLinksRelated)`

SetLast sets Last field to given value.

### HasLast

`func (o *V3PackagesGet200ResponsePagination) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *V3PackagesGet200ResponsePagination) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *V3PackagesGet200ResponsePagination) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *V3PackagesGet200ResponsePagination) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *V3PackagesGet200ResponsePagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *V3PackagesGet200ResponsePagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *V3PackagesGet200ResponsePagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *V3PackagesGet200ResponsePagination) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *V3PackagesGet200ResponsePagination) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *V3PackagesGet200ResponsePagination) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *V3PackagesGet200ResponsePagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *V3PackagesGet200ResponsePagination) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *V3PackagesGet200ResponsePagination) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalPages

`func (o *V3PackagesGet200ResponsePagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *V3PackagesGet200ResponsePagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *V3PackagesGet200ResponsePagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *V3PackagesGet200ResponsePagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *V3PackagesGet200ResponsePagination) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *V3PackagesGet200ResponsePagination) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *V3PackagesGet200ResponsePagination) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *V3PackagesGet200ResponsePagination) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


