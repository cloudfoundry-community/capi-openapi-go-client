# AuditEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**AuditEventListPagination**](AuditEventListPagination.md) |  | [optional] 
**Resources** | Pointer to [**[]AuditEvent**](AuditEvent.md) |  | [optional] 

## Methods

### NewAuditEventList

`func NewAuditEventList() *AuditEventList`

NewAuditEventList instantiates a new AuditEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventListWithDefaults

`func NewAuditEventListWithDefaults() *AuditEventList`

NewAuditEventListWithDefaults instantiates a new AuditEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AuditEventList) GetPagination() AuditEventListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AuditEventList) GetPaginationOk() (*AuditEventListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AuditEventList) SetPagination(v AuditEventListPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AuditEventList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResources

`func (o *AuditEventList) GetResources() []AuditEvent`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AuditEventList) GetResourcesOk() (*[]AuditEvent, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AuditEventList) SetResources(v []AuditEvent)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AuditEventList) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


