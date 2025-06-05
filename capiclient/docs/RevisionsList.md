# RevisionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**RevisionsListPagination**](RevisionsListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]Revision**](Revision.md) |  | [optional] 

## Methods

### NewRevisionsList

`func NewRevisionsList() *RevisionsList`

NewRevisionsList instantiates a new RevisionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionsListWithDefaults

`func NewRevisionsListWithDefaults() *RevisionsList`

NewRevisionsListWithDefaults instantiates a new RevisionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RevisionsList) GetPagination() RevisionsListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RevisionsList) GetPaginationOk() (*RevisionsListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RevisionsList) SetPagination(v RevisionsListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RevisionsList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *RevisionsList) GetResources() []Revision`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RevisionsList) GetResourcesOk() (*[]Revision, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RevisionsList) SetResources(v []Revision)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RevisionsList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


