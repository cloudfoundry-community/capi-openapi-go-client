# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp when the security group was created | 
**GloballyEnabled** | [**SecurityGroupGloballyEnabled**](SecurityGroupGloballyEnabled.md) |  | 
**Guid** | **string** | Unique identifier for the security group | 
**Links** | [**SecurityGroupLinks**](SecurityGroupLinks.md) |  | 
**Metadata** | Pointer to [**SecurityGroupMetadata**](SecurityGroupMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the security group | 
**Relationships** | [**SecurityGroupRelationships**](SecurityGroupRelationships.md) |  | 
**Rules** | [**[]Rule**](Rule.md) | Array of egress traffic rules | 
**UpdatedAt** | **time.Time** | Timestamp when the security group was last updated | 

## Methods

### NewSecurityGroup

`func NewSecurityGroup(createdAt time.Time, globallyEnabled SecurityGroupGloballyEnabled, guid string, links SecurityGroupLinks, name string, relationships SecurityGroupRelationships, rules []Rule, updatedAt time.Time, ) *SecurityGroup`

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


### GetMetadata

`func (o *SecurityGroup) GetMetadata() SecurityGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecurityGroup) GetMetadataOk() (*SecurityGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecurityGroup) SetMetadata(v SecurityGroupMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecurityGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


