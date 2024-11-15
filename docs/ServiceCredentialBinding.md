# ServiceCredentialBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**LastOperation** | Pointer to [**ServiceCredentialBindingLastOperation**](ServiceCredentialBindingLastOperation.md) |  | [optional] 
**Links** | Pointer to [**ServiceCredentialBindingLinks**](ServiceCredentialBindingLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceCredentialBindingRelationships**](ServiceCredentialBindingRelationships.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceCredentialBinding

`func NewServiceCredentialBinding() *ServiceCredentialBinding`

NewServiceCredentialBinding instantiates a new ServiceCredentialBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCredentialBindingWithDefaults

`func NewServiceCredentialBindingWithDefaults() *ServiceCredentialBinding`

NewServiceCredentialBindingWithDefaults instantiates a new ServiceCredentialBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceCredentialBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceCredentialBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceCredentialBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceCredentialBinding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceCredentialBinding) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceCredentialBinding) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceCredentialBinding) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceCredentialBinding) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastOperation

`func (o *ServiceCredentialBinding) GetLastOperation() ServiceCredentialBindingLastOperation`

GetLastOperation returns the LastOperation field if non-nil, zero value otherwise.

### GetLastOperationOk

`func (o *ServiceCredentialBinding) GetLastOperationOk() (*ServiceCredentialBindingLastOperation, bool)`

GetLastOperationOk returns a tuple with the LastOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperation

`func (o *ServiceCredentialBinding) SetLastOperation(v ServiceCredentialBindingLastOperation)`

SetLastOperation sets LastOperation field to given value.

### HasLastOperation

`func (o *ServiceCredentialBinding) HasLastOperation() bool`

HasLastOperation returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceCredentialBinding) GetLinks() ServiceCredentialBindingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceCredentialBinding) GetLinksOk() (*ServiceCredentialBindingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceCredentialBinding) SetLinks(v ServiceCredentialBindingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceCredentialBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceCredentialBinding) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceCredentialBinding) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceCredentialBinding) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceCredentialBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ServiceCredentialBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCredentialBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCredentialBinding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCredentialBinding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceCredentialBinding) GetRelationships() ServiceCredentialBindingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceCredentialBinding) GetRelationshipsOk() (*ServiceCredentialBindingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceCredentialBinding) SetRelationships(v ServiceCredentialBindingRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceCredentialBinding) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ServiceCredentialBinding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceCredentialBinding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceCredentialBinding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceCredentialBinding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceCredentialBinding) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceCredentialBinding) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceCredentialBinding) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceCredentialBinding) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


