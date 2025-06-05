# ServiceBroker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ServiceBrokerLinks**](ServiceBrokerLinks.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**V3ServiceBrokersPostRequestRelationships**](V3ServiceBrokersPostRequestRelationships.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceBroker

`func NewServiceBroker() *ServiceBroker`

NewServiceBroker instantiates a new ServiceBroker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBrokerWithDefaults

`func NewServiceBrokerWithDefaults() *ServiceBroker`

NewServiceBrokerWithDefaults instantiates a new ServiceBroker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceBroker) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceBroker) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceBroker) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceBroker) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceBroker) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceBroker) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceBroker) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceBroker) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceBroker) GetLinks() ServiceBrokerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceBroker) GetLinksOk() (*ServiceBrokerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceBroker) SetLinks(v ServiceBrokerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceBroker) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceBroker) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceBroker) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceBroker) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceBroker) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServiceBroker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceBroker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceBroker) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceBroker) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceBroker) GetRelationships() V3ServiceBrokersPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceBroker) GetRelationshipsOk() (*V3ServiceBrokersPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceBroker) SetRelationships(v V3ServiceBrokersPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceBroker) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceBroker) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceBroker) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceBroker) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceBroker) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceBroker) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceBroker) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceBroker) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceBroker) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


