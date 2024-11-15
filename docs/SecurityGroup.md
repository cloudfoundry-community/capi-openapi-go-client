# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GloballyEnabled** | Pointer to [**SecurityGroupGloballyEnabled**](SecurityGroupGloballyEnabled.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**SecurityGroupLinks**](SecurityGroupLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**SecurityGroupRelationships**](SecurityGroupRelationships.md) |  | [optional] 
**Rules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSecurityGroup

`func NewSecurityGroup() *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGloballyEnabled

`func (o *SecurityGroup) GetGloballyEnabled() SecurityGroupGloballyEnabled`

GetGloballyEnabled returns the GloballyEnabled field if non-nil, zero value otherwise.

### GetGloballyEnabledOk

`func (o *SecurityGroup) GetGloballyEnabledOk() (*SecurityGroupGloballyEnabled, bool)`

GetGloballyEnabledOk returns a tuple with the GloballyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyEnabled

`func (o *SecurityGroup) SetGloballyEnabled(v SecurityGroupGloballyEnabled)`

SetGloballyEnabled sets GloballyEnabled field to given value.

### HasGloballyEnabled

`func (o *SecurityGroup) HasGloballyEnabled() bool`

HasGloballyEnabled returns a boolean if a field has been set.

### GetGuid

`func (o *SecurityGroup) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SecurityGroup) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SecurityGroup) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SecurityGroup) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityGroup) GetLinks() SecurityGroupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityGroup) GetLinksOk() (*SecurityGroupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityGroup) SetLinks(v SecurityGroupLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *SecurityGroup) GetRelationships() SecurityGroupRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SecurityGroup) GetRelationshipsOk() (*SecurityGroupRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SecurityGroup) SetRelationships(v SecurityGroupRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SecurityGroup) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRules

`func (o *SecurityGroup) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroup) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroup) SetRules(v []Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SecurityGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecurityGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecurityGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SecurityGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


