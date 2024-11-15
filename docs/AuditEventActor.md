# AuditEventActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Unique identifier for the actor | [optional] 
**Name** | Pointer to **string** | The name of the actor | [optional] 
**Type** | Pointer to **string** | The actor type | [optional] 

## Methods

### NewAuditEventActor

`func NewAuditEventActor() *AuditEventActor`

NewAuditEventActor instantiates a new AuditEventActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventActorWithDefaults

`func NewAuditEventActorWithDefaults() *AuditEventActor`

NewAuditEventActorWithDefaults instantiates a new AuditEventActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AuditEventActor) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AuditEventActor) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AuditEventActor) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AuditEventActor) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *AuditEventActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditEventActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditEventActor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditEventActor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuditEventActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditEventActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditEventActor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditEventActor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


