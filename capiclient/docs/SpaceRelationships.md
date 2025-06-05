# SpaceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**ToOneRelationship**](ToOneRelationship.md) |  | [optional] 
**Quota** | Pointer to [**NullableSpaceRelationshipsQuota**](SpaceRelationshipsQuota.md) |  | [optional] 

## Methods

### NewSpaceRelationships

`func NewSpaceRelationships() *SpaceRelationships`

NewSpaceRelationships instantiates a new SpaceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRelationshipsWithDefaults

`func NewSpaceRelationshipsWithDefaults() *SpaceRelationships`

NewSpaceRelationshipsWithDefaults instantiates a new SpaceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SpaceRelationships) GetOrganization() ToOneRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SpaceRelationships) GetOrganizationOk() (*ToOneRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SpaceRelationships) SetOrganization(v ToOneRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SpaceRelationships) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetQuota

`func (o *SpaceRelationships) GetQuota() SpaceRelationshipsQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *SpaceRelationships) GetQuotaOk() (*SpaceRelationshipsQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *SpaceRelationships) SetQuota(v SpaceRelationshipsQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *SpaceRelationships) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *SpaceRelationships) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *SpaceRelationships) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


