# ServiceOfferingCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**BrokerCatalog** | Pointer to [**BrokerCatalog**](BrokerCatalog.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceOfferingRelationships**](ServiceOfferingRelationships.md) |  | [optional] 
**Requires** | Pointer to **[]string** |  | [optional] 
**Shareable** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceOfferingCreate

`func NewServiceOfferingCreate() *ServiceOfferingCreate`

NewServiceOfferingCreate instantiates a new ServiceOfferingCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferingCreateWithDefaults

`func NewServiceOfferingCreateWithDefaults() *ServiceOfferingCreate`

NewServiceOfferingCreateWithDefaults instantiates a new ServiceOfferingCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ServiceOfferingCreate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServiceOfferingCreate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServiceOfferingCreate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ServiceOfferingCreate) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBrokerCatalog

`func (o *ServiceOfferingCreate) GetBrokerCatalog() BrokerCatalog`

GetBrokerCatalog returns the BrokerCatalog field if non-nil, zero value otherwise.

### GetBrokerCatalogOk

`func (o *ServiceOfferingCreate) GetBrokerCatalogOk() (*BrokerCatalog, bool)`

GetBrokerCatalogOk returns a tuple with the BrokerCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerCatalog

`func (o *ServiceOfferingCreate) SetBrokerCatalog(v BrokerCatalog)`

SetBrokerCatalog sets BrokerCatalog field to given value.

### HasBrokerCatalog

`func (o *ServiceOfferingCreate) HasBrokerCatalog() bool`

HasBrokerCatalog returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceOfferingCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceOfferingCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceOfferingCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceOfferingCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *ServiceOfferingCreate) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *ServiceOfferingCreate) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *ServiceOfferingCreate) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *ServiceOfferingCreate) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetName

`func (o *ServiceOfferingCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceOfferingCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceOfferingCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceOfferingCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceOfferingCreate) GetRelationships() ServiceOfferingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceOfferingCreate) GetRelationshipsOk() (*ServiceOfferingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceOfferingCreate) SetRelationships(v ServiceOfferingRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceOfferingCreate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRequires

`func (o *ServiceOfferingCreate) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *ServiceOfferingCreate) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *ServiceOfferingCreate) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *ServiceOfferingCreate) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetShareable

`func (o *ServiceOfferingCreate) GetShareable() bool`

GetShareable returns the Shareable field if non-nil, zero value otherwise.

### GetShareableOk

`func (o *ServiceOfferingCreate) GetShareableOk() (*bool, bool)`

GetShareableOk returns a tuple with the Shareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareable

`func (o *ServiceOfferingCreate) SetShareable(v bool)`

SetShareable sets Shareable field to given value.

### HasShareable

`func (o *ServiceOfferingCreate) HasShareable() bool`

HasShareable returns a boolean if a field has been set.

### GetTags

`func (o *ServiceOfferingCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceOfferingCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceOfferingCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceOfferingCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


