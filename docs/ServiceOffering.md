# ServiceOffering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**BrokerCatalog** | Pointer to [**BrokerCatalog**](BrokerCatalog.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ServiceOfferingLinks**](ServiceOfferingLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3DropletsGuidPatchRequestMetadata**](V3DropletsGuidPatchRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceOfferingRelationships**](ServiceOfferingRelationships.md) |  | [optional] 
**Requires** | Pointer to **[]string** |  | [optional] 
**Shareable** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceOffering

`func NewServiceOffering() *ServiceOffering`

NewServiceOffering instantiates a new ServiceOffering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferingWithDefaults

`func NewServiceOfferingWithDefaults() *ServiceOffering`

NewServiceOfferingWithDefaults instantiates a new ServiceOffering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ServiceOffering) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServiceOffering) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServiceOffering) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ServiceOffering) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBrokerCatalog

`func (o *ServiceOffering) GetBrokerCatalog() BrokerCatalog`

GetBrokerCatalog returns the BrokerCatalog field if non-nil, zero value otherwise.

### GetBrokerCatalogOk

`func (o *ServiceOffering) GetBrokerCatalogOk() (*BrokerCatalog, bool)`

GetBrokerCatalogOk returns a tuple with the BrokerCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerCatalog

`func (o *ServiceOffering) SetBrokerCatalog(v BrokerCatalog)`

SetBrokerCatalog sets BrokerCatalog field to given value.

### HasBrokerCatalog

`func (o *ServiceOffering) HasBrokerCatalog() bool`

HasBrokerCatalog returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceOffering) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceOffering) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceOffering) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceOffering) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceOffering) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceOffering) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceOffering) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceOffering) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *ServiceOffering) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *ServiceOffering) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *ServiceOffering) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *ServiceOffering) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceOffering) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceOffering) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceOffering) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceOffering) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceOffering) GetLinks() ServiceOfferingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceOffering) GetLinksOk() (*ServiceOfferingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceOffering) SetLinks(v ServiceOfferingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceOffering) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceOffering) GetMetadata() V3DropletsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceOffering) GetMetadataOk() (*V3DropletsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceOffering) SetMetadata(v V3DropletsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceOffering) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServiceOffering) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceOffering) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceOffering) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceOffering) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceOffering) GetRelationships() ServiceOfferingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceOffering) GetRelationshipsOk() (*ServiceOfferingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceOffering) SetRelationships(v ServiceOfferingRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceOffering) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRequires

`func (o *ServiceOffering) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *ServiceOffering) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *ServiceOffering) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *ServiceOffering) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetShareable

`func (o *ServiceOffering) GetShareable() bool`

GetShareable returns the Shareable field if non-nil, zero value otherwise.

### GetShareableOk

`func (o *ServiceOffering) GetShareableOk() (*bool, bool)`

GetShareableOk returns a tuple with the Shareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareable

`func (o *ServiceOffering) SetShareable(v bool)`

SetShareable sets Shareable field to given value.

### HasShareable

`func (o *ServiceOffering) HasShareable() bool`

HasShareable returns a boolean if a field has been set.

### GetTags

`func (o *ServiceOffering) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceOffering) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceOffering) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceOffering) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceOffering) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceOffering) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceOffering) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceOffering) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


