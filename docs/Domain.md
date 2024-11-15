# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**DomainLinks**](DomainLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**DomainRelationships**](DomainRelationships.md) |  | [optional] 
**RouterGroup** | Pointer to [**V3DropletsPostRequestRelationshipsAppData**](V3DropletsPostRequestRelationshipsAppData.md) |  | [optional] 
**SupportedProtocols** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Domain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Domain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *Domain) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Domain) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Domain) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Domain) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetInternal

`func (o *Domain) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Domain) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Domain) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Domain) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLinks

`func (o *Domain) GetLinks() DomainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Domain) GetLinksOk() (*DomainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Domain) SetLinks(v DomainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Domain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Domain) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Domain) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Domain) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Domain) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *Domain) GetRelationships() DomainRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Domain) GetRelationshipsOk() (*DomainRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Domain) SetRelationships(v DomainRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Domain) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRouterGroup

`func (o *Domain) GetRouterGroup() V3DropletsPostRequestRelationshipsAppData`

GetRouterGroup returns the RouterGroup field if non-nil, zero value otherwise.

### GetRouterGroupOk

`func (o *Domain) GetRouterGroupOk() (*V3DropletsPostRequestRelationshipsAppData, bool)`

GetRouterGroupOk returns a tuple with the RouterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterGroup

`func (o *Domain) SetRouterGroup(v V3DropletsPostRequestRelationshipsAppData)`

SetRouterGroup sets RouterGroup field to given value.

### HasRouterGroup

`func (o *Domain) HasRouterGroup() bool`

HasRouterGroup returns a boolean if a field has been set.

### GetSupportedProtocols

`func (o *Domain) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *Domain) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *Domain) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.

### HasSupportedProtocols

`func (o *Domain) HasSupportedProtocols() bool`

HasSupportedProtocols returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Domain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


