# AuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to [**AuditEventActor**](AuditEventActor.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time with zone when the object was created | [optional] 
**Data** | Pointer to **map[string]interface{}** | Additional information about event | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the event | [optional] 
**Links** | Pointer to [**AuditEventLinks**](AuditEventLinks.md) |  | [optional] 
**Organization** | Pointer to [**AuditEventOrganization**](AuditEventOrganization.md) |  | [optional] 
**Space** | Pointer to [**AuditEventSpace**](AuditEventSpace.md) |  | [optional] 
**Target** | Pointer to [**AuditEventTarget**](AuditEventTarget.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the event | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time with zone when the object was last updated | [optional] 

## Methods

### NewAuditEvent

`func NewAuditEvent() *AuditEvent`

NewAuditEvent instantiates a new AuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventWithDefaults

`func NewAuditEventWithDefaults() *AuditEvent`

NewAuditEventWithDefaults instantiates a new AuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AuditEvent) GetActor() AuditEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AuditEvent) GetActorOk() (*AuditEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AuditEvent) SetActor(v AuditEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AuditEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *AuditEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AuditEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGuid

`func (o *AuditEvent) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AuditEvent) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AuditEvent) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AuditEvent) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *AuditEvent) GetLinks() AuditEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditEvent) GetLinksOk() (*AuditEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditEvent) SetLinks(v AuditEventLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuditEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrganization

`func (o *AuditEvent) GetOrganization() AuditEventOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AuditEvent) GetOrganizationOk() (*AuditEventOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AuditEvent) SetOrganization(v AuditEventOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AuditEvent) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSpace

`func (o *AuditEvent) GetSpace() AuditEventSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *AuditEvent) GetSpaceOk() (*AuditEventSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *AuditEvent) SetSpace(v AuditEventSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *AuditEvent) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTarget

`func (o *AuditEvent) GetTarget() AuditEventTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditEvent) GetTargetOk() (*AuditEventTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditEvent) SetTarget(v AuditEventTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditEvent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetType

`func (o *AuditEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


