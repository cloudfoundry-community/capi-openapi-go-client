# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**OrganizationLinks**](OrganizationLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3DropletsGuidPatchRequestMetadata**](V3DropletsGuidPatchRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**OrganizationRelationships**](OrganizationRelationships.md) |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *Organization) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Organization) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Organization) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Organization) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *Organization) GetLinks() OrganizationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Organization) GetLinksOk() (*OrganizationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Organization) SetLinks(v OrganizationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Organization) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Organization) GetMetadata() V3DropletsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Organization) GetMetadataOk() (*V3DropletsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Organization) SetMetadata(v V3DropletsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Organization) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *Organization) GetRelationships() OrganizationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Organization) GetRelationshipsOk() (*OrganizationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Organization) SetRelationships(v OrganizationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Organization) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSuspended

`func (o *Organization) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Organization) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Organization) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Organization) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Organization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


