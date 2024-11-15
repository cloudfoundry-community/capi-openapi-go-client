# ServiceOfferingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**BrokerCatalog** | Pointer to [**BrokerCatalog**](BrokerCatalog.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceOfferingRelationships**](ServiceOfferingRelationships.md) |  | [optional] 
**Requires** | Pointer to **[]string** |  | [optional] 
**Shareable** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceOfferingUpdate

`func NewServiceOfferingUpdate() *ServiceOfferingUpdate`

NewServiceOfferingUpdate instantiates a new ServiceOfferingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferingUpdateWithDefaults

`func NewServiceOfferingUpdateWithDefaults() *ServiceOfferingUpdate`

NewServiceOfferingUpdateWithDefaults instantiates a new ServiceOfferingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ServiceOfferingUpdate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServiceOfferingUpdate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServiceOfferingUpdate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ServiceOfferingUpdate) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBrokerCatalog

`func (o *ServiceOfferingUpdate) GetBrokerCatalog() BrokerCatalog`

GetBrokerCatalog returns the BrokerCatalog field if non-nil, zero value otherwise.

### GetBrokerCatalogOk

`func (o *ServiceOfferingUpdate) GetBrokerCatalogOk() (*BrokerCatalog, bool)`

GetBrokerCatalogOk returns a tuple with the BrokerCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerCatalog

`func (o *ServiceOfferingUpdate) SetBrokerCatalog(v BrokerCatalog)`

SetBrokerCatalog sets BrokerCatalog field to given value.

### HasBrokerCatalog

`func (o *ServiceOfferingUpdate) HasBrokerCatalog() bool`

HasBrokerCatalog returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceOfferingUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceOfferingUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceOfferingUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceOfferingUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *ServiceOfferingUpdate) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *ServiceOfferingUpdate) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *ServiceOfferingUpdate) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *ServiceOfferingUpdate) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceOfferingUpdate) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceOfferingUpdate) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceOfferingUpdate) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceOfferingUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServiceOfferingUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceOfferingUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceOfferingUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceOfferingUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceOfferingUpdate) GetRelationships() ServiceOfferingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceOfferingUpdate) GetRelationshipsOk() (*ServiceOfferingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceOfferingUpdate) SetRelationships(v ServiceOfferingRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceOfferingUpdate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRequires

`func (o *ServiceOfferingUpdate) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *ServiceOfferingUpdate) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *ServiceOfferingUpdate) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *ServiceOfferingUpdate) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetShareable

`func (o *ServiceOfferingUpdate) GetShareable() bool`

GetShareable returns the Shareable field if non-nil, zero value otherwise.

### GetShareableOk

`func (o *ServiceOfferingUpdate) GetShareableOk() (*bool, bool)`

GetShareableOk returns a tuple with the Shareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareable

`func (o *ServiceOfferingUpdate) SetShareable(v bool)`

SetShareable sets Shareable field to given value.

### HasShareable

`func (o *ServiceOfferingUpdate) HasShareable() bool`

HasShareable returns a boolean if a field has been set.

### GetTags

`func (o *ServiceOfferingUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceOfferingUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceOfferingUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceOfferingUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


