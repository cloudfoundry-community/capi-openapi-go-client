# AuditEventOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Unique identifier for the organization where the event occurred | [optional] 

## Methods

### NewAuditEventOrganization

`func NewAuditEventOrganization() *AuditEventOrganization`

NewAuditEventOrganization instantiates a new AuditEventOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventOrganizationWithDefaults

`func NewAuditEventOrganizationWithDefaults() *AuditEventOrganization`

NewAuditEventOrganizationWithDefaults instantiates a new AuditEventOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AuditEventOrganization) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AuditEventOrganization) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AuditEventOrganization) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AuditEventOrganization) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

