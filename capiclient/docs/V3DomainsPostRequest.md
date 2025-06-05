# V3DomainsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Internal** | Pointer to **bool** | Whether the domain is used for internal traffic | [optional] 
**Metadata** | Pointer to [**V3DomainsPostRequestMetadata**](V3DomainsPostRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the domain | [optional] 
**Organization** | Pointer to [**V3DomainsPostRequestOrganization**](V3DomainsPostRequestOrganization.md) |  | [optional] 
**RouterGroup** | Pointer to [**V3DomainsPostRequestRouterGroup**](V3DomainsPostRequestRouterGroup.md) |  | [optional] 
**SharedOrganizations** | Pointer to [**[]V3DomainsPostRequestOrganizationData**](V3DomainsPostRequestOrganizationData.md) |  | [optional] 

## Methods

### NewV3DomainsPostRequest

`func NewV3DomainsPostRequest() *V3DomainsPostRequest`

NewV3DomainsPostRequest instantiates a new V3DomainsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3DomainsPostRequestWithDefaults

`func NewV3DomainsPostRequestWithDefaults() *V3DomainsPostRequest`

NewV3DomainsPostRequestWithDefaults instantiates a new V3DomainsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *V3DomainsPostRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *V3DomainsPostRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *V3DomainsPostRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *V3DomainsPostRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMetadata

`func (o *V3DomainsPostRequest) GetMetadata() V3DomainsPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3DomainsPostRequest) GetMetadataOk() (*V3DomainsPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3DomainsPostRequest) SetMetadata(v V3DomainsPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3DomainsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3DomainsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3DomainsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3DomainsPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3DomainsPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *V3DomainsPostRequest) GetOrganization() V3DomainsPostRequestOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V3DomainsPostRequest) GetOrganizationOk() (*V3DomainsPostRequestOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V3DomainsPostRequest) SetOrganization(v V3DomainsPostRequestOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *V3DomainsPostRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRouterGroup

`func (o *V3DomainsPostRequest) GetRouterGroup() V3DomainsPostRequestRouterGroup`

GetRouterGroup returns the RouterGroup field if non-nil, zero value otherwise.

### GetRouterGroupOk

`func (o *V3DomainsPostRequest) GetRouterGroupOk() (*V3DomainsPostRequestRouterGroup, bool)`

GetRouterGroupOk returns a tuple with the RouterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterGroup

`func (o *V3DomainsPostRequest) SetRouterGroup(v V3DomainsPostRequestRouterGroup)`

SetRouterGroup sets RouterGroup field to given value.

### HasRouterGroup

`func (o *V3DomainsPostRequest) HasRouterGroup() bool`

HasRouterGroup returns a boolean if a field has been set.

### GetSharedOrganizations

`func (o *V3DomainsPostRequest) GetSharedOrganizations() []V3DomainsPostRequestOrganizationData`

GetSharedOrganizations returns the SharedOrganizations field if non-nil, zero value otherwise.

### GetSharedOrganizationsOk

`func (o *V3DomainsPostRequest) GetSharedOrganizationsOk() (*[]V3DomainsPostRequestOrganizationData, bool)`

GetSharedOrganizationsOk returns a tuple with the SharedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOrganizations

`func (o *V3DomainsPostRequest) SetSharedOrganizations(v []V3DomainsPostRequestOrganizationData)`

SetSharedOrganizations sets SharedOrganizations field to given value.

### HasSharedOrganizations

`func (o *V3DomainsPostRequest) HasSharedOrganizations() bool`

HasSharedOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


