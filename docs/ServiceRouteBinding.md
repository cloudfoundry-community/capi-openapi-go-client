# ServiceRouteBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**LastOperation** | Pointer to [**ServiceCredentialBindingLastOperation**](ServiceCredentialBindingLastOperation.md) |  | [optional] 
**Links** | Pointer to [**ServiceRouteBindingLinks**](ServiceRouteBindingLinks.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Relationships** | Pointer to [**ServiceRouteBindingRelationships**](ServiceRouteBindingRelationships.md) |  | [optional] 
**RouteServiceUrl** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceRouteBinding

`func NewServiceRouteBinding() *ServiceRouteBinding`

NewServiceRouteBinding instantiates a new ServiceRouteBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRouteBindingWithDefaults

`func NewServiceRouteBindingWithDefaults() *ServiceRouteBinding`

NewServiceRouteBindingWithDefaults instantiates a new ServiceRouteBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceRouteBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceRouteBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceRouteBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceRouteBinding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *ServiceRouteBinding) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ServiceRouteBinding) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ServiceRouteBinding) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ServiceRouteBinding) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastOperation

`func (o *ServiceRouteBinding) GetLastOperation() ServiceCredentialBindingLastOperation`

GetLastOperation returns the LastOperation field if non-nil, zero value otherwise.

### GetLastOperationOk

`func (o *ServiceRouteBinding) GetLastOperationOk() (*ServiceCredentialBindingLastOperation, bool)`

GetLastOperationOk returns a tuple with the LastOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperation

`func (o *ServiceRouteBinding) SetLastOperation(v ServiceCredentialBindingLastOperation)`

SetLastOperation sets LastOperation field to given value.

### HasLastOperation

`func (o *ServiceRouteBinding) HasLastOperation() bool`

HasLastOperation returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceRouteBinding) GetLinks() ServiceRouteBindingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceRouteBinding) GetLinksOk() (*ServiceRouteBindingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceRouteBinding) SetLinks(v ServiceRouteBindingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceRouteBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceRouteBinding) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceRouteBinding) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceRouteBinding) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceRouteBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceRouteBinding) GetRelationships() ServiceRouteBindingRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceRouteBinding) GetRelationshipsOk() (*ServiceRouteBindingRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceRouteBinding) SetRelationships(v ServiceRouteBindingRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceRouteBinding) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRouteServiceUrl

`func (o *ServiceRouteBinding) GetRouteServiceUrl() string`

GetRouteServiceUrl returns the RouteServiceUrl field if non-nil, zero value otherwise.

### GetRouteServiceUrlOk

`func (o *ServiceRouteBinding) GetRouteServiceUrlOk() (*string, bool)`

GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteServiceUrl

`func (o *ServiceRouteBinding) SetRouteServiceUrl(v string)`

SetRouteServiceUrl sets RouteServiceUrl field to given value.

### HasRouteServiceUrl

`func (o *ServiceRouteBinding) HasRouteServiceUrl() bool`

HasRouteServiceUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceRouteBinding) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceRouteBinding) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceRouteBinding) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceRouteBinding) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


